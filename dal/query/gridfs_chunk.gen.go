// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
)

func newGridfsChunk(db *gorm.DB, opts ...gen.DOOption) gridfsChunk {
	_gridfsChunk := gridfsChunk{}

	_gridfsChunk.gridfsChunkDo.UseDB(db, opts...)
	_gridfsChunk.gridfsChunkDo.UseModel(&model.GridfsChunk{})

	tableName := _gridfsChunk.gridfsChunkDo.TableName()
	_gridfsChunk.ALL = field.NewAsterisk(tableName)
	_gridfsChunk.ID = field.NewInt64(tableName, "id")
	_gridfsChunk.FileID = field.NewInt64(tableName, "file_id")
	_gridfsChunk.Serial = field.NewInt(tableName, "serial")
	_gridfsChunk.Data = field.NewBytes(tableName, "data")

	_gridfsChunk.fillFieldMap()

	return _gridfsChunk
}

type gridfsChunk struct {
	gridfsChunkDo gridfsChunkDo

	ALL    field.Asterisk
	ID     field.Int64
	FileID field.Int64
	Serial field.Int
	Data   field.Bytes

	fieldMap map[string]field.Expr
}

func (g gridfsChunk) Table(newTableName string) *gridfsChunk {
	g.gridfsChunkDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gridfsChunk) As(alias string) *gridfsChunk {
	g.gridfsChunkDo.DO = *(g.gridfsChunkDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gridfsChunk) updateTableName(table string) *gridfsChunk {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.FileID = field.NewInt64(table, "file_id")
	g.Serial = field.NewInt(table, "serial")
	g.Data = field.NewBytes(table, "data")

	g.fillFieldMap()

	return g
}

func (g *gridfsChunk) WithContext(ctx context.Context) *gridfsChunkDo {
	return g.gridfsChunkDo.WithContext(ctx)
}

func (g gridfsChunk) TableName() string { return g.gridfsChunkDo.TableName() }

func (g gridfsChunk) Alias() string { return g.gridfsChunkDo.Alias() }

func (g gridfsChunk) Columns(cols ...field.Expr) gen.Columns { return g.gridfsChunkDo.Columns(cols...) }

func (g *gridfsChunk) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gridfsChunk) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["id"] = g.ID
	g.fieldMap["file_id"] = g.FileID
	g.fieldMap["serial"] = g.Serial
	g.fieldMap["data"] = g.Data
}

func (g gridfsChunk) clone(db *gorm.DB) gridfsChunk {
	g.gridfsChunkDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gridfsChunk) replaceDB(db *gorm.DB) gridfsChunk {
	g.gridfsChunkDo.ReplaceDB(db)
	return g
}

type gridfsChunkDo struct{ gen.DO }

func (g gridfsChunkDo) Debug() *gridfsChunkDo {
	return g.withDO(g.DO.Debug())
}

func (g gridfsChunkDo) WithContext(ctx context.Context) *gridfsChunkDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gridfsChunkDo) ReadDB() *gridfsChunkDo {
	return g.Clauses(dbresolver.Read)
}

func (g gridfsChunkDo) WriteDB() *gridfsChunkDo {
	return g.Clauses(dbresolver.Write)
}

func (g gridfsChunkDo) Session(config *gorm.Session) *gridfsChunkDo {
	return g.withDO(g.DO.Session(config))
}

func (g gridfsChunkDo) Clauses(conds ...clause.Expression) *gridfsChunkDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gridfsChunkDo) Returning(value interface{}, columns ...string) *gridfsChunkDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gridfsChunkDo) Not(conds ...gen.Condition) *gridfsChunkDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gridfsChunkDo) Or(conds ...gen.Condition) *gridfsChunkDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gridfsChunkDo) Select(conds ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gridfsChunkDo) Where(conds ...gen.Condition) *gridfsChunkDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gridfsChunkDo) Order(conds ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gridfsChunkDo) Distinct(cols ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gridfsChunkDo) Omit(cols ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gridfsChunkDo) Join(table schema.Tabler, on ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gridfsChunkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gridfsChunkDo) RightJoin(table schema.Tabler, on ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gridfsChunkDo) Group(cols ...field.Expr) *gridfsChunkDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gridfsChunkDo) Having(conds ...gen.Condition) *gridfsChunkDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gridfsChunkDo) Limit(limit int) *gridfsChunkDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gridfsChunkDo) Offset(offset int) *gridfsChunkDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gridfsChunkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gridfsChunkDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gridfsChunkDo) Unscoped() *gridfsChunkDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gridfsChunkDo) Create(values ...*model.GridfsChunk) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gridfsChunkDo) CreateInBatches(values []*model.GridfsChunk, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gridfsChunkDo) Save(values ...*model.GridfsChunk) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gridfsChunkDo) First() (*model.GridfsChunk, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridfsChunk), nil
	}
}

func (g gridfsChunkDo) Take() (*model.GridfsChunk, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridfsChunk), nil
	}
}

func (g gridfsChunkDo) Last() (*model.GridfsChunk, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridfsChunk), nil
	}
}

func (g gridfsChunkDo) Find() ([]*model.GridfsChunk, error) {
	result, err := g.DO.Find()
	return result.([]*model.GridfsChunk), err
}

func (g gridfsChunkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GridfsChunk, err error) {
	buf := make([]*model.GridfsChunk, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gridfsChunkDo) FindInBatches(result *[]*model.GridfsChunk, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gridfsChunkDo) Attrs(attrs ...field.AssignExpr) *gridfsChunkDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gridfsChunkDo) Assign(attrs ...field.AssignExpr) *gridfsChunkDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gridfsChunkDo) Joins(fields ...field.RelationField) *gridfsChunkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gridfsChunkDo) Preload(fields ...field.RelationField) *gridfsChunkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gridfsChunkDo) FirstOrInit() (*model.GridfsChunk, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridfsChunk), nil
	}
}

func (g gridfsChunkDo) FirstOrCreate() (*model.GridfsChunk, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridfsChunk), nil
	}
}

func (g gridfsChunkDo) FindByPage(offset int, limit int) (result []*model.GridfsChunk, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gridfsChunkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gridfsChunkDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gridfsChunkDo) Delete(models ...*model.GridfsChunk) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gridfsChunkDo) withDO(do gen.Dao) *gridfsChunkDo {
	g.DO = *do.(*gen.DO)
	return g
}
