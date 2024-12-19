package dong

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type AlertConfig struct {
	SIEM    bool
	URL     string
	Token   string
	Account string
}

func (ac AlertConfig) joinPath(str string) (*url.URL, error) {
	pu, err := url.Parse(ac.URL)
	if err != nil {
		return nil, err
	}

	return pu.JoinPath(str), nil
}

type AlertConfigurer interface {
	AlertConfigure(ctx context.Context) (*AlertConfig, error)
}

func NewAlert(cfg AlertConfigurer, clis ...*http.Client) Client {
	ac := &alertClient{cfg: cfg}
	if len(clis) != 0 && clis[0] != nil {
		ac.cli = newCommonClient(clis[0])
	} else {
		ac.cli = newCommonClient(http.DefaultClient)
	}

	return ac
}

type alertClient struct {
	cli *commonClient
	cfg AlertConfigurer
}

func (ac *alertClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	cfg, err := ac.cfg.AlertConfigure(ctx)
	if err != nil {
		return err
	}
	if cfg.SIEM {
		return ac.sendSIEM(ctx, cfg, uids, gids, title, body)
	}

	return ac.sendDong(ctx, cfg, uids, gids, title, body)
}

func (ac *alertClient) sendSIEM(ctx context.Context, cfg *AlertConfig, uids, gids []string, title, body string) error {
	data := &requestBody{
		Prop: requestProp{JobNumbers: uids},
		Data: requestData{Detail: body, Title: title},
	}
	header := http.Header{"Authorization": []string{cfg.Token}}
	rawURL, err := cfg.joinPath("/api/dong/send/adaptive")
	if err != nil {
		return err
	}
	strURL := rawURL.String()

	return ac.cli.postJSON(ctx, strURL, data, header)
}

func (ac *alertClient) sendDong(ctx context.Context, cfg *AlertConfig, uids, gids []string, title, body string) error {
	rawURL, err := cfg.joinPath("/oa-chat-api/api/msg/broadcast/")
	if err != nil {
		return err
	}
	strURL := rawURL.String()

	data := &requestDirectBody{
		UserIDs:  strings.Join(uids, ","),
		GroupIDs: strings.Join(gids, ","),
		Title:    title,
		Detail:   body,
	}
	header := http.Header{"Token": []string{cfg.Token}, "Account": []string{cfg.Account}}

	return ac.cli.postJSON(ctx, strURL, data, header)
}
