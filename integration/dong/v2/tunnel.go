package dong

import (
	"context"
	"log/slog"
	"net/http"
)

func NewTunnel(cli *http.Client, log *slog.Logger) Client {
	return &tunnelClient{
		cli: newCommonClient(cli),
		log: log,
	}
}

type tunnelClient struct {
	cli *commonClient
	log *slog.Logger
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
		tc.log.Error("通过tunnel告警出错", slog.Any("error", err))
	}

	return err
}
