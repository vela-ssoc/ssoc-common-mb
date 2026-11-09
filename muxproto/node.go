package muxproto

import (
	"github.com/vela-ssoc/vela-tunnel/muxconn"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Node interface {
	MUX() muxconn.Muxer // 通信隧道
	Host() string       // 主机位，例如：test
	FQDN() string       // 全路径，例如：test.broker.internal
	Info() NodeInfo     // 附带信息
}

type virtualNode struct {
	mux  muxconn.Muxer
	host string
	fqdn string
	info NodeInfo
}

func (n *virtualNode) MUX() muxconn.Muxer { return n.mux }
func (n *virtualNode) Host() string       { return n.host }
func (n *virtualNode) FQDN() string       { return n.fqdn }
func (n *virtualNode) Info() NodeInfo     { return n.info }

type NodeInfo struct {
	ID      bson.ObjectID `json:"id"`   // 新版本 MongoDB 数据库 ID
	IDV1    int64         `json:"idv1"` // 兼容 V1 版本的数据
	Slug    string        `json:"slug"` // 仅中继节点有值
	Inet    string        `json:"inet"`
	GOOS    string        `json:"goos"`
	GOARCH  string        `json:"goarch"`
	Version string        `json:"version"`
}
