package dong

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type DirectConfig struct {
	URL     string
	Account string
	Token   string
}

func (dc DirectConfig) joinPath(str string) (*url.URL, error) {
	pu, err := url.Parse(dc.URL)
	if err != nil {
		return nil, err
	}

	return pu.JoinPath(str), nil
}

type DirectConfigurer interface {
	DirectConfigure(ctx context.Context) (*DirectConfig, error)
}

func NewDirect(cfg DirectConfigurer, clis ...*http.Client) Client {
	dc := &directClient{cfg: cfg}
	if len(clis) != 0 && clis[0] != nil {
		dc.cli = newCommonClient(clis[0])
	} else {
		dc.cli = newCommonClient(http.DefaultClient)
	}

	return dc
}

type directClient struct {
	cfg DirectConfigurer
	cli *commonClient
}

func (dc *directClient) Send(ctx context.Context, uids, gids []string, title, detail string) error {
	cfg, err := dc.cfg.DirectConfigure(ctx)
	if err != nil {
		return err
	}
	rawURL, err := cfg.joinPath("/oa-chat-api/api/msg/broadcast/")
	if err != nil {
		return err
	}
	strURL := rawURL.String()

	data := &requestDirectBody{
		UserIDs:  strings.Join(uids, ","),
		GroupIDs: strings.Join(gids, ","),
		Title:    title,
		Detail:   detail,
	}
	header := http.Header{"Token": []string{cfg.Token}, "Account": []string{cfg.Account}}

	return dc.cli.postJSON(ctx, strURL, data, header)
}
