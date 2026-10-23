package muxproto

import (
	"context"
	"net"
	"strings"

	"github.com/vela-ssoc/vela-tunnel/muxconn"
)

type ThirdDialer struct {
	host string
	mux  muxconn.Muxer
	hub  Huber
	next net.Dialer
}

// NewDualDialer 主要用于 Manager 节点中，Manager 需要能够和 Broker、外部网络进行通信。
func NewDualDialer(hub Huber, next net.Dialer) *ThirdDialer {
	return NewThirdDialer("", nil, hub, next)
}

// NewThirdDialer 主要用于 Broker 节点中，Broker 需要能够和 Manager、Agent、外部网络进行通信。
func NewThirdDialer(host string, mux muxconn.Muxer, hub Huber, next net.Dialer) *ThirdDialer {
	return &ThirdDialer{
		host: host,
		mux:  mux,
		hub:  hub,
		next: next,
	}
}

func (td *ThirdDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if (td.host == "" || td.mux == nil) && td.hub == nil {
		return td.next.DialContext(ctx, network, address)
	}

	fqdn, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	fqdn = strings.TrimSuffix(fqdn, ".") // example.com. 也是合法域名

	if td.host == "" && td.host == fqdn && td.mux != nil {
		return td.mux.Open(ctx)
	}

	tld := td.hub.TLD()
	host, found := strings.CutSuffix(fqdn, "."+tld)
	if !found {
		return td.next.DialContext(ctx, network, address)
	}

	if p := td.hub.Get(host); p != nil {
		m := p.MUX()
		return m.Open(ctx)
	}

	return nil, &net.OpError{
		Op:   "open-substream",
		Net:  network,
		Addr: &net.UnixAddr{Name: fqdn, Net: "tunnel"},
		Err:  ErrNodeUnavailable,
	}
}
