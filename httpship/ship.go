package httpship

import (
	"log/slog"

	"github.com/vela-ssoc/ssoc-common-mb/logger"
	"github.com/xgfone/ship/v5"
	"github.com/xgfone/ship/v5/middleware"
)

func NewDefault(v ship.Validator, h slog.Handler) *ship.Ship {
	sh := ship.Default()
	sh.Validator = v
	sh.NotFound = NotFound
	sh.HandleError = HandleError
	if h != nil {
		sh.Logger = logger.NewFormat(h, 4)
	}
	sh.Pre(middleware.Logger())

	return sh
}
