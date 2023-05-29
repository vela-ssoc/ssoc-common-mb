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

func newMinionTag(db *gorm.DB, opts ...gen.DOOption) minionTag {
	_minionTag := minionTag{}

	_minionTag.minionTagDo.UseDB(db, opts...)
	_minionTag.minionTagDo.UseModel(&model.MinionTag{})

	tableName := _minionTag.minionTagDo.TableName()
	_minionTag.ALL = field.NewAsterisk(tableName)
	_minionTag.ID = field.NewInt64(tableName, "id")
	_minionTag.Tag = field.NewString(tableName, "tag")
	_minionTag.MinionID = field.NewInt64(tableName, "minion_id")
	_minionTag.Kind = field.NewInt8(tableName, "kind")

	_minionTag.fillFieldMap()

	return _minionTag
}

type minionTag struct {
	minionTagDo minionTagDo

	ALL      field.Asterisk
	ID       field.Int64
	Tag      field.String
	MinionID field.Int64
	Kind     field.Int8

	fieldMap map[string]field.Expr
}

func (m minionTag) Table(newTableName string) *minionTag {
	m.minionTagDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m minionTag) As(alias string) *minionTag {
	m.minionTagDo.DO = *(m.minionTagDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *minionTag) updateTableName(table string) *minionTag {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.Tag = field.NewString(table, "tag")
	m.MinionID = field.NewInt64(table, "minion_id")
	m.Kind = field.NewInt8(table, "kind")

	m.fillFieldMap()

	return m
}

func (m *minionTag) WithContext(ctx context.Context) *minionTagDo {
	return m.minionTagDo.WithContext(ctx)
}

func (m minionTag) TableName() string { return m.minionTagDo.TableName() }

func (m minionTag) Alias() string { return m.minionTagDo.Alias() }

func (m *minionTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *minionTag) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["id"] = m.ID
	m.fieldMap["tag"] = m.Tag
	m.fieldMap["minion_id"] = m.MinionID
	m.fieldMap["kind"] = m.Kind
}

func (m minionTag) clone(db *gorm.DB) minionTag {
	m.minionTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m minionTag) replaceDB(db *gorm.DB) minionTag {
	m.minionTagDo.ReplaceDB(db)
	return m
}

type minionTagDo struct{ gen.DO }

func (m minionTagDo) Debug() *minionTagDo {
	return m.withDO(m.DO.Debug())
}

func (m minionTagDo) WithContext(ctx context.Context) *minionTagDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m minionTagDo) ReadDB() *minionTagDo {
	return m.Clauses(dbresolver.Read)
}

func (m minionTagDo) WriteDB() *minionTagDo {
	return m.Clauses(dbresolver.Write)
}

func (m minionTagDo) Session(config *gorm.Session) *minionTagDo {
	return m.withDO(m.DO.Session(config))
}

func (m minionTagDo) Clauses(conds ...clause.Expression) *minionTagDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m minionTagDo) Returning(value interface{}, columns ...string) *minionTagDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m minionTagDo) Not(conds ...gen.Condition) *minionTagDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m minionTagDo) Or(conds ...gen.Condition) *minionTagDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m minionTagDo) Select(conds ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m minionTagDo) Where(conds ...gen.Condition) *minionTagDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m minionTagDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *minionTagDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m minionTagDo) Order(conds ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m minionTagDo) Distinct(cols ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m minionTagDo) Omit(cols ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m minionTagDo) Join(table schema.Tabler, on ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m minionTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m minionTagDo) RightJoin(table schema.Tabler, on ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m minionTagDo) Group(cols ...field.Expr) *minionTagDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m minionTagDo) Having(conds ...gen.Condition) *minionTagDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m minionTagDo) Limit(limit int) *minionTagDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m minionTagDo) Offset(offset int) *minionTagDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m minionTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *minionTagDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m minionTagDo) Unscoped() *minionTagDo {
	return m.withDO(m.DO.Unscoped())
}

func (m minionTagDo) Create(values ...*model.MinionTag) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m minionTagDo) CreateInBatches(values []*model.MinionTag, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m minionTagDo) Save(values ...*model.MinionTag) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m minionTagDo) First() (*model.MinionTag, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionTag), nil
	}
}

func (m minionTagDo) Take() (*model.MinionTag, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionTag), nil
	}
}

func (m minionTagDo) Last() (*model.MinionTag, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionTag), nil
	}
}

func (m minionTagDo) Find() ([]*model.MinionTag, error) {
	result, err := m.DO.Find()
	return result.([]*model.MinionTag), err
}

func (m minionTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinionTag, err error) {
	buf := make([]*model.MinionTag, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m minionTagDo) FindInBatches(result *[]*model.MinionTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m minionTagDo) Attrs(attrs ...field.AssignExpr) *minionTagDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m minionTagDo) Assign(attrs ...field.AssignExpr) *minionTagDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m minionTagDo) Joins(fields ...field.RelationField) *minionTagDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m minionTagDo) Preload(fields ...field.RelationField) *minionTagDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m minionTagDo) FirstOrInit() (*model.MinionTag, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionTag), nil
	}
}

func (m minionTagDo) FirstOrCreate() (*model.MinionTag, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionTag), nil
	}
}

func (m minionTagDo) FindByPage(offset int, limit int) (result []*model.MinionTag, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m minionTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m minionTagDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m minionTagDo) Delete(models ...*model.MinionTag) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *minionTagDo) withDO(do gen.Dao) *minionTagDo {
	m.DO = *do.(*gen.DO)
	return m
}
