package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ManagerConfig struct {
	ID        bson.ObjectID        `bson:"_id,omitempty"        json:"id"`
	Server    *ManagerServerConfig `bson:"server"               json:"server"`
	Logger    *LoggerConfig        `bson:"logger"               json:"logger"`
	CreatedAt time.Time            `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at,omitempty" json:"updated_at"`
}

func (ManagerConfig) CollectionInfo() (string, string) {
	return "manager_config", "中心端配置"
}

type ManagerServerConfig struct {
	Addr   string                         `bson:"addr"   json:"addr"`   // HTTP 监听地址
	Static map[string]ManagerStaticConfig `bson:"static" json:"static"` // 静态资源映射
}

type ManagerStaticConfig struct {
	SPA bool   `bson:"spa" json:"spa"`                     // 是否单页面应用
	Dir string `bson:"dir" json:"dir" validate:"required"` // 本地目录
}
