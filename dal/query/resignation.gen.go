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

func newResignation(db *gorm.DB, opts ...gen.DOOption) resignation {
	_resignation := resignation{}

	_resignation.resignationDo.UseDB(db, opts...)
	_resignation.resignationDo.UseModel(&model.Resignation{})

	tableName := _resignation.resignationDo.TableName()
	_resignation.ALL = field.NewAsterisk(tableName)
	_resignation.ID = field.NewString(tableName, "id")
	_resignation.Name = field.NewString(tableName, "name")
	_resignation.CreatedAt = field.NewTime(tableName, "created_at")
	_resignation.UpdatedAt = field.NewTime(tableName, "updated_at")

	_resignation.fillFieldMap()

	return _resignation
}

type resignation struct {
	resignationDo resignationDo

	ALL       field.Asterisk
	ID        field.String
	Name      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (r resignation) Table(newTableName string) *resignation {
	r.resignationDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r resignation) As(alias string) *resignation {
	r.resignationDo.DO = *(r.resignationDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *resignation) updateTableName(table string) *resignation {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewString(table, "id")
	r.Name = field.NewString(table, "name")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *resignation) WithContext(ctx context.Context) *resignationDo {
	return r.resignationDo.WithContext(ctx)
}

func (r resignation) TableName() string { return r.resignationDo.TableName() }

func (r resignation) Alias() string { return r.resignationDo.Alias() }

func (r resignation) Columns(cols ...field.Expr) gen.Columns { return r.resignationDo.Columns(cols...) }

func (r *resignation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *resignation) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 4)
	r.fieldMap["id"] = r.ID
	r.fieldMap["name"] = r.Name
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r resignation) clone(db *gorm.DB) resignation {
	r.resignationDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r resignation) replaceDB(db *gorm.DB) resignation {
	r.resignationDo.ReplaceDB(db)
	return r
}

type resignationDo struct{ gen.DO }

func (r resignationDo) Debug() *resignationDo {
	return r.withDO(r.DO.Debug())
}

func (r resignationDo) WithContext(ctx context.Context) *resignationDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r resignationDo) ReadDB() *resignationDo {
	return r.Clauses(dbresolver.Read)
}

func (r resignationDo) WriteDB() *resignationDo {
	return r.Clauses(dbresolver.Write)
}

func (r resignationDo) Session(config *gorm.Session) *resignationDo {
	return r.withDO(r.DO.Session(config))
}

func (r resignationDo) Clauses(conds ...clause.Expression) *resignationDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r resignationDo) Returning(value interface{}, columns ...string) *resignationDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r resignationDo) Not(conds ...gen.Condition) *resignationDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r resignationDo) Or(conds ...gen.Condition) *resignationDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r resignationDo) Select(conds ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r resignationDo) Where(conds ...gen.Condition) *resignationDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r resignationDo) Order(conds ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r resignationDo) Distinct(cols ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r resignationDo) Omit(cols ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r resignationDo) Join(table schema.Tabler, on ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r resignationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *resignationDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r resignationDo) RightJoin(table schema.Tabler, on ...field.Expr) *resignationDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r resignationDo) Group(cols ...field.Expr) *resignationDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r resignationDo) Having(conds ...gen.Condition) *resignationDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r resignationDo) Limit(limit int) *resignationDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r resignationDo) Offset(offset int) *resignationDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r resignationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *resignationDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r resignationDo) Unscoped() *resignationDo {
	return r.withDO(r.DO.Unscoped())
}

func (r resignationDo) Create(values ...*model.Resignation) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r resignationDo) CreateInBatches(values []*model.Resignation, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r resignationDo) Save(values ...*model.Resignation) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r resignationDo) First() (*model.Resignation, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Resignation), nil
	}
}

func (r resignationDo) Take() (*model.Resignation, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Resignation), nil
	}
}

func (r resignationDo) Last() (*model.Resignation, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Resignation), nil
	}
}

func (r resignationDo) Find() ([]*model.Resignation, error) {
	result, err := r.DO.Find()
	return result.([]*model.Resignation), err
}

func (r resignationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Resignation, err error) {
	buf := make([]*model.Resignation, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r resignationDo) FindInBatches(result *[]*model.Resignation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r resignationDo) Attrs(attrs ...field.AssignExpr) *resignationDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r resignationDo) Assign(attrs ...field.AssignExpr) *resignationDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r resignationDo) Joins(fields ...field.RelationField) *resignationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r resignationDo) Preload(fields ...field.RelationField) *resignationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r resignationDo) FirstOrInit() (*model.Resignation, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Resignation), nil
	}
}

func (r resignationDo) FirstOrCreate() (*model.Resignation, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Resignation), nil
	}
}

func (r resignationDo) FindByPage(offset int, limit int) (result []*model.Resignation, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r resignationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r resignationDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r resignationDo) Delete(models ...*model.Resignation) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *resignationDo) withDO(do gen.Dao) *resignationDo {
	r.DO = *do.(*gen.DO)
	return r
}
