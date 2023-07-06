package devops

import (
	"context"

	"github.com/vela-ssoc/vela-common-mb/storage"
)

type Configurer interface {
	Load(ctx context.Context) (string, error)
}

func NewConfig(store storage.Storer) Configurer {
	return &config{
		store: store,
	}
}

type config struct {
	store storage.Storer
}

func (cfg *config) Load(ctx context.Context) (string, error) {
	return cfg.store.AlarmURL(ctx)
}
