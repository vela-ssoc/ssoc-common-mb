package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type RedisConfig struct {
	ID          bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Enabled     bool          `bson:"enabled"              json:"enabled"`
	Addresses   []string      `bson:"addresses"            json:"addresses"`    // 集群模式填写多个，哨兵模式填写哨兵地址
	Username    string        `bson:"username"             json:"username"`     // Redis 6+
	Password    string        `bson:"password"             json:"password"`     // 密码
	ClusterMode bool          `bson:"cluster_mode"         json:"cluster_mode"` // 是否集群模式
	MasterName  string        `bson:"master_name"          json:"master_name"`  // 哨兵模式必填
	DB          int           `bson:"db"                   json:"db"`
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (RedisConfig) CollectionInfo() (string, string) {
	return "redis_config", "redis 配置"
}
