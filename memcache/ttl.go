package memcache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type TTL[T any] struct {
	ttl time.Duration
	fun func(context.Context) (T, error)
	mtx sync.Mutex
	ptr atomic.Pointer[resultHodler[T]]
}

func NewTTL[T any](ttl time.Duration, fun func(context.Context) (T, error)) *TTL[T] {
	return &TTL[T]{
		ttl: ttl,
		fun: fun,
	}
}

func (c *TTL[T]) Load(ctx context.Context) (T, error) {
	now := time.Now()
	if r := c.ptr.Load(); r != nil && r.alive(now) {
		return r.t, nil
	}

	return c.slowLoad(ctx, now)
}

func (c *TTL[T]) Forget() {
	c.ptr.Store(nil)
}

func (c *TTL[T]) slowLoad(ctx context.Context, now time.Time) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if r := c.ptr.Load(); r != nil && r.alive(now) {
		return r.t, nil
	}

	t, e := c.fun(ctx)
	if c.ttl <= 0 || e != nil { // TTL <= 0 无需缓存，每次都缓存穿透。
		return t, e
	}

	now = time.Now()
	d := &resultHodler[T]{t: t, expires: now.Add(c.ttl)}
	c.ptr.Store(d)

	return t, e
}
