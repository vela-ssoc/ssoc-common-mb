package elastic

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/problem"
)

func NewConfigure(name string) Configurer {
	dialer := &net.Dialer{Timeout: 3 * time.Second}
	tlsDialer := &tls.Dialer{NetDialer: dialer}
	trip := &http.Transport{
		DialContext:    dialer.DialContext,
		DialTLSContext: tlsDialer.DialContext,
	}

	return &forwardConfig{
		name: name,
		trip: trip,
	}
}

type Configurer interface {
	Reset()
	Load(ctx context.Context) (http.Handler, error)
	LoadAddr(ctx context.Context) (host, auth string, err error)
}

type forwardConfig struct {
	name   string
	trip   http.RoundTripper
	mutex  sync.RWMutex
	loaded bool
	err    error
	proxy  *httputil.ReverseProxy

	addr string
	auth string
}

func (fc *forwardConfig) Reset() {
	fc.mutex.Lock()
	fc.loaded, fc.err, fc.proxy = false, nil, nil
	fc.mutex.Unlock()
}

func (fc *forwardConfig) Load(ctx context.Context) (http.Handler, error) {
	fc.mutex.RLock()
	loaded, err, proxy := fc.loaded, fc.err, fc.proxy
	fc.mutex.RUnlock()
	if loaded {
		return proxy, err
	}
	p, _, _, exx := fc.loadSlow(ctx)

	return p, exx
}

func (fc *forwardConfig) LoadAddr(ctx context.Context) (string, string, error) {
	fc.mutex.RLock()
	loaded, err, addr, auth := fc.loaded, fc.err, fc.addr, fc.auth
	fc.mutex.RUnlock()
	if loaded {
		return addr, auth, err
	}
	_, dest, ath, err := fc.loadSlow(ctx)

	return dest, ath, err
}

func (fc *forwardConfig) loadSlow(ctx context.Context) (http.Handler, string, string, error) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	if fc.loaded {
		return fc.proxy, fc.addr, fc.auth, fc.err
	}

	fc.loaded = true
	tbl := query.Elastic
	dat, err := tbl.WithContext(ctx).Where(tbl.Enable.Is(true)).First()
	if err != nil {
		fc.err = err
		return nil, "", "", err
	}

	var auth string
	addr := dat.Host
	if dat.Username != "" || dat.Password != "" {
		bs64 := base64.StdEncoding.EncodeToString([]byte(dat.Username + ":" + dat.Password))
		auth = "Basic " + bs64
	}

	proxy, err := fc.newProxy(dat.Host, dat.Username, dat.Password)
	if err != nil {
		fc.err = err
		return nil, "", "", err
	}
	fc.proxy = proxy
	fc.err = nil
	fc.addr = addr
	fc.auth = auth

	return proxy, addr, auth, nil
}

// newProxy 初始化创建代理，支持 BasicAuth
func (fc *forwardConfig) newProxy(addr, uname, passwd string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	rewriteFunc := func(r *httputil.ProxyRequest) {
		r.SetXForwarded()
		r.SetURL(u)
		r.Out.SetBasicAuth(uname, passwd)
	}

	px := &httputil.ReverseProxy{
		Transport:      fc.trip,
		Rewrite:        rewriteFunc,
		ModifyResponse: fc.modifyResponse,
		ErrorHandler:   fc.errorHandler,
	}

	return px, nil
}

func (fc *forwardConfig) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := err.(*net.OpError); ok {
		// 隐藏后端服务 IP
		e.Addr = nil
		e.Net += " elasticsearch"
		err = e
	}

	pd := &problem.Detail{
		Type:     fc.name,
		Title:    "代理错误",
		Status:   http.StatusBadRequest,
		Detail:   err.Error(),
		Instance: r.RequestURI,
	}
	_ = pd.JSON(w)
}

func (fc *forwardConfig) modifyResponse(w *http.Response) error {
	if w.StatusCode == http.StatusUnauthorized {
		w.StatusCode = http.StatusBadRequest
	}
	return nil
}
