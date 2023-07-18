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

func newJobPolicy(db *gorm.DB, opts ...gen.DOOption) jobPolicy {
	_jobPolicy := jobPolicy{}

	_jobPolicy.jobPolicyDo.UseDB(db, opts...)
	_jobPolicy.jobPolicyDo.UseModel(&model.JobPolicy{})

	tableName := _jobPolicy.jobPolicyDo.TableName()
	_jobPolicy.ALL = field.NewAsterisk(tableName)
	_jobPolicy.ID = field.NewInt64(tableName, "id")
	_jobPolicy.Name = field.NewString(tableName, "name")
	_jobPolicy.Desc = field.NewString(tableName, "desc")
	_jobPolicy.CodeID = field.NewInt64(tableName, "code_id")
	_jobPolicy.Timeout = field.NewInt(tableName, "timeout")
	_jobPolicy.Parallel = field.NewInt(tableName, "parallel")
	_jobPolicy.Args = field.NewBytes(tableName, "args")
	_jobPolicy.CreatedID = field.NewInt64(tableName, "created_id")
	_jobPolicy.CreatedAt = field.NewTime(tableName, "created_at")
	_jobPolicy.UpdatedAt = field.NewTime(tableName, "updated_at")

	_jobPolicy.fillFieldMap()

	return _jobPolicy
}

type jobPolicy struct {
	jobPolicyDo jobPolicyDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	Desc      field.String
	CodeID    field.Int64
	Timeout   field.Int
	Parallel  field.Int
	Args      field.Bytes
	CreatedID field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (j jobPolicy) Table(newTableName string) *jobPolicy {
	j.jobPolicyDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jobPolicy) As(alias string) *jobPolicy {
	j.jobPolicyDo.DO = *(j.jobPolicyDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jobPolicy) updateTableName(table string) *jobPolicy {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewInt64(table, "id")
	j.Name = field.NewString(table, "name")
	j.Desc = field.NewString(table, "desc")
	j.CodeID = field.NewInt64(table, "code_id")
	j.Timeout = field.NewInt(table, "timeout")
	j.Parallel = field.NewInt(table, "parallel")
	j.Args = field.NewBytes(table, "args")
	j.CreatedID = field.NewInt64(table, "created_id")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.UpdatedAt = field.NewTime(table, "updated_at")

	j.fillFieldMap()

	return j
}

func (j *jobPolicy) WithContext(ctx context.Context) *jobPolicyDo {
	return j.jobPolicyDo.WithContext(ctx)
}

func (j jobPolicy) TableName() string { return j.jobPolicyDo.TableName() }

func (j jobPolicy) Alias() string { return j.jobPolicyDo.Alias() }

func (j jobPolicy) Columns(cols ...field.Expr) gen.Columns { return j.jobPolicyDo.Columns(cols...) }

func (j *jobPolicy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jobPolicy) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 10)
	j.fieldMap["id"] = j.ID
	j.fieldMap["name"] = j.Name
	j.fieldMap["desc"] = j.Desc
	j.fieldMap["code_id"] = j.CodeID
	j.fieldMap["timeout"] = j.Timeout
	j.fieldMap["parallel"] = j.Parallel
	j.fieldMap["args"] = j.Args
	j.fieldMap["created_id"] = j.CreatedID
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["updated_at"] = j.UpdatedAt
}

func (j jobPolicy) clone(db *gorm.DB) jobPolicy {
	j.jobPolicyDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jobPolicy) replaceDB(db *gorm.DB) jobPolicy {
	j.jobPolicyDo.ReplaceDB(db)
	return j
}

type jobPolicyDo struct{ gen.DO }

func (j jobPolicyDo) Debug() *jobPolicyDo {
	return j.withDO(j.DO.Debug())
}

func (j jobPolicyDo) WithContext(ctx context.Context) *jobPolicyDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobPolicyDo) ReadDB() *jobPolicyDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobPolicyDo) WriteDB() *jobPolicyDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobPolicyDo) Session(config *gorm.Session) *jobPolicyDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobPolicyDo) Clauses(conds ...clause.Expression) *jobPolicyDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobPolicyDo) Returning(value interface{}, columns ...string) *jobPolicyDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobPolicyDo) Not(conds ...gen.Condition) *jobPolicyDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobPolicyDo) Or(conds ...gen.Condition) *jobPolicyDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobPolicyDo) Select(conds ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobPolicyDo) Where(conds ...gen.Condition) *jobPolicyDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobPolicyDo) Order(conds ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobPolicyDo) Distinct(cols ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobPolicyDo) Omit(cols ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobPolicyDo) Join(table schema.Tabler, on ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobPolicyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobPolicyDo) RightJoin(table schema.Tabler, on ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobPolicyDo) Group(cols ...field.Expr) *jobPolicyDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobPolicyDo) Having(conds ...gen.Condition) *jobPolicyDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobPolicyDo) Limit(limit int) *jobPolicyDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobPolicyDo) Offset(offset int) *jobPolicyDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobPolicyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *jobPolicyDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobPolicyDo) Unscoped() *jobPolicyDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobPolicyDo) Create(values ...*model.JobPolicy) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobPolicyDo) CreateInBatches(values []*model.JobPolicy, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobPolicyDo) Save(values ...*model.JobPolicy) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobPolicyDo) First() (*model.JobPolicy, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobPolicy), nil
	}
}

func (j jobPolicyDo) Take() (*model.JobPolicy, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobPolicy), nil
	}
}

func (j jobPolicyDo) Last() (*model.JobPolicy, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobPolicy), nil
	}
}

func (j jobPolicyDo) Find() ([]*model.JobPolicy, error) {
	result, err := j.DO.Find()
	return result.([]*model.JobPolicy), err
}

func (j jobPolicyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JobPolicy, err error) {
	buf := make([]*model.JobPolicy, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobPolicyDo) FindInBatches(result *[]*model.JobPolicy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobPolicyDo) Attrs(attrs ...field.AssignExpr) *jobPolicyDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobPolicyDo) Assign(attrs ...field.AssignExpr) *jobPolicyDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobPolicyDo) Joins(fields ...field.RelationField) *jobPolicyDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobPolicyDo) Preload(fields ...field.RelationField) *jobPolicyDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobPolicyDo) FirstOrInit() (*model.JobPolicy, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobPolicy), nil
	}
}

func (j jobPolicyDo) FirstOrCreate() (*model.JobPolicy, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobPolicy), nil
	}
}

func (j jobPolicyDo) FindByPage(offset int, limit int) (result []*model.JobPolicy, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobPolicyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobPolicyDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobPolicyDo) Delete(models ...*model.JobPolicy) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobPolicyDo) withDO(do gen.Dao) *jobPolicyDo {
	j.DO = *do.(*gen.DO)
	return j
}
