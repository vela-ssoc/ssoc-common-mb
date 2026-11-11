package memcache

import (
	"context"
	"errors"
	"time"
)

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

type result[T any] struct {
	val T
	err error
	exp time.Time
}

func (r *result[T]) isAlive(now time.Time) bool {
	return r != nil && r.exp.After(now)
}

func (r *result[T]) get() (T, error) {
	return r.val, r.err
}

type entry[K comparable, V any] struct {
	key   K
	value V
}
