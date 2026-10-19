package muxproto

import (
	"iter"
	"slices"
	"sync"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Huber interface {
	Put(id bson.ObjectID, mux muxconn.Muxer, data any) Peer
	Get(id bson.ObjectID) Peer
	Del(id bson.ObjectID) Peer
	Lookup(host string) Peer
	All() iter.Seq[Peer]
}

type safeMap struct {
	domain string
	mutex  sync.RWMutex
	peers  map[string]*peer
}

func NewBHub() Huber {
	return &safeMap{
		domain: "broker.internal",
		peers:  make(map[string]*peer, 32),
	}
}

func (sm *safeMap) Put(id bson.ObjectID, mux muxconn.Muxer, data any) Peer {
	host := id.Hex()

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if p := sm.peers[host]; p != nil {
		return nil
	}

	p := sm.newPeer(id, mux, data)
	sm.peers[host] = p

	return p
}

func (sm *safeMap) Get(id bson.ObjectID) Peer {
	return sm.Lookup(id.Hex())
}

func (sm *safeMap) Del(id bson.ObjectID) Peer {
	host := id.Hex()

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if p := sm.peers[host]; p != nil {
		delete(sm.peers, host)
		return p
	}

	return nil
}

func (sm *safeMap) Lookup(host string) Peer {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	if p := sm.peers[host]; p != nil {
		return p
	}

	return nil
}

func (sm *safeMap) All() iter.Seq[Peer] {
	sm.mutex.Lock()
	peers := make([]Peer, 0, len(sm.peers))
	for _, p := range sm.peers {
		peers = append(peers, p)
	}
	sm.mutex.Unlock()

	return slices.Values(peers)
}

func (sm *safeMap) newPeer(id bson.ObjectID, mux muxconn.Muxer, data any) *peer {
	host := id.Hex()
	fqdn := host + "." + sm.domain

	return &peer{
		id:   id,
		mux:  mux,
		host: host,
		fqdn: fqdn,
		data: data,
	}
}
