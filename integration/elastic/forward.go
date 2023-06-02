package elastic

import (
	"context"
	"net/http"
)

type Forwarder interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

func NewForward(cfg ForwardConfigurer) Forwarder {
	return &esForward{cfg: cfg}
}

type esForward struct {
	cfg ForwardConfigurer
}

func (es *esForward) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	h, err := es.cfg.Load(ctx)
	if err == nil {
		h.ServeHTTP(w, r)
	}
	return err
}
