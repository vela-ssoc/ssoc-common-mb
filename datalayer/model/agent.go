package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Agent struct {
	ID            bson.ObjectID   `bson:"_id,omitempty"           json:"id"`            // 数据库 ID
	IDV1          int64           `bson:"idv1"                    json:"idv1"`          // V1 版数据库 ID
	MachineID     string          `bson:"machine_id"              json:"machine_id"`    // 机器码
	Status        AgentStatus     `bson:"status"                  json:"status"`        // 状态
	Tags          AgentTags       `bson:"tags"                    json:"tags"`          // 标签
	Broker        *BrokerEndpoint `bson:"broker,omitempty"        json:"broker"`        // 在哪台中继节点上
	Unload        bool            `bson:"unload"                  json:"unload"`        // 静默模式
	CMDB          *NodeCMDB       `bson:"cmdb,omitempty"          json:"cmdb,omitzero"` // CMDB 信息
	TunnelStat    *TunnelStat     `bson:"tunnel_stat,omitempty"   json:"tunnel_stat,omitzero"`
	ExecuteStat   *ExecuteStat    `bson:"execute_stat,omitempty"  json:"execute_stat,omitzero"`
	ActiveStreams int64           `bson:"active_streams"          json:"active_streams"`
	HeartbeatAt   time.Time       `bson:"heartbeat_at,omitempty"  json:"heartbeat_at,omitzero"`
	DisconnectAt  time.Time       `bson:"disconnect_at,omitempty" json:"disconnect_at,omitzero"`
	CreatedAt     time.Time       `bson:"created_at,omitempty"    json:"created_at"`
	UpdatedAt     time.Time       `bson:"updated_at,omitempty"    json:"updated_at"`
}

func (Agent) CollectionInfo() (string, string) { return "agent", "探针节点" }

type AgentStatus int

const (
	AgentStatusOffline AgentStatus = iota // 离线
	AgentStatusOnline                     // 在线
	AgentStatusDeleted                    // 标记删除
)

type BrokerEndpoint struct {
	ID   bson.ObjectID `bson:"id"   json:"id"`
	IDV1 int64         `bson:"idv1" json:"idv1"`
	Slug string        `bson:"slug" json:"slug"`
	Name string        `bson:"name" json:"name"`
}

type AgentTag struct {
	Type AgentTagType `bson:"type" json:"type"`
	Name string       `bson:"name" json:"name"`
}

type AgentTags []AgentTag

type AgentTagType int

const (
	AgentTagBuiltin AgentTagType = iota // 内置标签（不可编辑）
	AgentTagManual                      // 管理员人工标记的标签
	AgentTagAgent                       // 探针程序自主上报的标签
)
