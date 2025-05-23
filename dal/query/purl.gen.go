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

func newPurl(db *gorm.DB, opts ...gen.DOOption) purl {
	_purl := purl{}

	_purl.purlDo.UseDB(db, opts...)
	_purl.purlDo.UseModel(&model.Purl{})

	tableName := _purl.purlDo.TableName()
	_purl.ALL = field.NewAsterisk(tableName)
	_purl.ID = field.NewString(tableName, "id")

	_purl.fillFieldMap()

	return _purl
}

type purl struct {
	purlDo purlDo

	ALL field.Asterisk
	ID  field.String // PURL

	fieldMap map[string]field.Expr
}

func (p purl) Table(newTableName string) *purl {
	p.purlDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p purl) As(alias string) *purl {
	p.purlDo.DO = *(p.purlDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *purl) updateTableName(table string) *purl {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")

	p.fillFieldMap()

	return p
}

func (p *purl) WithContext(ctx context.Context) *purlDo { return p.purlDo.WithContext(ctx) }

func (p purl) TableName() string { return p.purlDo.TableName() }

func (p purl) Alias() string { return p.purlDo.Alias() }

func (p purl) Columns(cols ...field.Expr) gen.Columns { return p.purlDo.Columns(cols...) }

func (p *purl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *purl) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 1)
	p.fieldMap["id"] = p.ID
}

func (p purl) clone(db *gorm.DB) purl {
	p.purlDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p purl) replaceDB(db *gorm.DB) purl {
	p.purlDo.ReplaceDB(db)
	return p
}

type purlDo struct{ gen.DO }

func (p purlDo) Debug() *purlDo {
	return p.withDO(p.DO.Debug())
}

func (p purlDo) WithContext(ctx context.Context) *purlDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p purlDo) ReadDB() *purlDo {
	return p.Clauses(dbresolver.Read)
}

func (p purlDo) WriteDB() *purlDo {
	return p.Clauses(dbresolver.Write)
}

func (p purlDo) Session(config *gorm.Session) *purlDo {
	return p.withDO(p.DO.Session(config))
}

func (p purlDo) Clauses(conds ...clause.Expression) *purlDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p purlDo) Returning(value interface{}, columns ...string) *purlDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p purlDo) Not(conds ...gen.Condition) *purlDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p purlDo) Or(conds ...gen.Condition) *purlDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p purlDo) Select(conds ...field.Expr) *purlDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p purlDo) Where(conds ...gen.Condition) *purlDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p purlDo) Order(conds ...field.Expr) *purlDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p purlDo) Distinct(cols ...field.Expr) *purlDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p purlDo) Omit(cols ...field.Expr) *purlDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p purlDo) Join(table schema.Tabler, on ...field.Expr) *purlDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p purlDo) LeftJoin(table schema.Tabler, on ...field.Expr) *purlDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p purlDo) RightJoin(table schema.Tabler, on ...field.Expr) *purlDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p purlDo) Group(cols ...field.Expr) *purlDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p purlDo) Having(conds ...gen.Condition) *purlDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p purlDo) Limit(limit int) *purlDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p purlDo) Offset(offset int) *purlDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p purlDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *purlDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p purlDo) Unscoped() *purlDo {
	return p.withDO(p.DO.Unscoped())
}

func (p purlDo) Create(values ...*model.Purl) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p purlDo) CreateInBatches(values []*model.Purl, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p purlDo) Save(values ...*model.Purl) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p purlDo) First() (*model.Purl, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Purl), nil
	}
}

func (p purlDo) Take() (*model.Purl, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Purl), nil
	}
}

func (p purlDo) Last() (*model.Purl, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Purl), nil
	}
}

func (p purlDo) Find() ([]*model.Purl, error) {
	result, err := p.DO.Find()
	return result.([]*model.Purl), err
}

func (p purlDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Purl, err error) {
	buf := make([]*model.Purl, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p purlDo) FindInBatches(result *[]*model.Purl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p purlDo) Attrs(attrs ...field.AssignExpr) *purlDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p purlDo) Assign(attrs ...field.AssignExpr) *purlDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p purlDo) Joins(fields ...field.RelationField) *purlDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p purlDo) Preload(fields ...field.RelationField) *purlDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p purlDo) FirstOrInit() (*model.Purl, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Purl), nil
	}
}

func (p purlDo) FirstOrCreate() (*model.Purl, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Purl), nil
	}
}

func (p purlDo) FindByPage(offset int, limit int) (result []*model.Purl, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p purlDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p purlDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p purlDo) Delete(models ...*model.Purl) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *purlDo) withDO(do gen.Dao) *purlDo {
	p.DO = *do.(*gen.DO)
	return p
}
