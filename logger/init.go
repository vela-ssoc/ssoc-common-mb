package logger

import (
	"io"
	"os"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewInit(filename string) (*slog.Logger, io.Closer) {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	var cls io.Closer = new(nopCloser)
	hs := []slog.Handler{NewTint(os.Stdout, opts)}
	if filename != "" {
		lumber := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    100,
			MaxAge:     180,
			MaxBackups: 100,
			LocalTime:  true,
			Compress:   true,
		}
		hs = append(hs, slog.NewJSONHandler(lumber, opts))
		cls = lumber
	}
	h := slog.NewMultiHandler(hs...)

	return slog.New(h), cls
}

type nopCloser struct{}

func (nopCloser) Close() error { return nil }
