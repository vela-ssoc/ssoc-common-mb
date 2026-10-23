package logger

import (
	"context"
	"log/slog"
	"runtime"
)

func Skip(h slog.Handler, n int) slog.Handler {
	if n <= 0 {
		return h
	}

	return &skipHandler{
		h: h,
		n: n,
	}
}

type skipHandler struct {
	h slog.Handler
	n int
}

func (s *skipHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return s.h.Enabled(ctx, level)
}

func (s *skipHandler) Handle(ctx context.Context, record slog.Record) error {
	var pcs [1]uintptr

	// 跳过 runtime.Callers 自身、当前 Handle，
	// 再根据 n 跳过指定数量的调用栈帧。
	if runtime.Callers(2+s.n, pcs[:]) > 0 {
		record.PC = pcs[0]
	}

	return s.h.Handle(ctx, record)
}

func (s *skipHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &skipHandler{
		h: s.h.WithAttrs(attrs),
		n: s.n,
	}
}

func (s *skipHandler) WithGroup(name string) slog.Handler {
	return &skipHandler{
		h: s.h.WithGroup(name),
		n: s.n,
	}
}
