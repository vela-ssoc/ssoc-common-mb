package memcache

import (
	"context"
	"sync"
)

type Map[K comparable, T any] struct {
	fun     func(context.Context, K) (T, error)
	nocache NocacheFunc
	mtx     sync.RWMutex
	idx     map[K]*resultHodler[T]
}

func NewMap[K comparable, T any](fun func(context.Context, K) (T, error), nocache ...NocacheFunc) *Map[K, T] {
	return &Map[K, T]{
		fun:     fun,
		nocache: firstNocache(nocache),
	}
}

func (c *Map[K, T]) Load(ctx context.Context, k K) (T, error) {
	c.mtx.RLock()
	r := c.idx[k]
	c.mtx.RUnlock()
	if r != nil {
		return r.t, nil
	}

	return c.slowLoad(ctx, k)
}

func (c *Map[K, T]) Forget(k K) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.idx, k)
}

func (c *Map[K, T]) ForgetAll() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.idx = nil
}

func (c *Map[K, T]) slowLoad(ctx context.Context, k K) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if r := c.idx[k]; r != nil {
		return r.t, nil
	}

	t, e := c.fun(ctx, k)
	if e != nil && c.nocache != nil && c.nocache(e) {
		return t, e
	}

	r := &resultHodler[T]{t: t}
	if c.idx == nil {
		c.idx = make(map[K]*resultHodler[T], 8)
	}
	c.idx[k] = r

	return t, e
}
