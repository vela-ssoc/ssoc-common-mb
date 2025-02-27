package velacommonmb

import (
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"github.com/vela-ssoc/opengauss"
	"github.com/xgfone/ship/v5"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	_ *gen.Dao
	_ *gorm.DB
	_ *validator.Validate
	_ *ship.Ship
	_ opengauss.Config
	_ cron.Cron
)
