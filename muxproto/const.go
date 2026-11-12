package muxproto

import (
	"errors"
	"net/url"
	"strconv"
)

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

func BuildAgentURL(host, path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   BuildAgentFQDN(host),
		Path:   path,
	}
}

func BuildBrokerURL(host, path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   BuildBrokerFQDN(host),
		Path:   path,
	}
}

func BuildManagerURL(path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   BuildManagerFQDN(),
		Path:   path,
	}
}

func BuildFQDNURL(fqdn, path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   fqdn,
		Path:   path,
	}
}

func Int64ToHost(i int64) string {
	return strconv.FormatInt(i, 10)
}
