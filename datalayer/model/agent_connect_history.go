package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AgentConnectHistory struct {
	ID             bson.ObjectID  `bson:"_id,omitempty"   json:"id"`
	AgentIDV1      int64          `bson:"agent_idv1"      json:"agent_idv1"`
	ExecuteStat    ExecuteStat    `bson:"execute_stat"    json:"execute_stat,omitzero"`
	TunnelStat     TunnelStat     `bson:"tunnel_stat"     json:"tunnel_stat"`
	Broker         BrokerEndpoint `bson:"broker"          json:"broker"`
	Seconds        int64          `bson:"seconds"         json:"seconds"`
	DisconnectedAt time.Time      `bson:"disconnected_at" json:"disconnected_at"`
}

func (AgentConnectHistory) CollectionInfo() (string, string) {
	return "agent_connect_history", "探针节点连接历史记录"
}
