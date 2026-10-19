package memcache

import (
	"container/list"
	"context"
	"sync"
)

type LRU[K comparable, V any] struct {
	size    int
	load    func(context.Context, K) (V, error)
	onevict func(K, V)
	items   map[K]*list.Element
	evict   *list.List
	mutex   sync.Mutex
}

// NewLRU 创建一个 LRU 缓存。
// 注意 load 函数的调用性能问题。
func NewLRU[K comparable, V any](size int, load func(context.Context, K) (V, error), onevict func(K, V)) *LRU[K, V] {
	if size <= 0 {
		size = 8
	}

	return &LRU[K, V]{
		size:    size,
		load:    load,
		onevict: onevict,
		items:   make(map[K]*list.Element, size),
		evict:   list.New(),
	}
}

func (l *LRU[K, V]) Get(ctx context.Context, key K) (V, error) {
	ent, elim, err := l.lockGet(ctx, key)
	if err != nil {
		var zero V
		return zero, err
	}

	if elim != nil && l.onevict != nil {
		l.onevict(elim.key, elim.value)
	}

	return ent.value, nil
}

func (l *LRU[K, V]) ForgetAll() {
	olds := l.lockForgetAll()
	if len(olds) == 0 || l.onevict == nil {
		return
	}

	for _, ele := range olds {
		l.onEvict(ele)
	}
}

func (l *LRU[K, V]) Forget(key K) {
	ele := l.lockForget(key)
	if ele != nil && l.onevict != nil {
		l.onEvict(ele)
	}
}

func (l *LRU[K, V]) lockGet(ctx context.Context, key K) (*entry[K, V], *entry[K, V], error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 命中缓存
	if ele := l.items[key]; ele != nil {
		val := ele.Value
		ent := val.(*entry[K, V])

		l.evict.MoveToFront(ele)

		return ent, nil, nil
	}

	v, err := l.load(ctx, key)
	if err != nil {
		return nil, nil, err
	}

	ent := &entry[K, V]{key: key, value: v}
	ele := l.evict.PushFront(ent)
	l.items[key] = ele

	// 检查容量并执行淘汰逻辑
	var elim *entry[K, V]
	if l.evict.Len() > l.size {
		back := l.evict.Back()
		if back != nil {
			l.evict.Remove(back)
			val := back.Value
			elim = val.(*entry[K, V])
			delete(l.items, elim.key)
		}
	}

	return ent, elim, nil
}

func (l *LRU[K, V]) lockForget(key K) *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if ele, ok := l.items[key]; ok {
		l.evict.Remove(ele)
		delete(l.items, key)

		return ele
	}

	return nil
}

func (l *LRU[K, V]) lockForgetAll() map[K]*list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	olds := l.items
	l.evict = list.New()
	l.items = make(map[K]*list.Element, l.size)

	return olds
}

func (l *LRU[K, V]) onEvict(ele *list.Element) {
	val := ele.Value
	ent := val.(*entry[K, V])

	l.onevict(ent.key, ent.value)
}
