package profile

import "time"

type Database struct {
	DSN         string   `json:"dsn"           validate:"required,lte=255"`                      // 数据库连接
	Level       string   `json:"level"         validate:"omitempty,oneof=DEBUG INFO WARN ERROR"` // 日志输出级别
	Migrate     bool     `json:"migrate"`                                                        // 是否合并差异（broker 无视该参数）
	MaxOpenConn int      `json:"max_open_conn" validate:"gte=0"`                                 // 最大连接数
	MaxIdleConn int      `json:"max_idle_conn" validate:"gte=0"`                                 // 最大空闲连接数
	MaxLifeTime duration `json:"max_life_time" validate:"gte=0"`                                 // 连接最大存活时长
	MaxIdleTime duration `json:"max_idle_time" validate:"gte=0"`                                 // 空闲连接最大时长
}

type duration time.Duration

func (du *duration) UnmarshalText(bs []byte) error {
	pd, err := time.ParseDuration(string(bs))
	if err == nil {
		*du = duration(pd)
	}

	return err
}

func (du duration) MarshalText() ([]byte, error) {
	return []byte(du.String()), nil
}

func (du *duration) String() string {
	return time.Duration(*du).String()
}

func (du *duration) Duration() time.Duration {
	return time.Duration(*du)
}
