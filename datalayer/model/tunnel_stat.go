package model

import "time"

type TunnelStat struct {
	Library       string    `bson:"library"                json:"library"`
	Module        string    `bson:"module"                 json:"module"`
	Inet          string    `bson:"inet,omitempty"         json:"inet,omitzero"`
	ConnectedAt   time.Time `bson:"connected_at,omitempty" json:"connected_at,omitzero"`
	Cumulative    int64     `bson:"cumulative"             json:"cumulative"`
	LocalAddr     string    `bson:"local_addr,omitempty"   json:"local_addr,omitzero"`
	RemoteAddr    string    `bson:"remote_addr,omitempty"  json:"remote_addr,omitzero"`
	ReceiveBytes  uint64    `bson:"receive_bytes"          json:"receive_bytes"`
	TransmitBytes uint64    `bson:"transmit_bytes"         json:"transmit_bytes"`
}
