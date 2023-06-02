package ssoauth

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/logback"
	"gorm.io/gorm"
)

type Configurer interface {
	Reset()
	Load(context.Context) (addr string, err error)
	ID() string
	Shared() bool
	Validate([]byte) error
}

func NewConfigure(slog logback.Logger) Configurer {
	return &fromDB{
		id:      "global.sso.url",
		slog:    slog,
		timeout: 5 * time.Second,
	}
}

type fromDB struct {
	id      string
	slog    logback.Logger
	timeout time.Duration
	mutex   sync.RWMutex
	loaded  bool
	err     error
	addr    string
}

func (f *fromDB) Reset() {
	f.mutex.Lock()
	f.loaded, f.err, f.addr = false, nil, ""
	f.mutex.Unlock()
}

func (f *fromDB) Load(ctx context.Context) (string, error) {
	f.mutex.RLock()
	loaded, err, addr := f.loaded, f.err, f.addr
	f.mutex.RUnlock()
	if loaded {
		return addr, err
	}
	return f.slowLoad(ctx)
}

func (f *fromDB) slowLoad(ctx context.Context) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// 类似于双重检查锁
	loaded, err, addr := f.loaded, f.err, f.addr
	if loaded {
		return addr, err
	}

	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	f.loaded = true // 标记为已经加载
	tbl := query.Store
	dat, err := tbl.WithContext(ctx).Where(tbl.ID.Eq(f.id)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("sso 服务器配置 %s 没有找到", f.id)
		} else {
			f.err = err
		}
		return "", err
	}

	addr = string(dat.Value)
	// 预先拼凑好 addr 地址
	if strings.Contains(addr, "?") {
		addr += "&"
	} else {
		addr += "?"
	}
	f.err = nil
	f.addr = addr

	return addr, nil
}

func (f *fromDB) ID() string {
	return f.id
}

func (f *fromDB) Shared() bool {
	return true
}

func (f *fromDB) Validate(dat []byte) error {
	u, err := url.Parse(string(dat))
	if err == nil && u.Scheme != "http" && u.Scheme != "https" {
		err = fmt.Errorf("必须是 http 链接地址")
	}
	return err
}
