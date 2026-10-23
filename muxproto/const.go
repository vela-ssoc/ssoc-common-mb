package muxproto

const (
	ManagerFQDN = "manager.internal"
	BrokerTLD   = "broker.internal"
	AgentTLD    = "agent.internal"
)

func BuildBrokerFQDN(host string) string {
	return buildFQDN(host, BrokerTLD)
}

func BuildAgentFQDN(host string) string {
	return buildFQDN(host, AgentTLD)
}

func buildFQDN(host, tld string) string { return host + "." + tld }
