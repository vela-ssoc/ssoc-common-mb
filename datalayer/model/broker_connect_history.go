package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type BrokerConnectHistory struct {
	ID             bson.ObjectID `bson:"_id,omitempty"   json:"id"`
	BrokerID       bson.ObjectID `bson:"broker_id"       json:"broker_id"`
	ExecuteStat    ExecuteStat   `bson:"execute_stat"    json:"execute_stat,omitzero"`
	TunnelStat     TunnelStat    `bson:"tunnel_stat"     json:"tunnel_stat"`
	DisconnectedAt time.Time     `bson:"disconnected_at" json:"disconnected_at"`
}

func (BrokerConnectHistory) CollectionInfo() (string, string) {
	return "broker_connect_history", "中继节点连接历史记录"
}
