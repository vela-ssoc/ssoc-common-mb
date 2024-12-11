package dong

import (
	"context"

	"github.com/vela-ssoc/vela-common-mba/netutil"
)

func NewTunnel(cli netutil.HTTPClient) Client {
	return &tunnelClient{cli: cli}
}

type tunnelClient struct {
	cli netutil.HTTPClient
}

func (tc *tunnelClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	// TODO implement me
	panic("implement me")
}
