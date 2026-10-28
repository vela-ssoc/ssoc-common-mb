package memcache

import (
	"context"
	"errors"
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

type NocacheFunc func(err error) bool

func NocacheDeadline(err error) bool {
	return err != nil && errors.Is(err, context.DeadlineExceeded)
}

func firstNocache(nocache []NocacheFunc) NocacheFunc {
	if len(nocache) == 0 {
		return nil
	}

	return nocache[0]
}
