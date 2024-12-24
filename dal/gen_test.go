package dal_test

import (
	"testing"

	"github.com/vela-ssoc/vela-common-mb/dal/entity"
	"gorm.io/gen"
)

func TestGen(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithDefaultQuery,
		OutPath: "query",
	})

	tbls := entity.List()
	g.ApplyBasic(tbls...)

	g.Execute()
}
