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

func newSBOMComponent(db *gorm.DB, opts ...gen.DOOption) sBOMComponent {
	_sBOMComponent := sBOMComponent{}

	_sBOMComponent.sBOMComponentDo.UseDB(db, opts...)
	_sBOMComponent.sBOMComponentDo.UseModel(&model.SBOMComponent{})

	tableName := _sBOMComponent.sBOMComponentDo.TableName()
	_sBOMComponent.ALL = field.NewAsterisk(tableName)
	_sBOMComponent.ID = field.NewInt64(tableName, "id")
	_sBOMComponent.MinionID = field.NewInt64(tableName, "minion_id")
	_sBOMComponent.Inet = field.NewString(tableName, "inet")
	_sBOMComponent.ProjectID = field.NewInt64(tableName, "project_id")
	_sBOMComponent.Filepath = field.NewString(tableName, "filepath")
	_sBOMComponent.SHA1 = field.NewString(tableName, "sha1")
	_sBOMComponent.Name = field.NewString(tableName, "name")
	_sBOMComponent.Version = field.NewString(tableName, "version")
	_sBOMComponent.Language = field.NewString(tableName, "language")
	_sBOMComponent.Licenses = field.NewField(tableName, "licenses")
	_sBOMComponent.PURL = field.NewString(tableName, "purl")
	_sBOMComponent.CriticalNum = field.NewInt(tableName, "critical_num")
	_sBOMComponent.CriticalScore = field.NewFloat64(tableName, "critical_score")
	_sBOMComponent.HighNum = field.NewInt(tableName, "high_num")
	_sBOMComponent.HighScore = field.NewFloat64(tableName, "high_score")
	_sBOMComponent.MediumNum = field.NewInt(tableName, "medium_num")
	_sBOMComponent.MediumScore = field.NewFloat64(tableName, "medium_score")
	_sBOMComponent.LowNum = field.NewInt(tableName, "low_num")
	_sBOMComponent.LowScore = field.NewFloat64(tableName, "low_score")
	_sBOMComponent.TotalNum = field.NewInt(tableName, "total_num")
	_sBOMComponent.TotalScore = field.NewFloat64(tableName, "total_score")
	_sBOMComponent.Status = field.NewUint8(tableName, "status")
	_sBOMComponent.Nonce = field.NewInt64(tableName, "nonce")
	_sBOMComponent.CreatedAt = field.NewTime(tableName, "created_at")
	_sBOMComponent.UpdatedAt = field.NewTime(tableName, "updated_at")

	_sBOMComponent.fillFieldMap()

	return _sBOMComponent
}

type sBOMComponent struct {
	sBOMComponentDo sBOMComponentDo

	ALL           field.Asterisk
	ID            field.Int64
	MinionID      field.Int64
	Inet          field.String
	ProjectID     field.Int64
	Filepath      field.String
	SHA1          field.String
	Name          field.String
	Version       field.String
	Language      field.String
	Licenses      field.Field
	PURL          field.String
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
	Status        field.Uint8
	Nonce         field.Int64
	CreatedAt     field.Time // 创建时间
	UpdatedAt     field.Time // 更新时间

	fieldMap map[string]field.Expr
}

func (s sBOMComponent) Table(newTableName string) *sBOMComponent {
	s.sBOMComponentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sBOMComponent) As(alias string) *sBOMComponent {
	s.sBOMComponentDo.DO = *(s.sBOMComponentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sBOMComponent) updateTableName(table string) *sBOMComponent {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.MinionID = field.NewInt64(table, "minion_id")
	s.Inet = field.NewString(table, "inet")
	s.ProjectID = field.NewInt64(table, "project_id")
	s.Filepath = field.NewString(table, "filepath")
	s.SHA1 = field.NewString(table, "sha1")
	s.Name = field.NewString(table, "name")
	s.Version = field.NewString(table, "version")
	s.Language = field.NewString(table, "language")
	s.Licenses = field.NewField(table, "licenses")
	s.PURL = field.NewString(table, "purl")
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
	s.Status = field.NewUint8(table, "status")
	s.Nonce = field.NewInt64(table, "nonce")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *sBOMComponent) WithContext(ctx context.Context) *sBOMComponentDo {
	return s.sBOMComponentDo.WithContext(ctx)
}

func (s sBOMComponent) TableName() string { return s.sBOMComponentDo.TableName() }

func (s sBOMComponent) Alias() string { return s.sBOMComponentDo.Alias() }

func (s sBOMComponent) Columns(cols ...field.Expr) gen.Columns {
	return s.sBOMComponentDo.Columns(cols...)
}

func (s *sBOMComponent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sBOMComponent) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 25)
	s.fieldMap["id"] = s.ID
	s.fieldMap["minion_id"] = s.MinionID
	s.fieldMap["inet"] = s.Inet
	s.fieldMap["project_id"] = s.ProjectID
	s.fieldMap["filepath"] = s.Filepath
	s.fieldMap["sha1"] = s.SHA1
	s.fieldMap["name"] = s.Name
	s.fieldMap["version"] = s.Version
	s.fieldMap["language"] = s.Language
	s.fieldMap["licenses"] = s.Licenses
	s.fieldMap["purl"] = s.PURL
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
	s.fieldMap["status"] = s.Status
	s.fieldMap["nonce"] = s.Nonce
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s sBOMComponent) clone(db *gorm.DB) sBOMComponent {
	s.sBOMComponentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sBOMComponent) replaceDB(db *gorm.DB) sBOMComponent {
	s.sBOMComponentDo.ReplaceDB(db)
	return s
}

