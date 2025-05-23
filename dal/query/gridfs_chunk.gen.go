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

	"github.com/vela-ssoc/ssoc-common-mb/dal/model"
)

func newGridChunk(db *gorm.DB, opts ...gen.DOOption) gridChunk {
	_gridChunk := gridChunk{}

	_gridChunk.gridChunkDo.UseDB(db, opts...)
	_gridChunk.gridChunkDo.UseModel(&model.GridChunk{})

	tableName := _gridChunk.gridChunkDo.TableName()
	_gridChunk.ALL = field.NewAsterisk(tableName)
	_gridChunk.ID = field.NewInt64(tableName, "id")
	_gridChunk.FileID = field.NewInt64(tableName, "file_id")
	_gridChunk.Serial = field.NewInt(tableName, "serial")
	_gridChunk.Data = field.NewBytes(tableName, "data")

	_gridChunk.fillFieldMap()

	return _gridChunk
}

type gridChunk struct {
	gridChunkDo gridChunkDo

	ALL    field.Asterisk
	ID     field.Int64 // ID
	FileID field.Int64 // 文件ID
	Serial field.Int   // 文件分片序号
	Data   field.Bytes // 文件内容分片

	fieldMap map[string]field.Expr
}

func (g gridChunk) Table(newTableName string) *gridChunk {
	g.gridChunkDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gridChunk) As(alias string) *gridChunk {
	g.gridChunkDo.DO = *(g.gridChunkDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gridChunk) updateTableName(table string) *gridChunk {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.FileID = field.NewInt64(table, "file_id")
	g.Serial = field.NewInt(table, "serial")
	g.Data = field.NewBytes(table, "data")

	g.fillFieldMap()

	return g
}

func (g *gridChunk) WithContext(ctx context.Context) *gridChunkDo {
	return g.gridChunkDo.WithContext(ctx)
}

func (g gridChunk) TableName() string { return g.gridChunkDo.TableName() }

func (g gridChunk) Alias() string { return g.gridChunkDo.Alias() }

func (g gridChunk) Columns(cols ...field.Expr) gen.Columns { return g.gridChunkDo.Columns(cols...) }

func (g *gridChunk) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gridChunk) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["id"] = g.ID
	g.fieldMap["file_id"] = g.FileID
	g.fieldMap["serial"] = g.Serial
	g.fieldMap["data"] = g.Data
}

func (g gridChunk) clone(db *gorm.DB) gridChunk {
	g.gridChunkDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gridChunk) replaceDB(db *gorm.DB) gridChunk {
	g.gridChunkDo.ReplaceDB(db)
	return g
}

type gridChunkDo struct{ gen.DO }

func (g gridChunkDo) Debug() *gridChunkDo {
	return g.withDO(g.DO.Debug())
}

func (g gridChunkDo) WithContext(ctx context.Context) *gridChunkDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gridChunkDo) ReadDB() *gridChunkDo {
	return g.Clauses(dbresolver.Read)
}

func (g gridChunkDo) WriteDB() *gridChunkDo {
	return g.Clauses(dbresolver.Write)
}

func (g gridChunkDo) Session(config *gorm.Session) *gridChunkDo {
	return g.withDO(g.DO.Session(config))
}

func (g gridChunkDo) Clauses(conds ...clause.Expression) *gridChunkDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gridChunkDo) Returning(value interface{}, columns ...string) *gridChunkDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gridChunkDo) Not(conds ...gen.Condition) *gridChunkDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gridChunkDo) Or(conds ...gen.Condition) *gridChunkDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gridChunkDo) Select(conds ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gridChunkDo) Where(conds ...gen.Condition) *gridChunkDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gridChunkDo) Order(conds ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gridChunkDo) Distinct(cols ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gridChunkDo) Omit(cols ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gridChunkDo) Join(table schema.Tabler, on ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gridChunkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gridChunkDo) RightJoin(table schema.Tabler, on ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gridChunkDo) Group(cols ...field.Expr) *gridChunkDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gridChunkDo) Having(conds ...gen.Condition) *gridChunkDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gridChunkDo) Limit(limit int) *gridChunkDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gridChunkDo) Offset(offset int) *gridChunkDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gridChunkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gridChunkDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gridChunkDo) Unscoped() *gridChunkDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gridChunkDo) Create(values ...*model.GridChunk) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gridChunkDo) CreateInBatches(values []*model.GridChunk, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gridChunkDo) Save(values ...*model.GridChunk) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gridChunkDo) First() (*model.GridChunk, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridChunk), nil
	}
}

func (g gridChunkDo) Take() (*model.GridChunk, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridChunk), nil
	}
}

func (g gridChunkDo) Last() (*model.GridChunk, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridChunk), nil
	}
}

func (g gridChunkDo) Find() ([]*model.GridChunk, error) {
	result, err := g.DO.Find()
	return result.([]*model.GridChunk), err
}

func (g gridChunkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GridChunk, err error) {
	buf := make([]*model.GridChunk, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gridChunkDo) FindInBatches(result *[]*model.GridChunk, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gridChunkDo) Attrs(attrs ...field.AssignExpr) *gridChunkDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gridChunkDo) Assign(attrs ...field.AssignExpr) *gridChunkDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gridChunkDo) Joins(fields ...field.RelationField) *gridChunkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gridChunkDo) Preload(fields ...field.RelationField) *gridChunkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gridChunkDo) FirstOrInit() (*model.GridChunk, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridChunk), nil
	}
}

func (g gridChunkDo) FirstOrCreate() (*model.GridChunk, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GridChunk), nil
	}
}

func (g gridChunkDo) FindByPage(offset int, limit int) (result []*model.GridChunk, count int64, err error) {
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

func (g gridChunkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gridChunkDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gridChunkDo) Delete(models ...*model.GridChunk) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gridChunkDo) withDO(do gen.Dao) *gridChunkDo {
	g.DO = *do.(*gen.DO)
	return g
}
