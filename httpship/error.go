package httpship

import (
	"github.com/xgfone/ship/v5"
)

func NotFound(*ship.Context) error {
	return ship.ErrNotFound.Newf("资源或数据不存在")
}

func HandleError(c *ship.Context, e error) {
	if c.IsResponded() {
		return
	}

	w, r := c.Response(), c.Request()
	ErrorHandle(w, r, e)
}
