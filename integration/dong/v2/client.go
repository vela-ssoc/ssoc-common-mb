package dong

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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

	res, err := cc.cli.Do(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	if code := res.StatusCode; code/100 != 2 { // 2xx
		part := make([]byte, 0, 1024)
		n, _ := res.Body.Read(part)
		msg := string(part[:n])
		if msg == "" {
			msg = "咚咚服务器 HTTP 响应状态异常"
		}

		return &ResponseEntity{
			Code: strconv.Itoa(code),
			Msg:  msg,
		}
	}

	reply := new(BroadcastResponse)
	resp := &ResponseEntity{Data: reply}
	if err = json.NewDecoder(res.Body).Decode(resp); err != nil {
		return err
	}

	return resp.checkError()
}
