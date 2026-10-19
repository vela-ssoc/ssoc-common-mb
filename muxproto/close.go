package muxproto

import (
	"io"
	"sync/atomic"
)

type OnceCloser interface {
	Close()
	Closed() bool
}

func NewOnceCloser(c io.Closer) OnceCloser {
	return &onceCloser{c: c}
}

type onceCloser struct {
	f atomic.Bool
	c io.Closer
}

func (o *onceCloser) Close() {
	if o.f.CompareAndSwap(false, true) {
		_ = o.c.Close()
	}
}

func (o *onceCloser) Closed() bool {
	return o.f.Load()
}
