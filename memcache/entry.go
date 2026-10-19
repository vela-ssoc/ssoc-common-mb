package memcache

import (
	"time"
)

type resultHodler[T any] struct {
	t       T         // 结果
	expires time.Time // 过期时间（仅在 TTL 模式下有效）
}

func (e *resultHodler[T]) alive(at time.Time) bool {
	return !e.expires.Before(at)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}
