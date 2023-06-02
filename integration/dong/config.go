package dong

import (
	"context"
	"fmt"
	"sync"

	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"gorm.io/gorm"
)

type Configurer interface {
	Reset()
	Load(ctx context.Context) (addr, account, token string, err error)
}

func NewConfigure() Configurer {
	return &configure{}
}

type configure struct {
	mutex   sync.RWMutex
	loaded  bool
	err     error
	addr    string
	account string
	token   string
}

func (cf *configure) Reset() {
	cf.mutex.Lock()
	cf.loaded, cf.err, cf.addr, cf.account, cf.token = false, nil, "", "", ""
	cf.mutex.Unlock()
}

func (cf *configure) Load(ctx context.Context) (string, string, string, error) {
	cf.mutex.RLock()
	loaded, err, addr, account, token := cf.loaded, cf.err, cf.addr, cf.account, cf.token
	cf.mutex.RUnlock()
	if loaded {
		return addr, account, token, err
	}
	return cf.slowLoad(ctx)
}

func (cf *configure) slowLoad(ctx context.Context) (string, string, string, error) {
	cf.mutex.Lock()
	defer cf.mutex.Unlock()
	if cf.loaded {
		return cf.addr, cf.account, cf.token, cf.err
	}

	cf.loaded = true

	tbl := query.Emc
	dat, err := tbl.WithContext(ctx).Where(tbl.Enable.Is(true)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("没有找到咚咚服务号配置")
		} else {
			cf.err = err
		}
		return "", "", "", err
	}
	account := dat.Account
	token := dat.Token
	addr := dat.Host
	cf.account = account
	cf.addr = addr
	cf.token = token

	return addr, account, token, nil
}
