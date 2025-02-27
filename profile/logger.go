package profile

import (
	"io"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*lumberjack.Logger
	Console bool   `json:"console"`
	Level   string `json:"level"   validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
}

func (l Logger) Close() error {
	if lg := l.Logger; lg != nil {
		return lg.Close()
	}

	return nil
}

func (l Logger) LogWriter() LogWriter {
	level := new(slog.LevelVar)
	if err := level.UnmarshalText([]byte(l.Level)); err != nil {
		level.Set(slog.LevelInfo)
	}
	writers := make(map[io.Writer]struct{}, 8)
	if l.Console {
		writers[os.Stdout] = struct{}{}
	}
	if lumber := l.Logger; lumber != nil && lumber.Filename != "" {
		writers[lumber] = struct{}{}
	}

	lw := &logWriters{level: level, writers: writers}
	lw.refreshWriter()

	return lw
}

type LogWriter interface {
	io.Writer

	// Level 获取当前日志级别。
	Level() *slog.LevelVar

	// Attach 附加一个 Writer 到日志输出。
	Attach(w io.Writer) bool

	// Detach 卸载掉输出。
	Detach(w io.Writer) bool

	// Discard 将日志输出导向 [io.Discard]，等价于重置日志输出。
	Discard()
}

type logWriters struct {
	level   *slog.LevelVar
	mutex   sync.Mutex
	writers map[io.Writer]struct{}
	val     atomic.Pointer[packWriter]
}

func (lw *logWriters) Write(p []byte) (int, error) {
	pw := lw.val.Load()
	return pw.Write(p)
}

func (lw *logWriters) Level() *slog.LevelVar {
	return lw.level
}

func (lw *logWriters) Attach(w io.Writer) bool {
	if w == nil {
		return false
	}

	lw.mutex.Lock()
	defer lw.mutex.Unlock()
	lw.writers[w] = struct{}{}
	lw.refreshWriter()

	return true
}

func (lw *logWriters) Detach(w io.Writer) bool {
	lw.mutex.Lock()
	defer lw.mutex.Unlock()

	if _, exists := lw.writers[w]; !exists {
		return false
	}

	delete(lw.writers, w)
	lw.refreshWriter()

	return true
}

func (lw *logWriters) Discard() {
	lw.mutex.Lock()
	defer lw.mutex.Unlock()

	clear(lw.writers)
	lw.refreshWriter()
}

func (lw *logWriters) refreshWriter() {
	var ws []io.Writer
	for w := range lw.writers {
		ws = append(ws, w)
	}

	pw := new(packWriter)
	switch len(ws) {
	case 0:
		pw.w = io.Discard
	case 1:
		pw.w = ws[0]
	default:
		pw.w = io.MultiWriter(ws...)
	}
	lw.val.Store(pw)
}

type packWriter struct {
	w io.Writer
}

func (pw *packWriter) Write(p []byte) (int, error) {
	return pw.w.Write(p)
}
