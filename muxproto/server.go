package muxproto

import "github.com/vela-ssoc/vela-tunnel/muxconn"

type MUXServer interface {
	ServeMUX(mux muxconn.Muxer) error
}
