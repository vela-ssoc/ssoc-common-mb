package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NewClient(clis ...*http.Client) Client {
	cli := Client{cli: http.DefaultClient}
	if len(clis) != 0 {
		cli.cli = clis[0]
	}

	return cli
}

// Client http 客户端。
type Client struct {
	cli *http.Client
}

// JSON 发送 GET 请求，响应 JSON 数据。
func (c Client) JSON(ctx context.Context, rawURL string, header http.Header, result any) error {
	resp, err := c.sendJSON(ctx, http.MethodGet, rawURL, header, http.NoBody)
	if err != nil {
		return err
	}
	err = c.unmarshalJSON(resp.Body, result)

	return err
}

// PostJSON 通过 POST 发送 JSON 报文，响应 JSON 数据。
func (c Client) PostJSON(ctx context.Context, rawURL string, header http.Header, body, result any) error {
	resp, err := c.sendJSON(ctx, http.MethodPost, rawURL, header, body)
	if err != nil {
		return err
	}
	err = c.unmarshalJSON(resp.Body, result)

	return err
}

// PostForm POST FromData 数据，注意只是普通的 FormData，不支持发送二进制文件。
func (c Client) PostForm(ctx context.Context, rawURL string, header http.Header, body url.Values, result any) error {
	if header == nil {
		header = make(http.Header, 4)
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	encode := body.Encode()
	req, err := c.newRequest(ctx, http.MethodPost, rawURL, header, strings.NewReader(encode))
	if err != nil {
		return err
	}

	resp, err := c.send(req)
	if err != nil {
		return err
	}

	return c.unmarshalJSON(resp.Body, result)
}

// Do 发送请求。
func (c Client) Do(req *http.Request) (*http.Response, error) {
	return c.send(req)
}

func (c Client) sendJSON(ctx context.Context, method, rawURL string, header http.Header, body any) (*http.Response, error) {
	if header == nil {
		header = make(http.Header, 4)
	}
	header.Set("Accept", "application/json")

	var r io.Reader
	if method != http.MethodGet && method != http.MethodHead {
		rd, err := c.marshalJSON(body)
		if err != nil {
			return nil, err
		}
		r = rd
		header.Set("Content-Type", "application/json; charset=utf-8")
	}

	req, err := c.newRequest(ctx, method, rawURL, header, r)
	if err != nil {
		return nil, err
	}

	return c.send(req)
}

func (Client) newRequest(ctx context.Context, method, rawURL string, header http.Header, body io.Reader) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, method, rawURL, body)
	if err != nil {
		return nil, err
	}
	if len(header) != 0 {
		req.Header = header
	}

	return req, nil
}

func (c Client) send(req *http.Request) (*http.Response, error) {
	h := req.Header
	if host := h.Get("Host"); host != "" {
		req.Host = host
	}
	if h.Get("User-Agent") == "" {
		chrome133 := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
		h.Set("User-Agent", chrome133)
	}
	if h.Get("Accept-Language") == "" {
		h.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	}

	resp, err := c.getClient().Do(req)
	if err != nil {
		return nil, err
	}

	code := resp.StatusCode
	rem := code / 100
	if rem == 2 || rem == 3 {
		return resp, nil
	}

	e := &Error{
		Code:    code,
		Header:  resp.Header,
		Request: req,
	}
	buf := make([]byte, 4096)
	n, _ := io.ReadFull(resp.Body, buf)
	_ = resp.Body.Close()
	e.Body = buf[:n]

	return nil, e
}

func (c Client) marshalJSON(v any) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	return buf, err
}

func (c Client) getClient() *http.Client {
	if cli := c.cli; cli != nil {
		return cli
	}
	return http.DefaultClient
}

func (c Client) unmarshalJSON(rc io.ReadCloser, v any) error {
	//goland:noinspection GoUnhandledErrorResult
	defer rc.Close()
	if v == nil {
		_, err := io.Copy(io.Discard, rc)
		return err
	}

	return json.NewDecoder(rc).Decode(v)
}
