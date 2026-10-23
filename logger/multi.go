package logger

import (
	"context"
	"log/slog"
	"reflect"
	"sync"
	"sync/atomic"
)

func NewMultiHandler(hs ...slog.Handler) *MultiHandler {
	m := new(MultiHandler)
	m.Replace(hs...)
	return m
}

type handlerSnapshot struct {
	handlers []slog.Handler
	multi    *slog.MultiHandler
}

type MultiHandler struct {
	holder atomic.Pointer[handlerSnapshot]
	mutex  sync.Mutex
	order  []slog.Handler
	unique map[slog.Handler]struct{}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return m.load().multi.Enabled(ctx, level)
}

func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	return m.load().multi.Handle(ctx, record)
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &multiHandlerView{
		parent: m,
		ops: []handlerOp{{
			attrs: append([]slog.Attr(nil), attrs...),
		}},
	}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	return &multiHandlerView{
		parent: m,
		ops: []handlerOp{{
			group:   name,
			isGroup: true,
		}},
	}
}

func (m *MultiHandler) Append(hs ...slog.Handler) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.unique == nil {
		m.unique = make(map[slog.Handler]struct{}, len(hs))
	}

	for _, h := range hs {
		if isNilHandler(h) {
			continue
		}
		requireComparable(h)

		if _, ok := m.unique[h]; ok {
			continue
		}

		m.unique[h] = struct{}{}
		m.order = append(m.order, h)
	}

	m.publish()
}

func (m *MultiHandler) Remove(hs ...slog.Handler) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, h := range hs {
		if isNilHandler(h) {
			continue
		}
		requireComparable(h)

		if _, ok := m.unique[h]; !ok {
			continue
		}

		delete(m.unique, h)

		for i, v := range m.order {
			if v == h {
				copy(m.order[i:], m.order[i+1:])
				m.order[len(m.order)-1] = nil
				m.order = m.order[:len(m.order)-1]
				break
			}
		}
	}

	m.publish()
}

func (m *MultiHandler) Replace(hs ...slog.Handler) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	unique := make(map[slog.Handler]struct{}, len(hs))
	order := make([]slog.Handler, 0, len(hs))

	for _, h := range hs {
		if isNilHandler(h) {
			continue
		}
		requireComparable(h)

		if _, ok := unique[h]; ok {
			continue
		}

		unique[h] = struct{}{}
		order = append(order, h)
	}

	m.order = order
	m.unique = unique
	m.publish()
}

func (m *MultiHandler) publish() {
	handlers := append([]slog.Handler(nil), m.order...)

	m.holder.Store(&handlerSnapshot{
		handlers: handlers,
		multi:    slog.NewMultiHandler(handlers...),
	})
}

func (m *MultiHandler) load() *handlerSnapshot {
	if s := m.holder.Load(); s != nil {
		return s
	}

	return &handlerSnapshot{
		multi: slog.NewMultiHandler(),
	}
}

func isNilHandler(h slog.Handler) bool {
	if h == nil {
		return true
	}

	v := reflect.ValueOf(h)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func requireComparable(h slog.Handler) {
	if !reflect.TypeOf(h).Comparable() {
		panic("logger: slog.Handler dynamic type must be comparable")
	}
}

// handlerOp 保留 WithAttrs 和 WithGroup 的调用顺序。
type handlerOp struct {
	attrs   []slog.Attr
	group   string
	isGroup bool
}

type multiHandlerView struct {
	parent *MultiHandler
	ops    []handlerOp
}

func (v *multiHandlerView) Enabled(
	ctx context.Context,
	level slog.Level,
) bool {
	return v.load().Enabled(ctx, level)
}

func (v *multiHandlerView) Handle(
	ctx context.Context,
	record slog.Record,
) error {
	return v.load().Handle(ctx, record)
}

func (v *multiHandlerView) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &multiHandlerView{
		parent: v.parent,
		ops: appendOps(v.ops, handlerOp{
			attrs: append([]slog.Attr(nil), attrs...),
		}),
	}
}

func (v *multiHandlerView) WithGroup(name string) slog.Handler {
	return &multiHandlerView{
		parent: v.parent,
		ops: appendOps(v.ops, handlerOp{
			group:   name,
			isGroup: true,
		}),
	}
}

func (v *multiHandlerView) load() *slog.MultiHandler {
	snapshot := v.parent.load()

	handlers := make([]slog.Handler, 0, len(snapshot.handlers))
	for _, h := range snapshot.handlers {
		for _, op := range v.ops {
			if op.isGroup {
				h = h.WithGroup(op.group)
			} else {
				h = h.WithAttrs(op.attrs)
			}
		}
		handlers = append(handlers, h)
	}

	return slog.NewMultiHandler(handlers...)
}

func appendOps(ops []handlerOp, op handlerOp) []handlerOp {
	result := make([]handlerOp, len(ops), len(ops)+1)
	copy(result, ops)
	return append(result, op)
}
