package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ReverseProxy struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"-"`          // ID
	Location  string        `bson:"location"             json:"location"`   // Location
	Name      string        `bson:"name"                 json:"name"`       // 名字方便用户理解
	Manager   bool          `bson:"manager"              json:"manager"`    // 是否在 manager 节点启用
	Broker    bool          `bson:"broker"               json:"broker"`     // 是否在 broker 节点启用
	Header    HTTPHeader    `bson:"header"               json:"header"`     // HTTP Header
	Upstream  string        `bson:"upstream"             json:"upstream"`   // 上游服务
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"` // 创建时间。
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"` // 更新时间。
}

func (ReverseProxy) CollectionInfo() (string, string) { return "reverse_proxy", "反向代理" }
