package dong

import (
	"context"
	"net/http"
)

func NewTunnel(cli *http.Client) Client {
	return &tunnelClient{
		cli: newCommonClient(cli),
	}
}

type tunnelClient struct {
	cli *commonClient
}

func (tc *tunnelClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	strURL := "http://vtun/alert/dong"
	return tc.cli.send(ctx, strURL, nil, uids, gids, title, body)
}
