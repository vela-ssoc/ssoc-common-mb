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

func newOplog(db *gorm.DB, opts ...gen.DOOption) oplog {
	_oplog := oplog{}

	_oplog.oplogDo.UseDB(db, opts...)
	_oplog.oplogDo.UseModel(&model.Oplog{})

	tableName := _oplog.oplogDo.TableName()
	_oplog.ALL = field.NewAsterisk(tableName)
	_oplog.ID = field.NewInt64(tableName, "id")
	_oplog.UserID = field.NewInt64(tableName, "user_id")
	_oplog.Username = field.NewString(tableName, "username")
	_oplog.Nickname = field.NewString(tableName, "nickname")
	_oplog.Name = field.NewString(tableName, "name")
	_oplog.ClientAddr = field.NewString(tableName, "client_addr")
	_oplog.DirectAddr = field.NewString(tableName, "direct_addr")
	_oplog.Method = field.NewString(tableName, "method")
	_oplog.Path = field.NewString(tableName, "path")
	_oplog.Query = field.NewString(tableName, "query")
	_oplog.Length = field.NewInt64(tableName, "length")
	_oplog.Content = field.NewBytes(tableName, "content")
	_oplog.Failed = field.NewBool(tableName, "failed")
	_oplog.Cause = field.NewString(tableName, "cause")
	_oplog.RequestAt = field.NewTime(tableName, "request_at")
	_oplog.Elapsed = field.NewInt64(tableName, "elapsed")
	_oplog.CreatedAt = field.NewTime(tableName, "created_at")

	_oplog.fillFieldMap()

	return _oplog
}

type oplog struct {
	oplogDo oplogDo

	ALL        field.Asterisk
	ID         field.Int64
	UserID     field.Int64
	Username   field.String
	Nickname   field.String
	Name       field.String
	ClientAddr field.String
	DirectAddr field.String
	Method     field.String
	Path       field.String
	Query      field.String
	Length     field.Int64
	Content    field.Bytes
	Failed     field.Bool
	Cause      field.String
	RequestAt  field.Time
	Elapsed    field.Int64
	CreatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (o oplog) Table(newTableName string) *oplog {
	o.oplogDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oplog) As(alias string) *oplog {
	o.oplogDo.DO = *(o.oplogDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oplog) updateTableName(table string) *oplog {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.UserID = field.NewInt64(table, "user_id")
	o.Username = field.NewString(table, "username")
	o.Nickname = field.NewString(table, "nickname")
	o.Name = field.NewString(table, "name")
	o.ClientAddr = field.NewString(table, "client_addr")
	o.DirectAddr = field.NewString(table, "direct_addr")
	o.Method = field.NewString(table, "method")
	o.Path = field.NewString(table, "path")
	o.Query = field.NewString(table, "query")
	o.Length = field.NewInt64(table, "length")
	o.Content = field.NewBytes(table, "content")
	o.Failed = field.NewBool(table, "failed")
	o.Cause = field.NewString(table, "cause")
	o.RequestAt = field.NewTime(table, "request_at")
	o.Elapsed = field.NewInt64(table, "elapsed")
	o.CreatedAt = field.NewTime(table, "created_at")

	o.fillFieldMap()

	return o
}

func (o *oplog) WithContext(ctx context.Context) *oplogDo { return o.oplogDo.WithContext(ctx) }

func (o oplog) TableName() string { return o.oplogDo.TableName() }

func (o oplog) Alias() string { return o.oplogDo.Alias() }

func (o *oplog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oplog) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 17)
	o.fieldMap["id"] = o.ID
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["username"] = o.Username
	o.fieldMap["nickname"] = o.Nickname
	o.fieldMap["name"] = o.Name
	o.fieldMap["client_addr"] = o.ClientAddr
	o.fieldMap["direct_addr"] = o.DirectAddr
	o.fieldMap["method"] = o.Method
	o.fieldMap["path"] = o.Path
	o.fieldMap["query"] = o.Query
	o.fieldMap["length"] = o.Length
	o.fieldMap["content"] = o.Content
	o.fieldMap["failed"] = o.Failed
	o.fieldMap["cause"] = o.Cause
	o.fieldMap["request_at"] = o.RequestAt
	o.fieldMap["elapsed"] = o.Elapsed
	o.fieldMap["created_at"] = o.CreatedAt
}

func (o oplog) clone(db *gorm.DB) oplog {
	o.oplogDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o oplog) replaceDB(db *gorm.DB) oplog {
	o.oplogDo.ReplaceDB(db)
	return o
}

type oplogDo struct{ gen.DO }

func (o oplogDo) Debug() *oplogDo {
	return o.withDO(o.DO.Debug())
}

func (o oplogDo) WithContext(ctx context.Context) *oplogDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oplogDo) ReadDB() *oplogDo {
	return o.Clauses(dbresolver.Read)
}

func (o oplogDo) WriteDB() *oplogDo {
	return o.Clauses(dbresolver.Write)
}

func (o oplogDo) Session(config *gorm.Session) *oplogDo {
	return o.withDO(o.DO.Session(config))
}

func (o oplogDo) Clauses(conds ...clause.Expression) *oplogDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oplogDo) Returning(value interface{}, columns ...string) *oplogDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o oplogDo) Not(conds ...gen.Condition) *oplogDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oplogDo) Or(conds ...gen.Condition) *oplogDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oplogDo) Select(conds ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oplogDo) Where(conds ...gen.Condition) *oplogDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oplogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oplogDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o oplogDo) Order(conds ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oplogDo) Distinct(cols ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oplogDo) Omit(cols ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oplogDo) Join(table schema.Tabler, on ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oplogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oplogDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oplogDo) RightJoin(table schema.Tabler, on ...field.Expr) *oplogDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oplogDo) Group(cols ...field.Expr) *oplogDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oplogDo) Having(conds ...gen.Condition) *oplogDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oplogDo) Limit(limit int) *oplogDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oplogDo) Offset(offset int) *oplogDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oplogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oplogDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oplogDo) Unscoped() *oplogDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oplogDo) Create(values ...*model.Oplog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oplogDo) CreateInBatches(values []*model.Oplog, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oplogDo) Save(values ...*model.Oplog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oplogDo) First() (*model.Oplog, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oplog), nil
	}
}

func (o oplogDo) Take() (*model.Oplog, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oplog), nil
	}
}

func (o oplogDo) Last() (*model.Oplog, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oplog), nil
	}
}

func (o oplogDo) Find() ([]*model.Oplog, error) {
	result, err := o.DO.Find()
	return result.([]*model.Oplog), err
}

func (o oplogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Oplog, err error) {
	buf := make([]*model.Oplog, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oplogDo) FindInBatches(result *[]*model.Oplog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oplogDo) Attrs(attrs ...field.AssignExpr) *oplogDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oplogDo) Assign(attrs ...field.AssignExpr) *oplogDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oplogDo) Joins(fields ...field.RelationField) *oplogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o oplogDo) Preload(fields ...field.RelationField) *oplogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o oplogDo) FirstOrInit() (*model.Oplog, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oplog), nil
	}
}

func (o oplogDo) FirstOrCreate() (*model.Oplog, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Oplog), nil
	}
}

func (o oplogDo) FindByPage(offset int, limit int) (result []*model.Oplog, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o oplogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o oplogDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o oplogDo) Delete(models ...*model.Oplog) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *oplogDo) withDO(do gen.Dao) *oplogDo {
	o.DO = *do.(*gen.DO)
	return o
}
