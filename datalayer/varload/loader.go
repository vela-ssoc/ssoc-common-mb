package varload

import (
	"context"
	"os"
)

type Loader interface {
	Load(ctx context.Context, v any) error
}

type LoadFunc func(ctx context.Context, v any) error

func (f LoadFunc) Load(ctx context.Context, v any) error {
	return f(ctx, v)
}

type JSONC string

func (j JSONC) Load(_ context.Context, v any) error {
	raw, err := os.ReadFile(string(j))
	if err != nil {
		return err
	}

	return unmarshalJSONC(raw, v)
}

type Loaders struct {
	loads  []Loader
	ignore func(error) bool
}

func NewLoaders(first Loader, ignore func(error) bool, loads ...Loader) Loaders {
	all := append([]Loader{first}, loads...)
	return Loaders{loads: all, ignore: ignore}
}

func (ls Loaders) Load(ctx context.Context, v any) error {
	var err error
	for _, ld := range ls.loads {
		if err = ld.Load(ctx, v); err == nil {
			return nil
		}
		if ls.ignore == nil || !ls.ignore(err) {
			return err
		}
	}

	return err
}