type sBOMComponentDo struct{ gen.DO }

func (s sBOMComponentDo) Debug() *sBOMComponentDo {
	return s.withDO(s.DO.Debug())
}

func (s sBOMComponentDo) WithContext(ctx context.Context) *sBOMComponentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sBOMComponentDo) ReadDB() *sBOMComponentDo {
	return s.Clauses(dbresolver.Read)
}

func (s sBOMComponentDo) WriteDB() *sBOMComponentDo {
	return s.Clauses(dbresolver.Write)
}

func (s sBOMComponentDo) Session(config *gorm.Session) *sBOMComponentDo {
	return s.withDO(s.DO.Session(config))
}

func (s sBOMComponentDo) Clauses(conds ...clause.Expression) *sBOMComponentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sBOMComponentDo) Returning(value interface{}, columns ...string) *sBOMComponentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sBOMComponentDo) Not(conds ...gen.Condition) *sBOMComponentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sBOMComponentDo) Or(conds ...gen.Condition) *sBOMComponentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sBOMComponentDo) Select(conds ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sBOMComponentDo) Where(conds ...gen.Condition) *sBOMComponentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sBOMComponentDo) Order(conds ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sBOMComponentDo) Distinct(cols ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sBOMComponentDo) Omit(cols ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sBOMComponentDo) Join(table schema.Tabler, on ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sBOMComponentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sBOMComponentDo) RightJoin(table schema.Tabler, on ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sBOMComponentDo) Group(cols ...field.Expr) *sBOMComponentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sBOMComponentDo) Having(conds ...gen.Condition) *sBOMComponentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sBOMComponentDo) Limit(limit int) *sBOMComponentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sBOMComponentDo) Offset(offset int) *sBOMComponentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sBOMComponentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sBOMComponentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sBOMComponentDo) Unscoped() *sBOMComponentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sBOMComponentDo) Create(values ...*model.SBOMComponent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sBOMComponentDo) CreateInBatches(values []*model.SBOMComponent, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sBOMComponentDo) Save(values ...*model.SBOMComponent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sBOMComponentDo) First() (*model.SBOMComponent, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMComponent), nil
	}
}

func (s sBOMComponentDo) Take() (*model.SBOMComponent, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMComponent), nil
	}
}

func (s sBOMComponentDo) Last() (*model.SBOMComponent, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMComponent), nil
	}
}

func (s sBOMComponentDo) Find() ([]*model.SBOMComponent, error) {
	result, err := s.DO.Find()
	return result.([]*model.SBOMComponent), err
}

func (s sBOMComponentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SBOMComponent, err error) {
	buf := make([]*model.SBOMComponent, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sBOMComponentDo) FindInBatches(result *[]*model.SBOMComponent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sBOMComponentDo) Attrs(attrs ...field.AssignExpr) *sBOMComponentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sBOMComponentDo) Assign(attrs ...field.AssignExpr) *sBOMComponentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sBOMComponentDo) Joins(fields ...field.RelationField) *sBOMComponentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sBOMComponentDo) Preload(fields ...field.RelationField) *sBOMComponentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sBOMComponentDo) FirstOrInit() (*model.SBOMComponent, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMComponent), nil
	}
}

func (s sBOMComponentDo) FirstOrCreate() (*model.SBOMComponent, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SBOMComponent), nil
	}
}

func (s sBOMComponentDo) FindByPage(offset int, limit int) (result []*model.SBOMComponent, count int64, err error) {
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

func (s sBOMComponentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sBOMComponentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sBOMComponentDo) Delete(models ...*model.SBOMComponent) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sBOMComponentDo) withDO(do gen.Dao) *sBOMComponentDo {
	s.DO = *do.(*gen.DO)
	return s
}
