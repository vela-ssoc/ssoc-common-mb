package memcache

import (
	"context"
	"sync"
	"weak"
)

type WeakMap[K comparable, T any] struct {
	fun     func(context.Context, K) (T, error)
	nocache NocacheFunc
	mtx     sync.RWMutex
	idx     map[K]weak.Pointer[resultHodler[T]]
}

func NewWeakMap[K comparable, T any](fun func(context.Context, K) (T, error), nocache ...NocacheFunc) *WeakMap[K, T] {
	return &WeakMap[K, T]{
		fun:     fun,
		nocache: firstNocache(nocache),
	}
}

func (wm *WeakMap[K, T]) Load(ctx context.Context, k K) (T, error) {
	wm.mtx.RLock()
	p, ok := wm.idx[k]
	wm.mtx.RUnlock()
	if ok { // 是否已缓存
		if v := p.Value(); v != nil { // 缓存用的是弱指针，检查是否被回收。
			return v.t, nil
		}
	}

	return wm.slowLoad(ctx, k)
}

func (wm *WeakMap[K, T]) Forget(k K) {
	wm.mtx.Lock()
	defer wm.mtx.Unlock()

	delete(wm.idx, k)
}

func (wm *WeakMap[K, T]) ForgetAll() {
	wm.mtx.Lock()
	defer wm.mtx.Unlock()

	wm.idx = nil
}

func (wm *WeakMap[K, T]) slowLoad(ctx context.Context, k K) (T, error) {
	wm.mtx.Lock()
	defer wm.mtx.Unlock()

	if p, ok := wm.idx[k]; ok {
		if v := p.Value(); v != nil {
			return v.t, nil
		}
		delete(wm.idx, k) // 弱指针已失效，顺带删除
	}

	t, e := wm.fun(ctx, k)
	if e != nil && wm.nocache != nil && wm.nocache(e) {
		return t, e
	}

	v := &resultHodler[T]{t: t}
	if wm.idx == nil {
		wm.idx = make(map[K]weak.Pointer[resultHodler[T]], 8)
	}
	wm.idx[k] = weak.Make(v)

	return t, e
}
