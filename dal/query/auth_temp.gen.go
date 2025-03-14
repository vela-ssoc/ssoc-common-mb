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

func newAuthTemp(db *gorm.DB, opts ...gen.DOOption) authTemp {
	_authTemp := authTemp{}

	_authTemp.authTempDo.UseDB(db, opts...)
	_authTemp.authTempDo.UseModel(&model.AuthTemp{})

	tableName := _authTemp.authTempDo.TableName()
	_authTemp.ALL = field.NewAsterisk(tableName)
	_authTemp.ID = field.NewInt64(tableName, "id")
	_authTemp.UID = field.NewString(tableName, "uid")
	_authTemp.CreatedAt = field.NewTime(tableName, "created_at")

	_authTemp.fillFieldMap()

	return _authTemp
}

type authTemp struct {
	authTempDo authTempDo

	ALL       field.Asterisk
	ID        field.Int64
	UID       field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a authTemp) Table(newTableName string) *authTemp {
	a.authTempDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authTemp) As(alias string) *authTemp {
	a.authTempDo.DO = *(a.authTempDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authTemp) updateTableName(table string) *authTemp {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UID = field.NewString(table, "uid")
	a.CreatedAt = field.NewTime(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *authTemp) WithContext(ctx context.Context) *authTempDo { return a.authTempDo.WithContext(ctx) }

func (a authTemp) TableName() string { return a.authTempDo.TableName() }

func (a authTemp) Alias() string { return a.authTempDo.Alias() }

func (a authTemp) Columns(cols ...field.Expr) gen.Columns { return a.authTempDo.Columns(cols...) }

func (a *authTemp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authTemp) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["uid"] = a.UID
	a.fieldMap["created_at"] = a.CreatedAt
}

func (a authTemp) clone(db *gorm.DB) authTemp {
	a.authTempDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authTemp) replaceDB(db *gorm.DB) authTemp {
	a.authTempDo.ReplaceDB(db)
	return a
}

type authTempDo struct{ gen.DO }

func (a authTempDo) Debug() *authTempDo {
	return a.withDO(a.DO.Debug())
}

func (a authTempDo) WithContext(ctx context.Context) *authTempDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authTempDo) ReadDB() *authTempDo {
	return a.Clauses(dbresolver.Read)
}

func (a authTempDo) WriteDB() *authTempDo {
	return a.Clauses(dbresolver.Write)
}

func (a authTempDo) Session(config *gorm.Session) *authTempDo {
	return a.withDO(a.DO.Session(config))
}

func (a authTempDo) Clauses(conds ...clause.Expression) *authTempDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authTempDo) Returning(value interface{}, columns ...string) *authTempDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authTempDo) Not(conds ...gen.Condition) *authTempDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authTempDo) Or(conds ...gen.Condition) *authTempDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authTempDo) Select(conds ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authTempDo) Where(conds ...gen.Condition) *authTempDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authTempDo) Order(conds ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authTempDo) Distinct(cols ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authTempDo) Omit(cols ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authTempDo) Join(table schema.Tabler, on ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authTempDo) LeftJoin(table schema.Tabler, on ...field.Expr) *authTempDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authTempDo) RightJoin(table schema.Tabler, on ...field.Expr) *authTempDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authTempDo) Group(cols ...field.Expr) *authTempDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authTempDo) Having(conds ...gen.Condition) *authTempDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authTempDo) Limit(limit int) *authTempDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authTempDo) Offset(offset int) *authTempDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authTempDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *authTempDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authTempDo) Unscoped() *authTempDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authTempDo) Create(values ...*model.AuthTemp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authTempDo) CreateInBatches(values []*model.AuthTemp, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authTempDo) Save(values ...*model.AuthTemp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authTempDo) First() (*model.AuthTemp, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthTemp), nil
	}
}

func (a authTempDo) Take() (*model.AuthTemp, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthTemp), nil
	}
}

func (a authTempDo) Last() (*model.AuthTemp, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthTemp), nil
	}
}

func (a authTempDo) Find() ([]*model.AuthTemp, error) {
	result, err := a.DO.Find()
	return result.([]*model.AuthTemp), err
}

func (a authTempDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthTemp, err error) {
	buf := make([]*model.AuthTemp, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authTempDo) FindInBatches(result *[]*model.AuthTemp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authTempDo) Attrs(attrs ...field.AssignExpr) *authTempDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authTempDo) Assign(attrs ...field.AssignExpr) *authTempDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authTempDo) Joins(fields ...field.RelationField) *authTempDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authTempDo) Preload(fields ...field.RelationField) *authTempDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authTempDo) FirstOrInit() (*model.AuthTemp, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthTemp), nil
	}
}

func (a authTempDo) FirstOrCreate() (*model.AuthTemp, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthTemp), nil
	}
}

func (a authTempDo) FindByPage(offset int, limit int) (result []*model.AuthTemp, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authTempDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authTempDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authTempDo) Delete(models ...*model.AuthTemp) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authTempDo) withDO(do gen.Dao) *authTempDo {
	a.DO = *do.(*gen.DO)
	return a
}
