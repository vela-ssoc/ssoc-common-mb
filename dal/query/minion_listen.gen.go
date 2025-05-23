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

func newMinionListen(db *gorm.DB, opts ...gen.DOOption) minionListen {
	_minionListen := minionListen{}

	_minionListen.minionListenDo.UseDB(db, opts...)
	_minionListen.minionListenDo.UseModel(&model.MinionListen{})

	tableName := _minionListen.minionListenDo.TableName()
	_minionListen.ALL = field.NewAsterisk(tableName)
	_minionListen.ID = field.NewInt64(tableName, "id")
	_minionListen.MinionID = field.NewInt64(tableName, "minion_id")
	_minionListen.Inet = field.NewString(tableName, "inet")
	_minionListen.RecordID = field.NewString(tableName, "record_id")
	_minionListen.PID = field.NewUint32(tableName, "pid")
	_minionListen.FD = field.NewInt(tableName, "fd")
	_minionListen.Family = field.NewUint8(tableName, "family")
	_minionListen.Protocol = field.NewUint8(tableName, "protocol")
	_minionListen.LocalIP = field.NewString(tableName, "local_ip")
	_minionListen.LocalPort = field.NewInt(tableName, "local_port")
	_minionListen.Path = field.NewString(tableName, "path")
	_minionListen.Process = field.NewString(tableName, "process")
	_minionListen.Username = field.NewString(tableName, "username")
	_minionListen.CreatedAt = field.NewTime(tableName, "created_at")
	_minionListen.UpdatedAt = field.NewTime(tableName, "updated_at")

	_minionListen.fillFieldMap()

	return _minionListen
}

type minionListen struct {
	minionListenDo minionListenDo

	ALL       field.Asterisk
	ID        field.Int64 // ID
	MinionID  field.Int64
	Inet      field.String
	RecordID  field.String
	PID       field.Uint32
	FD        field.Int
	Family    field.Uint8
	Protocol  field.Uint8
	LocalIP   field.String
	LocalPort field.Int
	Path      field.String
	Process   field.String
	Username  field.String
	CreatedAt field.Time // 创建时间
	UpdatedAt field.Time // 更新时间

	fieldMap map[string]field.Expr
}

func (m minionListen) Table(newTableName string) *minionListen {
	m.minionListenDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m minionListen) As(alias string) *minionListen {
	m.minionListenDo.DO = *(m.minionListenDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *minionListen) updateTableName(table string) *minionListen {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MinionID = field.NewInt64(table, "minion_id")
	m.Inet = field.NewString(table, "inet")
	m.RecordID = field.NewString(table, "record_id")
	m.PID = field.NewUint32(table, "pid")
	m.FD = field.NewInt(table, "fd")
	m.Family = field.NewUint8(table, "family")
	m.Protocol = field.NewUint8(table, "protocol")
	m.LocalIP = field.NewString(table, "local_ip")
	m.LocalPort = field.NewInt(table, "local_port")
	m.Path = field.NewString(table, "path")
	m.Process = field.NewString(table, "process")
	m.Username = field.NewString(table, "username")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *minionListen) WithContext(ctx context.Context) *minionListenDo {
	return m.minionListenDo.WithContext(ctx)
}

func (m minionListen) TableName() string { return m.minionListenDo.TableName() }

func (m minionListen) Alias() string { return m.minionListenDo.Alias() }

func (m minionListen) Columns(cols ...field.Expr) gen.Columns {
	return m.minionListenDo.Columns(cols...)
}

func (m *minionListen) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *minionListen) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 15)
	m.fieldMap["id"] = m.ID
	m.fieldMap["minion_id"] = m.MinionID
	m.fieldMap["inet"] = m.Inet
	m.fieldMap["record_id"] = m.RecordID
	m.fieldMap["pid"] = m.PID
	m.fieldMap["fd"] = m.FD
	m.fieldMap["family"] = m.Family
	m.fieldMap["protocol"] = m.Protocol
	m.fieldMap["local_ip"] = m.LocalIP
	m.fieldMap["local_port"] = m.LocalPort
	m.fieldMap["path"] = m.Path
	m.fieldMap["process"] = m.Process
	m.fieldMap["username"] = m.Username
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m minionListen) clone(db *gorm.DB) minionListen {
	m.minionListenDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m minionListen) replaceDB(db *gorm.DB) minionListen {
	m.minionListenDo.ReplaceDB(db)
	return m
}

