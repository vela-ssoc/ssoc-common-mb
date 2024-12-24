package vela_common_mb

import (
	"github.com/go-playground/validator/v10"
	"github.com/vela-ssoc/opengauss"
	"github.com/vela-ssoc/vela-common-mba/netutil"
	"github.com/xgfone/ship/v5"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	_ *gen.Dao
	_ *gorm.DB
	_ *zap.Logger
	_ *validator.Validate
	_ *ship.Ship
	_ netutil.HTTPError
	_ opengauss.Config
)
