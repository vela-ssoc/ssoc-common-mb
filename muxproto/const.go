package muxproto

import "errors"

const (
	ManagerTLD = "manager.internal"
	BrokerTLD  = "broker.internal"
	AgentTLD   = "agent.internal"

	HeaderXNodeIdentify = "X-Node-Identify"
)

var ErrNodeUnavailable = errors.New("节点不存在或已离线")

func BuildManagerFQDN() string {
	return BuildFQDN("main", ManagerTLD) // 只有一个中心端，主机位就暂且固定为 main
}

func BuildBrokerFQDN(host string) string {
	return BuildFQDN(host, BrokerTLD)
}

func BuildAgentFQDN(host string) string {
	return BuildFQDN(host, AgentTLD)
}

func BuildFQDN(host, tld string) string {
	return host + "." + tld
}
