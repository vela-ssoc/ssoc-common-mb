package muxproto

import (
	"github.com/vela-ssoc/vela-tunnel/muxconn"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Peer interface {
	ID() bson.ObjectID
	MUX() muxconn.Muxer
	Host() string
	FQDN() string
	Data() any
}

type peer struct {
	id   bson.ObjectID
	mux  muxconn.Muxer
	host string
	fqdn string
	data any
}

func (p *peer) ID() bson.ObjectID  { return p.id }
func (p *peer) MUX() muxconn.Muxer { return p.mux }
func (p *peer) Host() string       { return p.host }
func (p *peer) FQDN() string       { return p.fqdn }
func (p *peer) Data() any          { return p.data }
