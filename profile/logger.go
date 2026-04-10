package profile

import "gopkg.in/natefinch/lumberjack.v2"

type Logger struct {
	*lumberjack.Logger
	Console bool   `json:"console"`
	Level   string `json:"level"   validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
}

func (l Logger) Lumber() *lumberjack.Logger {
	if lum := l.Logger; lum != nil && lum.Filename != "" {
		return lum
	}

	return nil
}
