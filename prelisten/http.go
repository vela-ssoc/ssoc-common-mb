package prelisten

import (
	"crypto/tls"
	"net"
	"net/http"
)

func WrapHTTP(h http.Handler) func(net.Conn) {
	return (&httpFunc{h: h}).serve
}

func WrapHTTPS(h http.Handler) func(net.Conn) {
	return (&httpFunc{h: h}).serve
}

type httpFunc struct {
	h http.Handler
}

func (h *httpFunc) serve(conn net.Conn) {
	lis := &onceAcceptListener{conn: conn}
	_ = http.Serve(lis, h.h)
}

type httpsFunc struct {
	tlsCfg *tls.Config
	h      http.Handler
}

func (h *httpsFunc) serve(conn net.Conn) {
	lis := &onceAcceptListener{conn: conn}
	_ = http.Serve(tls.NewListener(lis, h.tlsCfg), h.h)
}
