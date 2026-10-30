package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CMDBConfig struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"-"`
	Endpoint  string        `bson:"endpoint"             json:"endpoint"`
	Enabled   bool          `bson:"enabled"              json:"enabled"`
	Key       string        `bson:"key"                  json:"key"`
	Secret    string        `bson:"secret"               json:"secret"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (CMDBConfig) CollectionInfo() (string, string) { return "cmdb_config", "CMDB 配置" }
