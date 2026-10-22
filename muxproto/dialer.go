package muxproto

import (
	"context"
	"net"
	"strings"
)

type DualDialer struct {
	hub  Huber
	next net.Dialer
}

func NewDualDialer(hub Huber, next net.Dialer) *DualDialer {
	return &DualDialer{hub: hub, next: next}
}

func (dd *DualDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	fqdn, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	fqdn = strings.TrimSuffix(fqdn, ".")

	domain := dd.hub.TLD()
	host, found := strings.CutSuffix(fqdn, "."+domain)
	if !found {
		return dd.next.DialContext(ctx, network, address)
	}

	if p := dd.hub.Get(host); p != nil {
		m := p.MUX()
		return m.Open(ctx)
	}

	return nil, &net.OpError{
		Op:   "open-substream",
		Net:  network,
		Addr: &net.UnixAddr{Name: fqdn, Net: "tunnel"},
		Err:  ErrPeerUnavailable,
	}
}
