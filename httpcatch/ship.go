package httpcatch

import "github.com/xgfone/ship/v5"

func NotFound(c *ship.Context) error {
	return ship.ErrNotFound.Newf("资源或数据不存在")
}

func HandleShipError(c *ship.Context, e error) {
	w, r := c.Response(), c.Request()
	if c.Response().Wrote {
		return
	}

	HandleError(w, r, e)
}
