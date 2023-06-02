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

type SearchConfigurer interface {
	Reset()
	Load() (addr string, auth string, err error)
}

func NewSearchConfigure() SearchConfigurer {
	return &searchConfigure{}
}

func NewForwardConfigure(name string) ForwardConfigurer {
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

type searchConfigure struct {
	mutex sync.RWMutex
	done  bool
	addr  string
	auth  string
	err   error
}

func (sc *searchConfigure) Reset() {
	sc.mutex.Lock()
	sc.done, sc.addr, sc.auth, sc.err = false, "", "", nil
	sc.mutex.Unlock()
}

func (sc *searchConfigure) Load() (addr string, auth string, err error) {
	var done bool
	sc.mutex.RLock()
	done, addr, auth, err = sc.done, sc.addr, sc.auth, sc.err
	sc.mutex.RUnlock()
	if done {
		return
	}
	return sc.slowLoad()
}

func (sc *searchConfigure) slowLoad() (addr, auth string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if sc.done {
		return sc.addr, sc.auth, sc.err
	}

	tbl := query.Elastic
	dat, err := tbl.WithContext(ctx).Where(tbl.Enable.Is(true)).First()
	if err != nil {
		sc.done, sc.err = true, err
		return "", "", err
	}

	addr = dat.Host
	if dat.Username != "" || dat.Password != "" {
		bs64 := base64.StdEncoding.EncodeToString([]byte(dat.Username + ":" + dat.Password))
		auth = "Basic " + bs64
	}

	sc.addr = addr
	sc.auth = auth
	sc.err = nil
	sc.done = true

	return addr, auth, nil
}

type ForwardConfigurer interface {
	Reset()
	Load(ctx context.Context) (http.Handler, error)
}

type forwardConfig struct {
	name   string
	trip   http.RoundTripper
	mutex  sync.RWMutex
	loaded bool
	err    error
	proxy  *httputil.ReverseProxy
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
	return fc.loadSlow(ctx)
}

func (fc *forwardConfig) loadSlow(ctx context.Context) (http.Handler, error) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()
	if fc.loaded {
		return fc.proxy, fc.err
	}

	tbl := query.Elastic
	dat, err := tbl.WithContext(ctx).Where(tbl.Enable.Is(true)).First()
	if err != nil {
		fc.err = err
		return nil, err
	}

	proxy, err := fc.newProxy(dat.Host, dat.Username, dat.Password)
	if err != nil {
		fc.err = err
		return nil, err
	}
	fc.proxy = proxy
	fc.err = nil

	return proxy, nil
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
