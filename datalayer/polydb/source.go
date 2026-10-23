package polydb

import (
	"github.com/vela-ssoc/ssoc-common-mb/dal/query"
	"github.com/vela-ssoc/ssoc-common-mb/datalayer/repository"
)

type Sourcer interface {
	// MDB MongoDB.
	MDB() *repository.BaseDB

	// RDB 旧的关系型数据库（将被逐步淘汰）。
	RDB() *query.Query
}

func NewSource(mdb *repository.BaseDB, qry *query.Query) Sourcer {
	return &polyDB{mdb: mdb, qry: qry}
}

type polyDB struct {
	mdb *repository.BaseDB
	qry *query.Query
}

func (p *polyDB) MDB() *repository.BaseDB { return p.mdb }
func (p *polyDB) RDB() *query.Query       { return p.qry }
