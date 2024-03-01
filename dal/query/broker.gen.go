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

func newBroker(db *gorm.DB, opts ...gen.DOOption) broker {
	_broker := broker{}

	_broker.brokerDo.UseDB(db, opts...)
	_broker.brokerDo.UseModel(&model.Broker{})

	tableName := _broker.brokerDo.TableName()
	_broker.ALL = field.NewAsterisk(tableName)
	_broker.ID = field.NewInt64(tableName, "id")
	_broker.Name = field.NewString(tableName, "name")
	_broker.Servername = field.NewString(tableName, "servername")
	_broker.LAN = field.NewField(tableName, "lan")
	_broker.VIP = field.NewField(tableName, "vip")
	_broker.Status = field.NewBool(tableName, "status")
	_broker.Secret = field.NewString(tableName, "secret")
	_broker.CertID = field.NewInt64(tableName, "cert_id")
	_broker.Bind = field.NewString(tableName, "bind")
	_broker.Semver = field.NewString(tableName, "semver")
	_broker.HeartbeatAt = field.NewTime(tableName, "heartbeat_at")
	_broker.CreatedAt = field.NewTime(tableName, "created_at")
	_broker.UpdatedAt = field.NewTime(tableName, "updated_at")

	_broker.fillFieldMap()

	return _broker
}

type broker struct {
	brokerDo brokerDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	Servername  field.String
	LAN         field.Field
	VIP         field.Field
	Status      field.Bool
	Secret      field.String
	CertID      field.Int64
	Bind        field.String
	Semver      field.String
	HeartbeatAt field.Time
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (b broker) Table(newTableName string) *broker {
	b.brokerDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b broker) As(alias string) *broker {
	b.brokerDo.DO = *(b.brokerDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *broker) updateTableName(table string) *broker {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Name = field.NewString(table, "name")
	b.Servername = field.NewString(table, "servername")
	b.LAN = field.NewField(table, "lan")
	b.VIP = field.NewField(table, "vip")
	b.Status = field.NewBool(table, "status")
	b.Secret = field.NewString(table, "secret")
	b.CertID = field.NewInt64(table, "cert_id")
	b.Bind = field.NewString(table, "bind")
	b.Semver = field.NewString(table, "semver")
	b.HeartbeatAt = field.NewTime(table, "heartbeat_at")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")

	b.fillFieldMap()

	return b
}

func (b *broker) WithContext(ctx context.Context) *brokerDo { return b.brokerDo.WithContext(ctx) }

func (b broker) TableName() string { return b.brokerDo.TableName() }

func (b broker) Alias() string { return b.brokerDo.Alias() }

func (b broker) Columns(cols ...field.Expr) gen.Columns { return b.brokerDo.Columns(cols...) }

func (b *broker) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *broker) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 13)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
	b.fieldMap["servername"] = b.Servername
	b.fieldMap["lan"] = b.LAN
	b.fieldMap["vip"] = b.VIP
	b.fieldMap["status"] = b.Status
	b.fieldMap["secret"] = b.Secret
	b.fieldMap["cert_id"] = b.CertID
	b.fieldMap["bind"] = b.Bind
	b.fieldMap["semver"] = b.Semver
	b.fieldMap["heartbeat_at"] = b.HeartbeatAt
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
}

func (b broker) clone(db *gorm.DB) broker {
	b.brokerDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b broker) replaceDB(db *gorm.DB) broker {
	b.brokerDo.ReplaceDB(db)
	return b
}

type brokerDo struct{ gen.DO }

func (b brokerDo) Debug() *brokerDo {
	return b.withDO(b.DO.Debug())
}

func (b brokerDo) WithContext(ctx context.Context) *brokerDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b brokerDo) ReadDB() *brokerDo {
	return b.Clauses(dbresolver.Read)
}

func (b brokerDo) WriteDB() *brokerDo {
	return b.Clauses(dbresolver.Write)
}

func (b brokerDo) Session(config *gorm.Session) *brokerDo {
	return b.withDO(b.DO.Session(config))
}

func (b brokerDo) Clauses(conds ...clause.Expression) *brokerDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b brokerDo) Returning(value interface{}, columns ...string) *brokerDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b brokerDo) Not(conds ...gen.Condition) *brokerDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b brokerDo) Or(conds ...gen.Condition) *brokerDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b brokerDo) Select(conds ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b brokerDo) Where(conds ...gen.Condition) *brokerDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b brokerDo) Order(conds ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b brokerDo) Distinct(cols ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b brokerDo) Omit(cols ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b brokerDo) Join(table schema.Tabler, on ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b brokerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *brokerDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b brokerDo) RightJoin(table schema.Tabler, on ...field.Expr) *brokerDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b brokerDo) Group(cols ...field.Expr) *brokerDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b brokerDo) Having(conds ...gen.Condition) *brokerDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b brokerDo) Limit(limit int) *brokerDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b brokerDo) Offset(offset int) *brokerDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b brokerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *brokerDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b brokerDo) Unscoped() *brokerDo {
	return b.withDO(b.DO.Unscoped())
}

func (b brokerDo) Create(values ...*model.Broker) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b brokerDo) CreateInBatches(values []*model.Broker, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b brokerDo) Save(values ...*model.Broker) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b brokerDo) First() (*model.Broker, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Broker), nil
	}
}

func (b brokerDo) Take() (*model.Broker, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Broker), nil
	}
}

func (b brokerDo) Last() (*model.Broker, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Broker), nil
	}
}

func (b brokerDo) Find() ([]*model.Broker, error) {
	result, err := b.DO.Find()
	return result.([]*model.Broker), err
}

func (b brokerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Broker, err error) {
	buf := make([]*model.Broker, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b brokerDo) FindInBatches(result *[]*model.Broker, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b brokerDo) Attrs(attrs ...field.AssignExpr) *brokerDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b brokerDo) Assign(attrs ...field.AssignExpr) *brokerDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b brokerDo) Joins(fields ...field.RelationField) *brokerDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b brokerDo) Preload(fields ...field.RelationField) *brokerDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b brokerDo) FirstOrInit() (*model.Broker, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Broker), nil
	}
}

func (b brokerDo) FirstOrCreate() (*model.Broker, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Broker), nil
	}
}

func (b brokerDo) FindByPage(offset int, limit int) (result []*model.Broker, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b brokerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b brokerDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b brokerDo) Delete(models ...*model.Broker) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *brokerDo) withDO(do gen.Dao) *brokerDo {
	b.DO = *do.(*gen.DO)
	return b
}
