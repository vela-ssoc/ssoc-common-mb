package model

type LoggerConfig struct {
	Level      string `bson:"level"      json:"level"      validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
	Console    bool   `bson:"console"    json:"console"`
	Filename   string `bson:"filename"   json:"filename"   validate:"lte=255"`
	MaxSize    int    `bson:"maxsize"    json:"maxsize"    validate:"gte=0"`
	MaxAge     int    `bson:"maxage"     json:"maxage"     validate:"gte=0"`
	MaxBackups int    `bson:"maxbackups" json:"maxbackups" validate:"gte=0"`
	LocalTime  bool   `bson:"localtime"  json:"localtime"`
	Compress   bool   `bson:"compress"   json:"compress"`
}
