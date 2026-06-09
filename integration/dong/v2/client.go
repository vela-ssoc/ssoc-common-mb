package dong

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Client interface {
	Send(ctx context.Context, uids, gids []string, title, body string) error
}

func newCommonClient(cli *http.Client, log *slog.Logger) *commonClient {
	return &commonClient{cli: cli, log: log}
}

type commonClient struct {
	cli *http.Client
	log *slog.Logger
}

func (cc *commonClient) send(ctx context.Context, strURL string, header http.Header, uids, gids []string, title, body string) (*http.Response, error) {
	data := &requestBody{
		Prop: requestProp{JobNumbers: uids},
		Data: requestData{Detail: body, Title: title},
	}

	return cc.postJSON(ctx, strURL, data, header)
}

func (cc *commonClient) postJSON(ctx context.Context, strURL string, data any, header http.Header) (*http.Response, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strURL, buf)
	if err != nil {
		return nil, err
	}
	if len(header) != 0 {
		req.Header = header
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return cc.cli.Do(req)
}
