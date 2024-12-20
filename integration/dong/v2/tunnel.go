package dong

import (
	"context"
	"net/http"

	"github.com/vela-ssoc/vela-common-mb/logback"
)

func NewTunnel(cli *http.Client, log logback.Logger) Client {
	return &tunnelClient{
		cli: newCommonClient(cli),
		log: log,
	}
}

type tunnelClient struct {
	cli *commonClient
	log logback.Logger
}

func (tc *tunnelClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	tc.log.Info("通过tunnel发送告警")
	strURL := "http://vtun/alert/dong"
	data := &requestTunnelBody{
		UserIDs:  uids,
		GroupIDs: gids,
		Detail:   body,
		Title:    title,
	}
	err := tc.cli.postJSON(ctx, strURL, data, nil)
	if err != nil {
		tc.log.Errorf("通过tunnel告警出错: %s", err)
	}

	return err
}
