package muxproto

import (
	"net/http"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type MUXServer interface {
	ServeMUX(mux muxconn.Muxer) error
}

type V1Accepter interface {
	AcceptTCP(w http.ResponseWriter, r *http.Request) error
}
