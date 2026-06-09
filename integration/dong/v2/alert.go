package dong

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AlertConfig struct {
	SIEM    bool   `json:"siem"`
	URL     string `json:"url"`
	Token   string `json:"token"`
	Account string `json:"account"`
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

func NewAlert(cfg AlertConfigurer, log *slog.Logger, clis ...*http.Client) Client {
	ac := &alertClient{cfg: cfg, log: log}
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
	log *slog.Logger
}

func (ac *alertClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	cfg, err := ac.cfg.AlertConfigure(ctx)
	if err != nil {
		return err
	}
	ac.log.Info("发送告警配置信息", "config", cfg)
	if cfg.SIEM {
		return ac.sendSIEM(ctx, cfg, uids, gids, title, body)
	}

	return ac.sendDong(ctx, cfg, uids, gids, title, body)
}

func (ac *alertClient) sendSIEM(ctx context.Context, cfg *AlertConfig, uids, gids []string, title, body string) error {
	var ruleID string
	if len(gids) != 0 {
		ruleID = gids[0]
	}
	data := &requestBody{
		Prop: requestProp{JobNumbers: uids, RuleID: ruleID},
		Data: requestData{Detail: body, Title: title},
	}
	header := http.Header{"Authorization": []string{cfg.Token}}
	rawURL, err := cfg.joinPath("/api/dong/send/adaptive")
	if err != nil {
		return err
	}
	strURL := rawURL.String()

	res, err := ac.cli.postJSON(ctx, strURL, data, header)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode/100 == 2 {
		return nil
	}

	buf := make([]byte, 1024)
	n, err := io.ReadFull(res.Body, buf)
	if err != nil {
		return err
	}

	return errors.New(string(buf[:n]))
}

func (ac *alertClient) sendDong(ctx context.Context, cfg *AlertConfig, uids, gids []string, title, body string) error {
	rawURL, err := cfg.joinPath("/oa-chat-api/api/msg/broadcast/")
	if err != nil {
		return err
	}
	strURL := rawURL.String()

	data := &requestDirectBody{
		UserIDs: strings.Join(uids, ","),
		Title:   title,
		Detail:  body,
	}
	header := http.Header{"Token": []string{cfg.Token}, "Account": []string{cfg.Account}}

	res, err := ac.cli.postJSON(ctx, strURL, data, header)
	if err != nil {
		return err
	}

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
