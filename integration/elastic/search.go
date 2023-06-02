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
	Bulk(ctx context.Context, r io.Reader) (*BulkResponse, error)
}

func NewSearch(cfg SearchConfigurer, client netutil.HTTPClient) Searcher {
	ua := "elastic-ssoc-broker-" + runtime.GOOS + "-" + runtime.GOARCH
	return &esClient{
		cfg:    cfg,
		client: client,
		ua:     ua,
	}
}

type esClient struct {
	cfg    SearchConfigurer
	client netutil.HTTPClient
	ua     string
}

func (es *esClient) Bulk(parent context.Context, r io.Reader) (*BulkResponse, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
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

func (es *esClient) fetch(req *http.Request) (*response, error) {
	res, err := es.client.Do(req)
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

func (es *esClient) newRequest(ctx context.Context, method, path string, r io.Reader) (*http.Request, error) {
	addr, auth, err := es.cfg.Load()
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
