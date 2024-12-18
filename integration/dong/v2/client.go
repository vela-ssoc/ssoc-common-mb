package dong

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client interface {
	Send(ctx context.Context, uids, gids []string, title, body string) error
}

func newCommonClient(cli *http.Client) *commonClient {
	return &commonClient{cli: cli}
}

type commonClient struct {
	cli *http.Client
}

func (cc *commonClient) send(ctx context.Context, strURL string, header http.Header, uids, gids []string, title, body string) error {
	data := &requestBody{
		Prop: requestProp{JobNumbers: uids},
		Data: requestData{Detail: body, Title: title},
	}

	return cc.postJSON(ctx, strURL, data, header)
}

func (cc *commonClient) postJSON(ctx context.Context, strURL string, data any, header http.Header) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strURL, buf)
	if err != nil {
		return err
	}
	if len(header) != 0 {
		req.Header = header
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := cc.cli.Do(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	if resp.StatusCode/100 == 2 {
		return nil
	}

	msg := make([]byte, 1024)
	n, _ := io.ReadFull(resp.Body, msg)

	return errors.New(string(msg[:n]))
}
