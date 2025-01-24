package dal_test

import (
	"testing"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"gorm.io/gen"
)

func TestGen(t *testing.T) {
	cfg := gen.Config{OutPath: "query"}
	g := gen.NewGenerator(cfg)
	g.ApplyBasic(model.All()...)
	g.Execute()
}
