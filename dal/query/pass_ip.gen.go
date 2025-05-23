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

func newPassIP(db *gorm.DB, opts ...gen.DOOption) passIP {
	_passIP := passIP{}

	_passIP.passIPDo.UseDB(db, opts...)
	_passIP.passIPDo.UseModel(&model.PassIP{})

	tableName := _passIP.passIPDo.TableName()
	_passIP.ALL = field.NewAsterisk(tableName)
	_passIP.ID = field.NewInt64(tableName, "id")
	_passIP.IP = field.NewString(tableName, "ip")
	_passIP.Kind = field.NewString(tableName, "kind")
	_passIP.BeforeAt = field.NewTime(tableName, "before_at")
	_passIP.CreatedAt = field.NewTime(tableName, "created_at")
	_passIP.UpdatedAt = field.NewTime(tableName, "updated_at")

	_passIP.fillFieldMap()

	return _passIP
}

type passIP struct {
	passIPDo passIPDo

	ALL       field.Asterisk
	ID        field.Int64 // ID
	IP        field.String
	Kind      field.String
	BeforeAt  field.Time
	CreatedAt field.Time // 创建时间
	UpdatedAt field.Time // 更新时间

	fieldMap map[string]field.Expr
}

func (p passIP) Table(newTableName string) *passIP {
	p.passIPDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p passIP) As(alias string) *passIP {
	p.passIPDo.DO = *(p.passIPDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *passIP) updateTableName(table string) *passIP {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.IP = field.NewString(table, "ip")
	p.Kind = field.NewString(table, "kind")
	p.BeforeAt = field.NewTime(table, "before_at")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *passIP) WithContext(ctx context.Context) *passIPDo { return p.passIPDo.WithContext(ctx) }

func (p passIP) TableName() string { return p.passIPDo.TableName() }

func (p passIP) Alias() string { return p.passIPDo.Alias() }

func (p passIP) Columns(cols ...field.Expr) gen.Columns { return p.passIPDo.Columns(cols...) }

func (p *passIP) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *passIP) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["ip"] = p.IP
	p.fieldMap["kind"] = p.Kind
	p.fieldMap["before_at"] = p.BeforeAt
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p passIP) clone(db *gorm.DB) passIP {
	p.passIPDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p passIP) replaceDB(db *gorm.DB) passIP {
	p.passIPDo.ReplaceDB(db)
	return p
}

type passIPDo struct{ gen.DO }

func (p passIPDo) Debug() *passIPDo {
	return p.withDO(p.DO.Debug())
}

func (p passIPDo) WithContext(ctx context.Context) *passIPDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passIPDo) ReadDB() *passIPDo {
	return p.Clauses(dbresolver.Read)
}

func (p passIPDo) WriteDB() *passIPDo {
	return p.Clauses(dbresolver.Write)
}

func (p passIPDo) Session(config *gorm.Session) *passIPDo {
	return p.withDO(p.DO.Session(config))
}

func (p passIPDo) Clauses(conds ...clause.Expression) *passIPDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passIPDo) Returning(value interface{}, columns ...string) *passIPDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p passIPDo) Not(conds ...gen.Condition) *passIPDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passIPDo) Or(conds ...gen.Condition) *passIPDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passIPDo) Select(conds ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passIPDo) Where(conds ...gen.Condition) *passIPDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passIPDo) Order(conds ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passIPDo) Distinct(cols ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passIPDo) Omit(cols ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passIPDo) Join(table schema.Tabler, on ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passIPDo) LeftJoin(table schema.Tabler, on ...field.Expr) *passIPDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passIPDo) RightJoin(table schema.Tabler, on ...field.Expr) *passIPDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passIPDo) Group(cols ...field.Expr) *passIPDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passIPDo) Having(conds ...gen.Condition) *passIPDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passIPDo) Limit(limit int) *passIPDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passIPDo) Offset(offset int) *passIPDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passIPDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *passIPDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passIPDo) Unscoped() *passIPDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passIPDo) Create(values ...*model.PassIP) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passIPDo) CreateInBatches(values []*model.PassIP, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passIPDo) Save(values ...*model.PassIP) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passIPDo) First() (*model.PassIP, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassIP), nil
	}
}

func (p passIPDo) Take() (*model.PassIP, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassIP), nil
	}
}

func (p passIPDo) Last() (*model.PassIP, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassIP), nil
	}
}

func (p passIPDo) Find() ([]*model.PassIP, error) {
	result, err := p.DO.Find()
	return result.([]*model.PassIP), err
}

func (p passIPDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PassIP, err error) {
	buf := make([]*model.PassIP, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p passIPDo) FindInBatches(result *[]*model.PassIP, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p passIPDo) Attrs(attrs ...field.AssignExpr) *passIPDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passIPDo) Assign(attrs ...field.AssignExpr) *passIPDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passIPDo) Joins(fields ...field.RelationField) *passIPDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p passIPDo) Preload(fields ...field.RelationField) *passIPDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p passIPDo) FirstOrInit() (*model.PassIP, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassIP), nil
	}
}

func (p passIPDo) FirstOrCreate() (*model.PassIP, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassIP), nil
	}
}

func (p passIPDo) FindByPage(offset int, limit int) (result []*model.PassIP, count int64, err error) {
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

func (p passIPDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p passIPDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p passIPDo) Delete(models ...*model.PassIP) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *passIPDo) withDO(do gen.Dao) *passIPDo {
	p.DO = *do.(*gen.DO)
	return p
}
