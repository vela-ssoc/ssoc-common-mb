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

func newKVAudit(db *gorm.DB, opts ...gen.DOOption) kVAudit {
	_kVAudit := kVAudit{}

	_kVAudit.kVAuditDo.UseDB(db, opts...)
	_kVAudit.kVAuditDo.UseModel(&model.KVAudit{})

	tableName := _kVAudit.kVAuditDo.TableName()
	_kVAudit.ALL = field.NewAsterisk(tableName)
	_kVAudit.ID = field.NewInt64(tableName, "id")
	_kVAudit.MinionID = field.NewInt64(tableName, "minion_id")
	_kVAudit.Inet = field.NewString(tableName, "inet")
	_kVAudit.Bucket = field.NewString(tableName, "bucket")
	_kVAudit.Key = field.NewString(tableName, "key")
	_kVAudit.CreatedAt = field.NewTime(tableName, "created_at")
	_kVAudit.UpdatedAt = field.NewTime(tableName, "updated_at")

	_kVAudit.fillFieldMap()

	return _kVAudit
}

type kVAudit struct {
	kVAuditDo kVAuditDo

	ALL       field.Asterisk
	ID        field.Int64
	MinionID  field.Int64
	Inet      field.String
	Bucket    field.String
	Key       field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (k kVAudit) Table(newTableName string) *kVAudit {
	k.kVAuditDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kVAudit) As(alias string) *kVAudit {
	k.kVAuditDo.DO = *(k.kVAuditDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kVAudit) updateTableName(table string) *kVAudit {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewInt64(table, "id")
	k.MinionID = field.NewInt64(table, "minion_id")
	k.Inet = field.NewString(table, "inet")
	k.Bucket = field.NewString(table, "bucket")
	k.Key = field.NewString(table, "key")
	k.CreatedAt = field.NewTime(table, "created_at")
	k.UpdatedAt = field.NewTime(table, "updated_at")

	k.fillFieldMap()

	return k
}

func (k *kVAudit) WithContext(ctx context.Context) *kVAuditDo { return k.kVAuditDo.WithContext(ctx) }

func (k kVAudit) TableName() string { return k.kVAuditDo.TableName() }

func (k kVAudit) Alias() string { return k.kVAuditDo.Alias() }

func (k kVAudit) Columns(cols ...field.Expr) gen.Columns { return k.kVAuditDo.Columns(cols...) }

func (k *kVAudit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kVAudit) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 7)
	k.fieldMap["id"] = k.ID
	k.fieldMap["minion_id"] = k.MinionID
	k.fieldMap["inet"] = k.Inet
	k.fieldMap["bucket"] = k.Bucket
	k.fieldMap["key"] = k.Key
	k.fieldMap["created_at"] = k.CreatedAt
	k.fieldMap["updated_at"] = k.UpdatedAt
}

func (k kVAudit) clone(db *gorm.DB) kVAudit {
	k.kVAuditDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k kVAudit) replaceDB(db *gorm.DB) kVAudit {
	k.kVAuditDo.ReplaceDB(db)
	return k
}

type kVAuditDo struct{ gen.DO }

func (k kVAuditDo) Debug() *kVAuditDo {
	return k.withDO(k.DO.Debug())
}

func (k kVAuditDo) WithContext(ctx context.Context) *kVAuditDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kVAuditDo) ReadDB() *kVAuditDo {
	return k.Clauses(dbresolver.Read)
}

func (k kVAuditDo) WriteDB() *kVAuditDo {
	return k.Clauses(dbresolver.Write)
}

func (k kVAuditDo) Session(config *gorm.Session) *kVAuditDo {
	return k.withDO(k.DO.Session(config))
}

func (k kVAuditDo) Clauses(conds ...clause.Expression) *kVAuditDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kVAuditDo) Returning(value interface{}, columns ...string) *kVAuditDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kVAuditDo) Not(conds ...gen.Condition) *kVAuditDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kVAuditDo) Or(conds ...gen.Condition) *kVAuditDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kVAuditDo) Select(conds ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kVAuditDo) Where(conds ...gen.Condition) *kVAuditDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kVAuditDo) Order(conds ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kVAuditDo) Distinct(cols ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kVAuditDo) Omit(cols ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kVAuditDo) Join(table schema.Tabler, on ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kVAuditDo) LeftJoin(table schema.Tabler, on ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kVAuditDo) RightJoin(table schema.Tabler, on ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kVAuditDo) Group(cols ...field.Expr) *kVAuditDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kVAuditDo) Having(conds ...gen.Condition) *kVAuditDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kVAuditDo) Limit(limit int) *kVAuditDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kVAuditDo) Offset(offset int) *kVAuditDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kVAuditDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *kVAuditDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kVAuditDo) Unscoped() *kVAuditDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kVAuditDo) Create(values ...*model.KVAudit) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kVAuditDo) CreateInBatches(values []*model.KVAudit, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kVAuditDo) Save(values ...*model.KVAudit) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kVAuditDo) First() (*model.KVAudit, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.KVAudit), nil
	}
}

func (k kVAuditDo) Take() (*model.KVAudit, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.KVAudit), nil
	}
}

func (k kVAuditDo) Last() (*model.KVAudit, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.KVAudit), nil
	}
}

func (k kVAuditDo) Find() ([]*model.KVAudit, error) {
	result, err := k.DO.Find()
	return result.([]*model.KVAudit), err
}

func (k kVAuditDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.KVAudit, err error) {
	buf := make([]*model.KVAudit, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kVAuditDo) FindInBatches(result *[]*model.KVAudit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kVAuditDo) Attrs(attrs ...field.AssignExpr) *kVAuditDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kVAuditDo) Assign(attrs ...field.AssignExpr) *kVAuditDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kVAuditDo) Joins(fields ...field.RelationField) *kVAuditDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kVAuditDo) Preload(fields ...field.RelationField) *kVAuditDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kVAuditDo) FirstOrInit() (*model.KVAudit, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.KVAudit), nil
	}
}

func (k kVAuditDo) FirstOrCreate() (*model.KVAudit, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.KVAudit), nil
	}
}

func (k kVAuditDo) FindByPage(offset int, limit int) (result []*model.KVAudit, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kVAuditDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k kVAuditDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k kVAuditDo) Delete(models ...*model.KVAudit) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *kVAuditDo) withDO(do gen.Dao) *kVAuditDo {
	k.DO = *do.(*gen.DO)
	return k
}
