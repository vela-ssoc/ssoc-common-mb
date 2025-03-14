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

func newExtensionMarket(db *gorm.DB, opts ...gen.DOOption) extensionMarket {
	_extensionMarket := extensionMarket{}

	_extensionMarket.extensionMarketDo.UseDB(db, opts...)
	_extensionMarket.extensionMarketDo.UseModel(&model.ExtensionMarket{})

	tableName := _extensionMarket.extensionMarketDo.TableName()
	_extensionMarket.ALL = field.NewAsterisk(tableName)
	_extensionMarket.ID = field.NewInt64(tableName, "id")
	_extensionMarket.Name = field.NewString(tableName, "name")
	_extensionMarket.Intro = field.NewString(tableName, "intro")
	_extensionMarket.Category = field.NewString(tableName, "category")
	_extensionMarket.Version = field.NewInt64(tableName, "version")
	_extensionMarket.Content = field.NewString(tableName, "content")
	_extensionMarket.ContentSHA1 = field.NewString(tableName, "content_sha1")
	_extensionMarket.Changelog = field.NewString(tableName, "changelog")
	_extensionMarket.CreatedBy = field.NewField(tableName, "created_by")
	_extensionMarket.UpdatedBy = field.NewField(tableName, "updated_by")
	_extensionMarket.CreatedAt = field.NewTime(tableName, "created_at")
	_extensionMarket.UpdatedAt = field.NewTime(tableName, "updated_at")

	_extensionMarket.fillFieldMap()

	return _extensionMarket
}

type extensionMarket struct {
	extensionMarketDo extensionMarketDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	Intro       field.String
	Category    field.String
	Version     field.Int64
	Content     field.String
	ContentSHA1 field.String
	Changelog   field.String
	CreatedBy   field.Field
	UpdatedBy   field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (e extensionMarket) Table(newTableName string) *extensionMarket {
	e.extensionMarketDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e extensionMarket) As(alias string) *extensionMarket {
	e.extensionMarketDo.DO = *(e.extensionMarketDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *extensionMarket) updateTableName(table string) *extensionMarket {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.Name = field.NewString(table, "name")
	e.Intro = field.NewString(table, "intro")
	e.Category = field.NewString(table, "category")
	e.Version = field.NewInt64(table, "version")
	e.Content = field.NewString(table, "content")
	e.ContentSHA1 = field.NewString(table, "content_sha1")
	e.Changelog = field.NewString(table, "changelog")
	e.CreatedBy = field.NewField(table, "created_by")
	e.UpdatedBy = field.NewField(table, "updated_by")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")

	e.fillFieldMap()

	return e
}

func (e *extensionMarket) WithContext(ctx context.Context) *extensionMarketDo {
	return e.extensionMarketDo.WithContext(ctx)
}

func (e extensionMarket) TableName() string { return e.extensionMarketDo.TableName() }

func (e extensionMarket) Alias() string { return e.extensionMarketDo.Alias() }

func (e extensionMarket) Columns(cols ...field.Expr) gen.Columns {
	return e.extensionMarketDo.Columns(cols...)
}

func (e *extensionMarket) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *extensionMarket) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 12)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["intro"] = e.Intro
	e.fieldMap["category"] = e.Category
	e.fieldMap["version"] = e.Version
	e.fieldMap["content"] = e.Content
	e.fieldMap["content_sha1"] = e.ContentSHA1
	e.fieldMap["changelog"] = e.Changelog
	e.fieldMap["created_by"] = e.CreatedBy
	e.fieldMap["updated_by"] = e.UpdatedBy
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
}

func (e extensionMarket) clone(db *gorm.DB) extensionMarket {
	e.extensionMarketDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e extensionMarket) replaceDB(db *gorm.DB) extensionMarket {
	e.extensionMarketDo.ReplaceDB(db)
	return e
}

type extensionMarketDo struct{ gen.DO }

func (e extensionMarketDo) Debug() *extensionMarketDo {
	return e.withDO(e.DO.Debug())
}

func (e extensionMarketDo) WithContext(ctx context.Context) *extensionMarketDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e extensionMarketDo) ReadDB() *extensionMarketDo {
	return e.Clauses(dbresolver.Read)
}

func (e extensionMarketDo) WriteDB() *extensionMarketDo {
	return e.Clauses(dbresolver.Write)
}

func (e extensionMarketDo) Session(config *gorm.Session) *extensionMarketDo {
	return e.withDO(e.DO.Session(config))
}

func (e extensionMarketDo) Clauses(conds ...clause.Expression) *extensionMarketDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e extensionMarketDo) Returning(value interface{}, columns ...string) *extensionMarketDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e extensionMarketDo) Not(conds ...gen.Condition) *extensionMarketDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e extensionMarketDo) Or(conds ...gen.Condition) *extensionMarketDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e extensionMarketDo) Select(conds ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e extensionMarketDo) Where(conds ...gen.Condition) *extensionMarketDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e extensionMarketDo) Order(conds ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e extensionMarketDo) Distinct(cols ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e extensionMarketDo) Omit(cols ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e extensionMarketDo) Join(table schema.Tabler, on ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e extensionMarketDo) LeftJoin(table schema.Tabler, on ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e extensionMarketDo) RightJoin(table schema.Tabler, on ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e extensionMarketDo) Group(cols ...field.Expr) *extensionMarketDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e extensionMarketDo) Having(conds ...gen.Condition) *extensionMarketDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e extensionMarketDo) Limit(limit int) *extensionMarketDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e extensionMarketDo) Offset(offset int) *extensionMarketDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e extensionMarketDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *extensionMarketDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e extensionMarketDo) Unscoped() *extensionMarketDo {
	return e.withDO(e.DO.Unscoped())
}

func (e extensionMarketDo) Create(values ...*model.ExtensionMarket) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e extensionMarketDo) CreateInBatches(values []*model.ExtensionMarket, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e extensionMarketDo) Save(values ...*model.ExtensionMarket) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e extensionMarketDo) First() (*model.ExtensionMarket, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExtensionMarket), nil
	}
}

func (e extensionMarketDo) Take() (*model.ExtensionMarket, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExtensionMarket), nil
	}
}

func (e extensionMarketDo) Last() (*model.ExtensionMarket, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExtensionMarket), nil
	}
}

func (e extensionMarketDo) Find() ([]*model.ExtensionMarket, error) {
	result, err := e.DO.Find()
	return result.([]*model.ExtensionMarket), err
}

func (e extensionMarketDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ExtensionMarket, err error) {
	buf := make([]*model.ExtensionMarket, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e extensionMarketDo) FindInBatches(result *[]*model.ExtensionMarket, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e extensionMarketDo) Attrs(attrs ...field.AssignExpr) *extensionMarketDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e extensionMarketDo) Assign(attrs ...field.AssignExpr) *extensionMarketDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e extensionMarketDo) Joins(fields ...field.RelationField) *extensionMarketDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e extensionMarketDo) Preload(fields ...field.RelationField) *extensionMarketDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e extensionMarketDo) FirstOrInit() (*model.ExtensionMarket, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExtensionMarket), nil
	}
}

func (e extensionMarketDo) FirstOrCreate() (*model.ExtensionMarket, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExtensionMarket), nil
	}
}

func (e extensionMarketDo) FindByPage(offset int, limit int) (result []*model.ExtensionMarket, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e extensionMarketDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e extensionMarketDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e extensionMarketDo) Delete(models ...*model.ExtensionMarket) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *extensionMarketDo) withDO(do gen.Dao) *extensionMarketDo {
	e.DO = *do.(*gen.DO)
	return e
}
