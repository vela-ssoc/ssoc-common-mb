package shiputil

import (
	"log/slog"

	"github.com/vela-ssoc/ssoc-common-mb/httpcatch"
	"github.com/vela-ssoc/ssoc-common-mb/logger"
	"github.com/xgfone/ship/v5"
)

func NewShip(v ship.Validator, h slog.Handler) *ship.Ship {
	sh := ship.Default()
	sh.Validator = v
	sh.NotFound = httpcatch.NotFound
	sh.HandleError = httpcatch.HandleShipError
	if h != nil {
		sh.Logger = logger.NewFormat(h, 4)
	}

	return sh
}
