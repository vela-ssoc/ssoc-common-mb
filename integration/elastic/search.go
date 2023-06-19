package elastic

import (
	"context"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/vela-ssoc/vela-common-mba/netutil"
)

type Searcher interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	Bulk(ctx context.Context, rd io.Reader) (*BulkResponse, error)
}

func NewSearch(cfg Configurer, cli netutil.HTTPClient) Searcher {
	ua := "elastic-ssoc-broker-" + runtime.GOOS + "-" + runtime.GOARCH
	return &searchClient{
		cfg:     cfg,
		ua:      ua,
		cli:     cli,
		timeout: 5 * time.Second,
	}
}

type searchClient struct {
	cfg     Configurer
	ua      string
	cli     netutil.HTTPClient
	timeout time.Duration
}

func (es *searchClient) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	h, err := es.cfg.Load(ctx)
	if err == nil {
		h.ServeHTTP(w, r)
	}
	return err
}

func (es *searchClient) Bulk(parent context.Context, r io.Reader) (*BulkResponse, error) {
	ctx, cancel := context.WithTimeout(parent, es.timeout)
	defer cancel()

	req, err := es.newRequest(ctx, http.MethodPost, "/_bulk", r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-ndjson")

	res, err := es.fetch(req)
	if err != nil {
		return nil, err
	}
	ret := new(BulkResponse)
	err = res.Unmarshal(ret)

	return ret, err
}

func (es *searchClient) newRequest(ctx context.Context, method, path string, r io.Reader) (*http.Request, error) {
	addr, auth, err := es.cfg.LoadAddr(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, addr+path, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", es.ua)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (es *searchClient) fetch(req *http.Request) (*response, error) {
	res, err := es.cli.Do(req)
	if err != nil {
		return nil, err
	}
	ret := &response{
		StatusCode:          res.StatusCode,
		Header:              res.Header,
		DeprecationWarnings: res.Header["Warning"],
	}
	buf := make([]byte, 40960)
	n, _ := io.ReadFull(res.Body, buf)
	_ = res.Body.Close()
	ret.Body = buf[:n]

	return ret, nil
}
