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

func newDing(db *gorm.DB, opts ...gen.DOOption) ding {
	_ding := ding{}

	_ding.dingDo.UseDB(db, opts...)
	_ding.dingDo.UseModel(&model.Ding{})

	tableName := _ding.dingDo.TableName()
	_ding.ALL = field.NewAsterisk(tableName)
	_ding.ID = field.NewString(tableName, "id")
	_ding.Code = field.NewString(tableName, "code")
	_ding.Tries = field.NewInt(tableName, "tries")
	_ding.CreatedAt = field.NewTime(tableName, "created_at")

	_ding.fillFieldMap()

	return _ding
}

type ding struct {
	dingDo dingDo

	ALL       field.Asterisk
	ID        field.String
	Code      field.String
	Tries     field.Int
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (d ding) Table(newTableName string) *ding {
	d.dingDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d ding) As(alias string) *ding {
	d.dingDo.DO = *(d.dingDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *ding) updateTableName(table string) *ding {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Code = field.NewString(table, "code")
	d.Tries = field.NewInt(table, "tries")
	d.CreatedAt = field.NewTime(table, "created_at")

	d.fillFieldMap()

	return d
}

func (d *ding) WithContext(ctx context.Context) *dingDo { return d.dingDo.WithContext(ctx) }

func (d ding) TableName() string { return d.dingDo.TableName() }

func (d ding) Alias() string { return d.dingDo.Alias() }

func (d ding) Columns(cols ...field.Expr) gen.Columns { return d.dingDo.Columns(cols...) }

func (d *ding) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *ding) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 4)
	d.fieldMap["id"] = d.ID
	d.fieldMap["code"] = d.Code
	d.fieldMap["tries"] = d.Tries
	d.fieldMap["created_at"] = d.CreatedAt
}

func (d ding) clone(db *gorm.DB) ding {
	d.dingDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d ding) replaceDB(db *gorm.DB) ding {
	d.dingDo.ReplaceDB(db)
	return d
}

type dingDo struct{ gen.DO }

func (d dingDo) Debug() *dingDo {
	return d.withDO(d.DO.Debug())
}

func (d dingDo) WithContext(ctx context.Context) *dingDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dingDo) ReadDB() *dingDo {
	return d.Clauses(dbresolver.Read)
}

func (d dingDo) WriteDB() *dingDo {
	return d.Clauses(dbresolver.Write)
}

func (d dingDo) Session(config *gorm.Session) *dingDo {
	return d.withDO(d.DO.Session(config))
}

func (d dingDo) Clauses(conds ...clause.Expression) *dingDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dingDo) Returning(value interface{}, columns ...string) *dingDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dingDo) Not(conds ...gen.Condition) *dingDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dingDo) Or(conds ...gen.Condition) *dingDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dingDo) Select(conds ...field.Expr) *dingDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dingDo) Where(conds ...gen.Condition) *dingDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dingDo) Order(conds ...field.Expr) *dingDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dingDo) Distinct(cols ...field.Expr) *dingDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dingDo) Omit(cols ...field.Expr) *dingDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dingDo) Join(table schema.Tabler, on ...field.Expr) *dingDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *dingDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dingDo) RightJoin(table schema.Tabler, on ...field.Expr) *dingDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dingDo) Group(cols ...field.Expr) *dingDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dingDo) Having(conds ...gen.Condition) *dingDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dingDo) Limit(limit int) *dingDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dingDo) Offset(offset int) *dingDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *dingDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dingDo) Unscoped() *dingDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dingDo) Create(values ...*model.Ding) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dingDo) CreateInBatches(values []*model.Ding, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dingDo) Save(values ...*model.Ding) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dingDo) First() (*model.Ding, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ding), nil
	}
}

func (d dingDo) Take() (*model.Ding, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ding), nil
	}
}

func (d dingDo) Last() (*model.Ding, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ding), nil
	}
}

func (d dingDo) Find() ([]*model.Ding, error) {
	result, err := d.DO.Find()
	return result.([]*model.Ding), err
}

func (d dingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ding, err error) {
	buf := make([]*model.Ding, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dingDo) FindInBatches(result *[]*model.Ding, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dingDo) Attrs(attrs ...field.AssignExpr) *dingDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dingDo) Assign(attrs ...field.AssignExpr) *dingDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dingDo) Joins(fields ...field.RelationField) *dingDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dingDo) Preload(fields ...field.RelationField) *dingDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dingDo) FirstOrInit() (*model.Ding, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ding), nil
	}
}

func (d dingDo) FirstOrCreate() (*model.Ding, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ding), nil
	}
}

func (d dingDo) FindByPage(offset int, limit int) (result []*model.Ding, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dingDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dingDo) Delete(models ...*model.Ding) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dingDo) withDO(do gen.Dao) *dingDo {
	d.DO = *do.(*gen.DO)
	return d
}
