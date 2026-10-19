package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Broker struct {
	ID           bson.ObjectID   `bson:"_id,omitempty"           json:"id"`
	Slug         string          `bson:"slug"                    json:"slug"` // 唯一标识符
	Name         string          `bson:"name"                    json:"name"`
	Secret       string          `bson:"secret"                  json:"-"`
	Exposes      ExposeAddresses `bson:"exposes"                 json:"exposes"`
	Config       BrokerConfig    `bson:"config"                  json:"config"`
	Status       bool            `bson:"status"                  json:"status"`
	TunnelStat   *TunnelStat     `bson:"tunnel_stat,omitempty"   json:"tunnel_stat,omitzero"`
	ExecuteStat  *ExecuteStat    `bson:"execute_stat,omitempty"  json:"execute_stat,omitzero"`
	HeartbeatAt  time.Time       `bson:"heartbeat_at,omitempty"  json:"heartbeat_at,omitzero"`
	DisconnectAt time.Time       `bson:"disconnect_at,omitempty" json:"disconnect_at,omitzero"`
	CreatedAt    time.Time       `bson:"created_at,omitempty"    json:"created_at"`
	UpdatedAt    time.Time       `bson:"updated_at,omitempty"    json:"updated_at"`
}

func (Broker) CollectionInfo() (string, string) { return "broker", "中继节点" }

type ExposeAddress struct {
	Name string `bson:"name" json:"name" validate:"required"`
	Addr string `bson:"addr" json:"addr" validate:"required"`
}

type ExposeAddresses []*ExposeAddress

func (eas ExposeAddresses) Addresses() []string {
	rets := make([]string, 0, 10)
	uniq := make(map[string]struct{}, 8)
	for _, ea := range eas {
		addr := ea.Addr
		if _, exists := uniq[addr]; !exists {
			uniq[addr] = struct{}{}
			rets = append(rets, addr)
		}
	}

	return rets
}
