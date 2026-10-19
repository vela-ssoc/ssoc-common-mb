package memcache

import (
	"context"

	"golang.org/x/sync/singleflight"
)

type SingleFlight[I, V any] struct {
	sig singleflight.Group
	fun func(ctx context.Context, in I) (V, error)
}

func NewSingleFlight[I, V any](fun func(ctx context.Context, in I) (V, error)) *SingleFlight[I, V] {
	return &SingleFlight[I, V]{
		fun: fun,
	}
}

func (sf *SingleFlight[I, V]) Do(ctx context.Context, key string, args I) (V, error, bool) {
	val, err, shared := sf.sig.Do(key, func() (any, error) { return sf.fun(ctx, args) })
	if err != nil {
		var v V
		return v, err, shared
	}
	v, _ := val.(V)

	return v, nil, shared
}
