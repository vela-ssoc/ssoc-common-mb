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

func newSBOMMinion(db *gorm.DB, opts ...gen.DOOption) sBOMMinion {
	_sBOMMinion := sBOMMinion{}

	_sBOMMinion.sBOMMinionDo.UseDB(db, opts...)
	_sBOMMinion.sBOMMinionDo.UseModel(&model.SBOMMinion{})

	tableName := _sBOMMinion.sBOMMinionDo.TableName()
	_sBOMMinion.ALL = field.NewAsterisk(tableName)
	_sBOMMinion.ID = field.NewInt64(tableName, "id")
	_sBOMMinion.Inet = field.NewString(tableName, "inet")
	_sBOMMinion.CriticalNum = field.NewInt(tableName, "critical_num")
	_sBOMMinion.CriticalScore = field.NewFloat64(tableName, "critical_score")
	_sBOMMinion.HighNum = field.NewInt(tableName, "high_num")
	_sBOMMinion.HighScore = field.NewFloat64(tableName, "high_score")
	_sBOMMinion.MediumNum = field.NewInt(tableName, "medium_num")
	_sBOMMinion.MediumScore = field.NewFloat64(tableName, "medium_score")
	_sBOMMinion.LowNum = field.NewInt(tableName, "low_num")
	_sBOMMinion.LowScore = field.NewFloat64(tableName, "low_score")
	_sBOMMinion.TotalNum = field.NewInt(tableName, "total_num")
	_sBOMMinion.TotalScore = field.NewFloat64(tableName, "total_score")
	_sBOMMinion.Nonce = field.NewInt64(tableName, "nonce")
	_sBOMMinion.UpdatedAt = field.NewTime(tableName, "updated_at")

	_sBOMMinion.fillFieldMap()

	return _sBOMMinion
}

