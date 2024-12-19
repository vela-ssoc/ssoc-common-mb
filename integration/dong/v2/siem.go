package dong

import (
	"context"
	"net/http"
	"net/url"
)

type SIEMConfig struct {
	URL   string // SIEM 服务器地址
	Token string // SIEM 服务调用 Token
}

func (sc SIEMConfig) SIEMConfigure(ctx context.Context) (*SIEMConfig, error) {
	return &sc, nil
}

func (sc SIEMConfig) joinPath(str string) (*url.URL, error) {
	pu, err := url.Parse(sc.URL)
	if err != nil {
		return nil, err
	}

	return pu.JoinPath(str), nil
}

type SIEMConfigurer interface {
	SIEMConfigure(ctx context.Context) (*SIEMConfig, error)
}

// NewSIEM 通过 SIEM 平台发送告警。
//
// https://siem.eastmoney.com/oas3/stoplight.html#/operations/post-dong-send-adaptive
func NewSIEM(cfg SIEMConfigurer, clis ...*http.Client) Client {
	sc := &siemClient{cfg: cfg}
	if len(clis) != 0 && clis[0] != nil {
		sc.cli = newCommonClient(clis[0])
	} else {
		sc.cli = newCommonClient(http.DefaultClient)
	}

	return sc
}

type siemClient struct {
	cfg SIEMConfigurer
	cli *commonClient
}

func (sc *siemClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	cfg, err := sc.cfg.SIEMConfigure(ctx)
	if err != nil {
		return err
	}
	rawURL, err := cfg.joinPath("/api/dong/send/adaptive")
	if err != nil {
		return err
	}
	strURL := rawURL.String()
	header := http.Header{"Authorization": []string{cfg.Token}}

	return sc.cli.send(ctx, strURL, header, uids, gids, title, body)
}
