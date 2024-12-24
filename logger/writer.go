package logger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewWriter(console bool, lumber *lumberjack.Logger) io.WriteCloser {
	return nil
}

type logWriter struct {
	w io.Writer
	c io.Closer
}

func (lw *logWriter) Write(p []byte) (int, error) {
	return lw.w.Write(p)
}

func (lw *logWriter) Close() error {
	if c := lw.c; c != nil {
		return c.Close()
	}

	return nil
}
