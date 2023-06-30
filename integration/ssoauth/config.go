package ssoauth

import (
	"context"

	"github.com/vela-ssoc/vela-common-mb/storage"
)

type Configurer interface {
	Load(context.Context) (addr string, err error)
}

func NewConfigure(store storage.Storer) Configurer {
	return &fromDB{
		store: store,
	}
}

type fromDB struct {
	store storage.Storer
}

func (f *fromDB) Load(ctx context.Context) (string, error) {
	return f.store.SsoURL(ctx)
}
