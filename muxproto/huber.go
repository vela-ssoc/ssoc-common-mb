package muxproto

import (
	"iter"
	"slices"
	"sync"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type Huber interface {
	Put(slug string, mux muxconn.Muxer, data any) Peer
	Get(slug string) Peer
	Del(slug string) Peer
	All() iter.Seq[Peer]
	TLD() string
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

func (sm *safeMap) TLD() string {
	return sm.domain
}

func (sm *safeMap) Put(slug string, mux muxconn.Muxer, data any) Peer {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if p := sm.peers[slug]; p != nil {
		return nil
	}

	p := sm.newPeer(slug, mux, data)
	sm.peers[slug] = p

	return p
}

func (sm *safeMap) Get(slug string) Peer {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	if p := sm.peers[slug]; p != nil {
		return p
	}

	return nil
}

func (sm *safeMap) Del(slug string) Peer {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if p := sm.peers[slug]; p != nil {
		delete(sm.peers, slug)
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

func (sm *safeMap) newPeer(slug string, mux muxconn.Muxer, data any) *peer {
	fqdn := slug + "." + sm.domain

	return &peer{
		slug: slug,
		mux:  mux,
		fqdn: fqdn,
		data: data,
	}
}
