package dong

import (
	"context"
	"encoding/json"
	"log"
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
	data := &siemDongCard{
		Prop: siemCardProp{
			JobNumbers: uids,
			RuleID:     ruleID,
		},
		Card: json.RawMessage(body),
	}

	ret := make(map[string]any, 4)
	header := http.Header{
		"Authorization": []string{sc.siem.Token},
	}
	addr := sc.siem.URL + "/api/dong/card/send"
	if err := sc.cli.JSON(ctx, http.MethodPost, addr, data, &ret, header); err != nil {
		log.Println(err)
		log.Println(ret)
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
