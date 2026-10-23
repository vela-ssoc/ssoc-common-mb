package logger

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func NewTint(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = new(slog.HandlerOptions)
	}

	topts := &tint.Options{
		AddSource:   opts.AddSource,
		Level:       opts.Level,
		ReplaceAttr: opts.ReplaceAttr,
		TimeFormat:  time.RFC3339,
	}

	return tint.NewTextHandler(w, topts)
}
