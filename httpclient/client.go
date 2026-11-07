package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	cli *http.Client
}

func NewClient(cli *http.Client) Client {
	return Client{cli: cli}
}

func (c Client) PostJSON(ctx context.Context, addr string, body, result any) error {
	var br io.Reader
	if body != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
		br = buf
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, addr, br)
	if err != nil {
		return err
	}

	resp, err := c.send(req)
	if err != nil {
		return err
	}

	return c.marshalJSON(resp, result)
}

func (c Client) Do(req *http.Request) (*http.Response, error) {
	return c.send(req)
}

func (c Client) send(req *http.Request) (*http.Response, error) {
	h := req.Header
	if host := h.Get("Host"); host != "" {
		req.Host = host
		h.Del("Host")
	}
	if h.Get("User-Agent") == "" {
		h.Set("User-Agent", "ssoc-http-client")
	}
	if h.Get("Accept-Language") == "" {
		h.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	}

	resp, err := c.getClient().Do(req)
	if err != nil {
		return nil, err
	}

	// 检查状态码
	flag := resp.StatusCode / 100
	if flag != 2 && flag != 3 {
		defer resp.Body.Close()
		err = c.catch(req, resp)

		return nil, err
	}

	return resp, nil
}

func (c Client) getClient() *http.Client {
	if c.cli != nil {
		return c.cli
	}

	return http.DefaultClient
}

// catch 无法解析为 RFC 9457 时，返回通用 HTTP 错误。
func (Client) catch(req *http.Request, resp *http.Response) error {
	// 安全起见至多读取 1K 数据，防止对端返回超大数据。
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))

	ret := &HTTPError{
		Code:     resp.StatusCode,
		Body:     body,
		Request:  req,
		Response: resp,
	}

	// 错误信息格式是 RFC 9457
	rfc := new(ProblemDetails)
	if err := json.Unmarshal(body, rfc); err == nil {
		ret.RFC9457 = rfc
	}

	return ret
}

func (Client) marshalJSON(resp *http.Response, result any) error {
	body := resp.Body
	defer body.Close()

	if result == nil || body == http.NoBody {
		return nil
	}

	return json.NewDecoder(body).Decode(result)
}