type minionListenDo struct{ gen.DO }

func (m minionListenDo) Debug() *minionListenDo {
	return m.withDO(m.DO.Debug())
}

func (m minionListenDo) WithContext(ctx context.Context) *minionListenDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m minionListenDo) ReadDB() *minionListenDo {
	return m.Clauses(dbresolver.Read)
}

func (m minionListenDo) WriteDB() *minionListenDo {
	return m.Clauses(dbresolver.Write)
}

func (m minionListenDo) Session(config *gorm.Session) *minionListenDo {
	return m.withDO(m.DO.Session(config))
}

func (m minionListenDo) Clauses(conds ...clause.Expression) *minionListenDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m minionListenDo) Returning(value interface{}, columns ...string) *minionListenDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m minionListenDo) Not(conds ...gen.Condition) *minionListenDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m minionListenDo) Or(conds ...gen.Condition) *minionListenDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m minionListenDo) Select(conds ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m minionListenDo) Where(conds ...gen.Condition) *minionListenDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m minionListenDo) Order(conds ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m minionListenDo) Distinct(cols ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m minionListenDo) Omit(cols ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m minionListenDo) Join(table schema.Tabler, on ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m minionListenDo) LeftJoin(table schema.Tabler, on ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m minionListenDo) RightJoin(table schema.Tabler, on ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m minionListenDo) Group(cols ...field.Expr) *minionListenDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m minionListenDo) Having(conds ...gen.Condition) *minionListenDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m minionListenDo) Limit(limit int) *minionListenDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m minionListenDo) Offset(offset int) *minionListenDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m minionListenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *minionListenDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m minionListenDo) Unscoped() *minionListenDo {
	return m.withDO(m.DO.Unscoped())
}

func (m minionListenDo) Create(values ...*model.MinionListen) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m minionListenDo) CreateInBatches(values []*model.MinionListen, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m minionListenDo) Save(values ...*model.MinionListen) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m minionListenDo) First() (*model.MinionListen, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionListen), nil
	}
}

func (m minionListenDo) Take() (*model.MinionListen, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionListen), nil
	}
}

func (m minionListenDo) Last() (*model.MinionListen, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionListen), nil
	}
}

func (m minionListenDo) Find() ([]*model.MinionListen, error) {
	result, err := m.DO.Find()
	return result.([]*model.MinionListen), err
}

func (m minionListenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinionListen, err error) {
	buf := make([]*model.MinionListen, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m minionListenDo) FindInBatches(result *[]*model.MinionListen, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m minionListenDo) Attrs(attrs ...field.AssignExpr) *minionListenDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m minionListenDo) Assign(attrs ...field.AssignExpr) *minionListenDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m minionListenDo) Joins(fields ...field.RelationField) *minionListenDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m minionListenDo) Preload(fields ...field.RelationField) *minionListenDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m minionListenDo) FirstOrInit() (*model.MinionListen, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionListen), nil
	}
}

func (m minionListenDo) FirstOrCreate() (*model.MinionListen, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionListen), nil
	}
}

func (m minionListenDo) FindByPage(offset int, limit int) (result []*model.MinionListen, count int64, err error) {
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

func (m minionListenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m minionListenDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m minionListenDo) Delete(models ...*model.MinionListen) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *minionListenDo) withDO(do gen.Dao) *minionListenDo {
	m.DO = *do.(*gen.DO)
	return m
}
