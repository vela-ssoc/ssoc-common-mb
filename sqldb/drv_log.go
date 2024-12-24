package sqldb

import (
	"context"
	"log/slog"

	pq "gitee.com/opengauss/openGauss-connector-go-pq"
)

type mysqlLog struct {
	log *slog.Logger
}

func (ml *mysqlLog) Print(vs ...any) {
	size := len(vs)
	if size == 0 {
		return
	}

	msg := "mysql"
	arg0 := vs[0]
	switch v := arg0.(type) {
	case error:
		msg = v.Error()
		vs = vs[1:]
	case string:
		msg = v
		vs = vs[1:]
	}
	if len(vs) == 0 {
		ml.log.Info(msg)
	} else {
		ml.log.Info(msg, slog.Any("data", vs))
	}
}

type gaussLog struct {
	log *slog.Logger
}

func (gl *gaussLog) Log(ctx context.Context, level pq.LogLevel, msg string, data map[string]any) {
	lvl := gl.mappingLevel(level)
	gl.log.Log(ctx, lvl, msg, data)
}

func (gl *gaussLog) mappingLevel(lvl pq.LogLevel) slog.Level {
	switch lvl {
	case pq.LogLevelError:
		return slog.LevelError
	case pq.LogLevelWarn:
		return slog.LevelWarn
	case pq.LogLevelInfo:
		return slog.LevelInfo
	case pq.LogLevelDebug, pq.LogLevelTrace:
		return slog.LevelDebug
	default:
		return slog.LevelDebug - 1
	}
}
