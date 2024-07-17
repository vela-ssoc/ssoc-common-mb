package dong

import (
	"context"
	"encoding/json"
)

func NewCompatible(htm Client, jsn Client) Client {
	return &compatibleClient{
		htm: htm,
		jsn: jsn,
	}
}

type compatibleClient struct {
	htm Client
	jsn Client
}

func (c *compatibleClient) Send(ctx context.Context, uids, gids []string, title, body string) error {
	if json.Valid([]byte(body)) {
		return c.jsn.Send(ctx, uids, gids, title, body)
	}
	return c.htm.Send(ctx, uids, gids, title, body)
}
