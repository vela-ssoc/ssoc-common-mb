package muxproto

import (
	"iter"
	"slices"
	"sync"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type Huber interface {
	Put(slug string, mux muxconn.Muxer, data any) Node
	Get(slug string) Node
	Del(slug string) Node
	All() iter.Seq[Node]
	TLD() string
}

type safeMap struct {
	tld   string
	mutex sync.RWMutex
	nodes map[string]*virtualNode
}

func NewBHub() Huber {
	return &safeMap{
		tld:   BrokerTLD,
		nodes: make(map[string]*virtualNode, 32),
	}
}

func (sm *safeMap) TLD() string { return sm.tld }

func (sm *safeMap) Put(host string, mux muxconn.Muxer, data any) Node {
	node := sm.buildNode(mux, host, data)

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if val := sm.nodes[host]; val != nil {
		return nil
	}
	sm.nodes[host] = node

	return node
}

func (sm *safeMap) Get(host string) Node {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	node := sm.nodes[host]

	return node
}

func (sm *safeMap) Del(host string) Node {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if p := sm.nodes[host]; p != nil {
		delete(sm.nodes, host)
		return p
	}

	return nil
}

func (sm *safeMap) All() iter.Seq[Node] {
	sm.mutex.Lock()
	nodes := make([]Node, 0, len(sm.nodes))
	for _, p := range sm.nodes {
		nodes = append(nodes, p)
	}
	sm.mutex.Unlock()

	return slices.Values(nodes)
}

func (sm *safeMap) buildNode(mux muxconn.Muxer, host string, data any) *virtualNode {
	fqdn := buildFQDN(host, sm.tld)
	return &virtualNode{
		mux:  mux,
		host: host,
		fqdn: fqdn,
		data: data,
	}
}
