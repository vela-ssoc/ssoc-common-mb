package muxproto

import (
	"errors"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

var ErrPeerUnavailable = errors.New("节点不存在或已离线")

type Peer interface {
	MUX() muxconn.Muxer
	Slug() string
	FQDN() string
	Data() any
}

type peer struct {
	slug string
	mux  muxconn.Muxer
	host string
	fqdn string
	data any
}

func (p *peer) MUX() muxconn.Muxer { return p.mux }
func (p *peer) Slug() string       { return p.slug }
func (p *peer) FQDN() string       { return p.fqdn }
func (p *peer) Data() any          { return p.data }
