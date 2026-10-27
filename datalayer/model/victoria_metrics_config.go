package model

import (
	"path"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VictoriaMetricsConfig struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Name      string        `bson:"name"                 json:"name"`
	Enabled   bool          `bson:"enabled"              json:"enabled"`
	Address   string        `bson:"address"              json:"address"`
	TenantID  string        `bson:"tenant_id"            json:"tenant_id"`
	Username  string        `bson:"username"             json:"username"`
	Password  string        `bson:"password"             json:"password"`
	Headers   HTTPHeader    `bson:"headers,omitempty"    json:"headers"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (VictoriaMetricsConfig) CollectionInfo() (string, string) {
	return "victoria_metrics_config", "VictoriaMetrics 数据库配置"
}

func (c VictoriaMetricsConfig) SelectPrefixPath() string {
	return c.prefixPath("select")
}

func (c VictoriaMetricsConfig) InsertPrefixPath() string {
	return c.prefixPath("insert")
}

func (c VictoriaMetricsConfig) prefixPath(typ string) string {
	if c.TenantID == "" {
		return ""
	}

	return path.Join(typ, c.TenantID, "prometheus")
}
