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

func newStore(db *gorm.DB, opts ...gen.DOOption) store {
	_store := store{}

	_store.storeDo.UseDB(db, opts...)
	_store.storeDo.UseModel(&model.Store{})

	tableName := _store.storeDo.TableName()
	_store.ALL = field.NewAsterisk(tableName)
	_store.ID = field.NewString(tableName, "id")
	_store.Value = field.NewBytes(tableName, "value")
	_store.Escape = field.NewBool(tableName, "escape")
	_store.Desc = field.NewString(tableName, "desc")
	_store.Version = field.NewInt64(tableName, "version")
	_store.CreatedAt = field.NewTime(tableName, "created_at")
	_store.UpdatedAt = field.NewTime(tableName, "updated_at")

	_store.fillFieldMap()

	return _store
}

type store struct {
	storeDo storeDo

	ALL       field.Asterisk
	ID        field.String
	Value     field.Bytes
	Escape    field.Bool
	Desc      field.String
	Version   field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s store) Table(newTableName string) *store {
	s.storeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s store) As(alias string) *store {
	s.storeDo.DO = *(s.storeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *store) updateTableName(table string) *store {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.Value = field.NewBytes(table, "value")
	s.Escape = field.NewBool(table, "escape")
	s.Desc = field.NewString(table, "desc")
	s.Version = field.NewInt64(table, "version")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *store) WithContext(ctx context.Context) *storeDo { return s.storeDo.WithContext(ctx) }

func (s store) TableName() string { return s.storeDo.TableName() }

func (s store) Alias() string { return s.storeDo.Alias() }

func (s store) Columns(cols ...field.Expr) gen.Columns { return s.storeDo.Columns(cols...) }

func (s *store) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *store) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["value"] = s.Value
	s.fieldMap["escape"] = s.Escape
	s.fieldMap["desc"] = s.Desc
	s.fieldMap["version"] = s.Version
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s store) clone(db *gorm.DB) store {
	s.storeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s store) replaceDB(db *gorm.DB) store {
	s.storeDo.ReplaceDB(db)
	return s
}

type storeDo struct{ gen.DO }

func (s storeDo) Debug() *storeDo {
	return s.withDO(s.DO.Debug())
}

func (s storeDo) WithContext(ctx context.Context) *storeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeDo) ReadDB() *storeDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeDo) WriteDB() *storeDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeDo) Session(config *gorm.Session) *storeDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeDo) Clauses(conds ...clause.Expression) *storeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeDo) Returning(value interface{}, columns ...string) *storeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeDo) Not(conds ...gen.Condition) *storeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeDo) Or(conds ...gen.Condition) *storeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeDo) Select(conds ...field.Expr) *storeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeDo) Where(conds ...gen.Condition) *storeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeDo) Order(conds ...field.Expr) *storeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeDo) Distinct(cols ...field.Expr) *storeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeDo) Omit(cols ...field.Expr) *storeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeDo) Join(table schema.Tabler, on ...field.Expr) *storeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *storeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeDo) RightJoin(table schema.Tabler, on ...field.Expr) *storeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeDo) Group(cols ...field.Expr) *storeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeDo) Having(conds ...gen.Condition) *storeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeDo) Limit(limit int) *storeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeDo) Offset(offset int) *storeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *storeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeDo) Unscoped() *storeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeDo) Create(values ...*model.Store) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeDo) CreateInBatches(values []*model.Store, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeDo) Save(values ...*model.Store) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeDo) First() (*model.Store, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Store), nil
	}
}

func (s storeDo) Take() (*model.Store, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Store), nil
	}
}

func (s storeDo) Last() (*model.Store, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Store), nil
	}
}

func (s storeDo) Find() ([]*model.Store, error) {
	result, err := s.DO.Find()
	return result.([]*model.Store), err
}

func (s storeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Store, err error) {
	buf := make([]*model.Store, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeDo) FindInBatches(result *[]*model.Store, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeDo) Attrs(attrs ...field.AssignExpr) *storeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeDo) Assign(attrs ...field.AssignExpr) *storeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeDo) Joins(fields ...field.RelationField) *storeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeDo) Preload(fields ...field.RelationField) *storeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeDo) FirstOrInit() (*model.Store, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Store), nil
	}
}

func (s storeDo) FirstOrCreate() (*model.Store, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Store), nil
	}
}

func (s storeDo) FindByPage(offset int, limit int) (result []*model.Store, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s storeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeDo) Delete(models ...*model.Store) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeDo) withDO(do gen.Dao) *storeDo {
	s.DO = *do.(*gen.DO)
	return s
}
