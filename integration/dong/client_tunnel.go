package dong

import (
	"context"
	"log"
	"net/http"

	"github.com/vela-ssoc/vela-common-mba/netutil"
)

func NewTunnel(cli netutil.HTTPClient) Client {
	return &tunnelClient{cli: cli}
}

type tunnelClient struct {
	cli netutil.HTTPClient
}

func (tc *tunnelClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	req := &tunnelRequest{
		JobNumbers: uids,
		GroupIDs:   gids,
		Title:      title,
		Detail:     body,
	}
	res := new(struct{})

	err := tc.cli.JSON(ctx, http.MethodPost, "http://vtun/alert/dong", req, res, nil)
	if err != nil {
		log.Println(err)
	}

	return err
}

type tunnelRequest struct {
	JobNumbers []string `json:"job_numbers"`
	GroupIDs   []string `json:"group_ids"`
	Title      string   `json:"title"`
	Detail     string   `json:"detail"`
}
