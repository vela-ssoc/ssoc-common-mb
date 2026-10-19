package memcache

import (
	"context"
	"sync"
	"sync/atomic"
)

type Value[T any] struct {
	fun func(context.Context) (T, error)
	mtx sync.Mutex
	ptr atomic.Pointer[resultHodler[T]]
}

func NewValue[T any](fun func(context.Context) (T, error)) *Value[T] {
	return &Value[T]{
		fun: fun,
	}
}

func (c *Value[T]) Load(ctx context.Context) (T, error) {
	if r := c.ptr.Load(); r != nil {
		return r.t, nil
	}

	return c.slowLoad(ctx)
}

func (c *Value[T]) Forget() (T, error) {
	if r := c.ptr.Swap(nil); r != nil {
		return r.t, nil
	}

	var t T

	return t, nil
}

func (c *Value[T]) slowLoad(ctx context.Context) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if d := c.ptr.Load(); d != nil {
		return d.t, nil
	}

	t, e := c.fun(ctx)
	if e != nil {
		return t, e
	}

	d := &resultHodler[T]{t: t}
	c.ptr.Store(d)

	return t, e
}