type sBOMMinion struct {
	sBOMMinionDo sBOMMinionDo

	ALL           field.Asterisk
	ID            field.Int64
	Inet          field.String
	CriticalNum   field.Int
	CriticalScore field.Float64
	HighNum       field.Int
	HighScore     field.Float64
	MediumNum     field.Int
	MediumScore   field.Float64
	LowNum        field.Int
	LowScore      field.Float64
	TotalNum      field.Int
	TotalScore    field.Float64
	Nonce         field.Int64
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (s sBOMMinion) Table(newTableName string) *sBOMMinion {
	s.sBOMMinionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sBOMMinion) As(alias string) *sBOMMinion {
	s.sBOMMinionDo.DO = *(s.sBOMMinionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sBOMMinion) updateTableName(table string) *sBOMMinion {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Inet = field.NewString(table, "inet")
	s.CriticalNum = field.NewInt(table, "critical_num")
	s.CriticalScore = field.NewFloat64(table, "critical_score")
	s.HighNum = field.NewInt(table, "high_num")
	s.HighScore = field.NewFloat64(table, "high_score")
	s.MediumNum = field.NewInt(table, "medium_num")
	s.MediumScore = field.NewFloat64(table, "medium_score")
	s.LowNum = field.NewInt(table, "low_num")
	s.LowScore = field.NewFloat64(table, "low_score")
	s.TotalNum = field.NewInt(table, "total_num")
	s.TotalScore = field.NewFloat64(table, "total_score")
	s.Nonce = field.NewInt64(table, "nonce")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *sBOMMinion) WithContext(ctx context.Context) *sBOMMinionDo {
	return s.sBOMMinionDo.WithContext(ctx)
}

func (s sBOMMinion) TableName() string { return s.sBOMMinionDo.TableName() }

func (s sBOMMinion) Alias() string { return s.sBOMMinionDo.Alias() }

func (s sBOMMinion) Columns(cols ...field.Expr) gen.Columns { return s.sBOMMinionDo.Columns(cols...) }

func (s *sBOMMinion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sBOMMinion) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["inet"] = s.Inet
	s.fieldMap["critical_num"] = s.CriticalNum
	s.fieldMap["critical_score"] = s.CriticalScore
	s.fieldMap["high_num"] = s.HighNum
	s.fieldMap["high_score"] = s.HighScore
	s.fieldMap["medium_num"] = s.MediumNum
	s.fieldMap["medium_score"] = s.MediumScore
	s.fieldMap["low_num"] = s.LowNum
	s.fieldMap["low_score"] = s.LowScore
	s.fieldMap["total_num"] = s.TotalNum
	s.fieldMap["total_score"] = s.TotalScore
	s.fieldMap["nonce"] = s.Nonce
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s sBOMMinion) clone(db *gorm.DB) sBOMMinion {
	s.sBOMMinionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sBOMMinion) replaceDB(db *gorm.DB) sBOMMinion {
	s.sBOMMinionDo.ReplaceDB(db)
	return s
}

type sBOMMinionDo struct{ gen.DO }

func (s sBOMMinionDo) Debug() *sBOMMinionDo {
	return s.withDO(s.DO.Debug())
}

func (s sBOMMinionDo) WithContext(ctx context.Context) *sBOMMinionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sBOMMinionDo) ReadDB() *sBOMMinionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sBOMMinionDo) WriteDB() *sBOMMinionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sBOMMinionDo) Session(config *gorm.Session) *sBOMMinionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sBOMMinionDo) Clauses(conds ...clause.Expression) *sBOMMinionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sBOMMinionDo) Returning(value interface{}, columns ...string) *sBOMMinionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sBOMMinionDo) Not(conds ...gen.Condition) *sBOMMinionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sBOMMinionDo) Or(conds ...gen.Condition) *sBOMMinionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sBOMMinionDo) Select(conds ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sBOMMinionDo) Where(conds ...gen.Condition) *sBOMMinionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sBOMMinionDo) Order(conds ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sBOMMinionDo) Distinct(cols ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sBOMMinionDo) Omit(cols ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sBOMMinionDo) Join(table schema.Tabler, on ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sBOMMinionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sBOMMinionDo) RightJoin(table schema.Tabler, on ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sBOMMinionDo) Group(cols ...field.Expr) *sBOMMinionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sBOMMinionDo) Having(conds ...gen.Condition) *sBOMMinionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sBOMMinionDo) Limit(limit int) *sBOMMinionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sBOMMinionDo) Offset(offset int) *sBOMMinionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sBOMMinionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sBOMMinionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sBOMMinionDo) Unscoped() *sBOMMinionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sBOMMinionDo) Create(values ...*model.SBOMMinion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sBOMMinionDo) CreateInBatches(values []*model.SBOMMinion, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sBOMMinionDo) Save(values ...*model.SBOMMinion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sBOMMinionDo) First() (*model.SBOMMinion, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMMinion), nil
	}
}

func (s sBOMMinionDo) Take() (*model.SBOMMinion, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMMinion), nil
	}
}

func (s sBOMMinionDo) Last() (*model.SBOMMinion, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMMinion), nil
	}
}

func (s sBOMMinionDo) Find() ([]*model.SBOMMinion, error) {
	result, err := s.DO.Find()
	return result.([]*model.SBOMMinion), err
}

func (s sBOMMinionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SBOMMinion, err error) {
	buf := make([]*model.SBOMMinion, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sBOMMinionDo) FindInBatches(result *[]*model.SBOMMinion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sBOMMinionDo) Attrs(attrs ...field.AssignExpr) *sBOMMinionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sBOMMinionDo) Assign(attrs ...field.AssignExpr) *sBOMMinionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sBOMMinionDo) Joins(fields ...field.RelationField) *sBOMMinionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sBOMMinionDo) Preload(fields ...field.RelationField) *sBOMMinionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sBOMMinionDo) FirstOrInit() (*model.SBOMMinion, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMMinion), nil
	}
}

func (s sBOMMinionDo) FirstOrCreate() (*model.SBOMMinion, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMMinion), nil
	}
}

func (s sBOMMinionDo) FindByPage(offset int, limit int) (result []*model.SBOMMinion, count int64, err error) {
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

func (s sBOMMinionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sBOMMinionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sBOMMinionDo) Delete(models ...*model.SBOMMinion) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sBOMMinionDo) withDO(do gen.Dao) *sBOMMinionDo {
	s.DO = *do.(*gen.DO)
	return s
}
