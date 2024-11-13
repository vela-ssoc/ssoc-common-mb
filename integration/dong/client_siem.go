package dong

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vela-ssoc/vela-common-mba/netutil"
)

type SIEMConfig struct {
	URL   string `json:"url"   yaml:"url"`
	Token string `json:"token" yaml:"token"`
}

func NewSIEM(cli netutil.HTTPClient, siem SIEMConfig) Client {
	return &siemClient{
		siem: siem,
		cli:  cli,
	}
}

type siemClient struct {
	siem SIEMConfig
	cli  netutil.HTTPClient
}

func (sc *siemClient) Send(ctx context.Context, uids, ruleIDs []string, title, body string) error {
	var ruleID string
	if len(ruleIDs) > 0 {
		ruleID = ruleIDs[0]
	}

	if json.Valid([]byte(body)) {
		return sc.sendCard(ctx, uids, ruleID, json.RawMessage(body))
	}

	return sc.sendBroadcast(ctx, uids, ruleID, title, body)
}

func (sc *siemClient) sendCard(ctx context.Context, jobNumbers []string, ruleID string, body json.RawMessage) error {
	data := &siemDongCard{
		Prop: siemCardProp{
			JobNumbers: jobNumbers,
			RuleID:     ruleID,
		},
		Card: body,
	}

	ret := make(map[string]any, 4)
	header := http.Header{
		"Authorization": []string{sc.siem.Token},
	}
	addr := sc.siem.URL + "/dong/send/card"
	if err := sc.cli.JSON(ctx, http.MethodPost, addr, data, &ret, header); err != nil {
		return err
	}

	return nil
}

func (sc *siemClient) sendBroadcast(ctx context.Context, jobNumbers []string, ruleID string, title, body string) error {
	data := &siemDongBroadcastBody{
		Prop: siemCardProp{
			JobNumbers: jobNumbers,
			RuleID:     ruleID,
		},
		Broadcast: siemDongBroadcast{
			Title:  title,
			Detail: body,
		},
	}

	ret := make(map[string]any, 4)
	header := http.Header{
		"Authorization": []string{sc.siem.Token},
	}
	addr := sc.siem.URL + "/api/dong/send/broadcast"
	if err := sc.cli.JSON(ctx, http.MethodPost, addr, data, &ret, header); err != nil {
		return err
	}

	return nil
}

type siemDongCard struct {
	Prop siemCardProp    `json:"prop"`
	Card json.RawMessage `json:"card"`
}

type siemCardProp struct {
	JobNumbers []string `json:"job_numbers"`
	RuleID     string   `json:"rule_id"`
}

type siemDongBroadcast struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Detail string `json:"detail"`
}

type siemDongBroadcastBody struct {
	Prop      siemCardProp      `json:"prop"`
	Broadcast siemDongBroadcast `json:"broadcast"`
}
