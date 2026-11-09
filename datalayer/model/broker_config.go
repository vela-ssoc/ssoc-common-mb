package model

import "go.mongodb.org/mongo-driver/v2/bson"

type BrokerConfig struct {
	Server   *BrokerServerConfig   `bson:"server"   json:"server"`
	Logger   *LoggerConfig         `bson:"logger"   json:"logger"`
	Database *BrokerDatabaseConfig `bson:"database" json:"database"`
}

type BrokerServerConfig struct {
	Addr string `bson:"addr" json:"addr"` // 中继节点自身服务监听地址
}

// BrokerDatabaseConfig 中继节点的数据库和中心端的数据库必须是同一个数据库，
// 但可能存在特殊网络环境，同一个数据库，从不同节点视角看，地址是不一样的。
//
// 例如在 中心端直连 10.0.0.2:27017，但是网络环境不一样，在中继节点可能只能通过 HAProxy
// 才可以访问。
//
// 注意：错误填写可能会导致中继节点服务启动异常。
type BrokerDatabaseConfig struct {
	URI string `bson:"uri" json:"uri"` // MongoDB 地址
	DSN string `bson:"dsn" json:"dsn"` // 关系数据库地址
}

type BrokerConfigParams struct {
	ID   bson.ObjectID `json:"id"   validate:"required"`
	Slug string        `json:"slug" validate:"required"`
	IDV1 int64         `json:"idv1"`
	Name string        `json:"name"`
	BrokerConfig
}
