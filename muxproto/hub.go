package muxproto

import (
	"sync"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type Huber interface {
	Put(host string, mux muxconn.Muxer, data any) (newly Node)
	Get(host string) Node
	Del(host string) (last Node)
	All() []Node
	TLD() string
}

type SafeMap struct {
	tld   string // 一般固定 broker.internal
	mutex sync.RWMutex
	nodes map[string]Node // KEY: host
}

func NewBHub() *SafeMap {
	return &SafeMap{
		tld:   BrokerTLD,
		nodes: make(map[string]Node, 32),
	}
}

func (m *SafeMap) Put(host string, mux muxconn.Muxer, data any) Node {
	node := m.newNode(host, mux, data)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if val := m.nodes[host]; val != nil {
		return nil
	}
	m.nodes[host] = node

	return node
}

func (m *SafeMap) Get(host string) Node {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if node := m.nodes[host]; node != nil {
		return node
	}

	return nil
}

func (m *SafeMap) Del(host string) Node {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if node := m.nodes[host]; node != nil {
		delete(m.nodes, host)
		return node
	}

	return nil
}

func (m *SafeMap) All() []Node {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	nodes := make([]Node, 0, len(m.nodes))
	for _, node := range m.nodes {
		nodes = append(nodes, node)
	}

	return nodes
}

func (m *SafeMap) TLD() string {
	return m.tld
}

func (m *SafeMap) newNode(host string, mux muxconn.Muxer, data any) *virtualNode {
	fqdn := BuildFQDN(host, m.tld)
	return &virtualNode{
		mux:  mux,
		host: host,
		fqdn: fqdn,
		data: data,
	}
}
