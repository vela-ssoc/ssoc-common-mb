package muxproto

import (
	"context"
	"net"
	"strings"
)

type TLDDialer struct {
	dnss []TLDResolver
	back net.Dialer
}

func NewTLDDialer(back net.Dialer, dnss []TLDResolver) *TLDDialer {
	return &TLDDialer{dnss: dnss, back: back}
}

func (td *TLDDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if len(td.dnss) == 0 {
		return td.back.DialContext(ctx, network, address)
	}

	fqdn, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}
	fqdn = strings.TrimSuffix(fqdn, ".") // example.com. 也是合法域名
	for _, dns := range td.dnss {
		tld := strings.TrimSuffix(dns.TLD(), ".") // manager.internal.
		host, found := strings.CutSuffix(fqdn, "."+tld)
		if !found {
			continue
		}

		if mux := dns.MUX(host); mux != nil {
			return mux.Open(ctx)
		}

		return nil, &net.OpError{
			Op:   "open-substream",
			Net:  network,
			Addr: &net.UnixAddr{Name: fqdn, Net: "tunnel"},
			Err:  ErrNodeUnavailable,
		}
	}

	return td.back.DialContext(ctx, network, address)
}
