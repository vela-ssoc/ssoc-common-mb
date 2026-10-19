package polydb

import (
	"github.com/vela-ssoc/ssoc-common-mb/datalayer/repository"
	"gorm.io/gorm"
)

type Sourcer interface {
	// DB MySQL/OpenGauss DB.
	DB() *gorm.DB

	// MDB MongoDB.
	MDB() *repository.BaseDB
}

func NewSource(db *gorm.DB, mdb *repository.BaseDB) Sourcer {
	return &polyDB{db: db, mdb: mdb}
}

type polyDB struct {
	db  *gorm.DB
	mdb *repository.BaseDB
}

func (p *polyDB) DB() *gorm.DB            { return p.db }
func (p *polyDB) MDB() *repository.BaseDB { return p.mdb }
