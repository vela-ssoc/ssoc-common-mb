package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type RedisConfig struct {
	ID         bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Mode       string        `bson:"mode,omitempty"       json:"mode"` // standalone/replication/sentinel/cluster
	Enabled    bool          `bson:"enabled"              json:"enabled"`
	Addresses  []string      `bson:"addresses"            json:"addresses"`
	Username   string        `bson:"username"             json:"username"`
	Password   string        `bson:"password"             json:"password"`
	DB         int           `bson:"db"                   json:"db"`
	MasterName string        `bson:"master_name"          json:"master_name"` // 哨兵模式必填
	CreatedAt  time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (RedisConfig) CollectionInfo() (string, string) {
	return "redis_config", "redis 配置"
}
