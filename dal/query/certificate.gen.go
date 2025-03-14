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

func newCertificate(db *gorm.DB, opts ...gen.DOOption) certificate {
	_certificate := certificate{}

	_certificate.certificateDo.UseDB(db, opts...)
	_certificate.certificateDo.UseModel(&model.Certificate{})

	tableName := _certificate.certificateDo.TableName()
	_certificate.ALL = field.NewAsterisk(tableName)
	_certificate.ID = field.NewInt64(tableName, "id")
	_certificate.Name = field.NewString(tableName, "name")
	_certificate.Certificate = field.NewString(tableName, "certificate")
	_certificate.PrivateKey = field.NewString(tableName, "private_key")
	_certificate.Version = field.NewInt(tableName, "version")
	_certificate.IssCountry = field.NewField(tableName, "iss_country")
	_certificate.IssProvince = field.NewField(tableName, "iss_province")
	_certificate.IssOrg = field.NewField(tableName, "iss_org")
	_certificate.IssCN = field.NewString(tableName, "iss_cn")
	_certificate.IssOrgUnit = field.NewField(tableName, "iss_org_unit")
	_certificate.SubCountry = field.NewField(tableName, "sub_country")
	_certificate.SubOrg = field.NewField(tableName, "sub_org")
	_certificate.SubProvince = field.NewField(tableName, "sub_province")
	_certificate.SubCN = field.NewString(tableName, "sub_cn")
	_certificate.DNSNames = field.NewField(tableName, "dns_names")
	_certificate.IPAddresses = field.NewField(tableName, "ip_addresses")
	_certificate.EmailAddresses = field.NewField(tableName, "email_addresses")
	_certificate.URIs = field.NewField(tableName, "uris")
	_certificate.NotBefore = field.NewTime(tableName, "not_before")
	_certificate.NotAfter = field.NewTime(tableName, "not_after")
	_certificate.CreatedAt = field.NewTime(tableName, "created_at")
	_certificate.UpdatedAt = field.NewTime(tableName, "updated_at")

	_certificate.fillFieldMap()

	return _certificate
}

type certificate struct {
	certificateDo certificateDo

	ALL            field.Asterisk
	ID             field.Int64
	Name           field.String
	Certificate    field.String
	PrivateKey     field.String
	Version        field.Int
	IssCountry     field.Field
	IssProvince    field.Field
	IssOrg         field.Field
	IssCN          field.String
	IssOrgUnit     field.Field
	SubCountry     field.Field
	SubOrg         field.Field
	SubProvince    field.Field
	SubCN          field.String
	DNSNames       field.Field
	IPAddresses    field.Field
	EmailAddresses field.Field
	URIs           field.Field
	NotBefore      field.Time
	NotAfter       field.Time
	CreatedAt      field.Time
	UpdatedAt      field.Time

	fieldMap map[string]field.Expr
}

func (c certificate) Table(newTableName string) *certificate {
	c.certificateDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c certificate) As(alias string) *certificate {
	c.certificateDo.DO = *(c.certificateDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *certificate) updateTableName(table string) *certificate {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Certificate = field.NewString(table, "certificate")
	c.PrivateKey = field.NewString(table, "private_key")
	c.Version = field.NewInt(table, "version")
	c.IssCountry = field.NewField(table, "iss_country")
	c.IssProvince = field.NewField(table, "iss_province")
	c.IssOrg = field.NewField(table, "iss_org")
	c.IssCN = field.NewString(table, "iss_cn")
	c.IssOrgUnit = field.NewField(table, "iss_org_unit")
	c.SubCountry = field.NewField(table, "sub_country")
	c.SubOrg = field.NewField(table, "sub_org")
	c.SubProvince = field.NewField(table, "sub_province")
	c.SubCN = field.NewString(table, "sub_cn")
	c.DNSNames = field.NewField(table, "dns_names")
	c.IPAddresses = field.NewField(table, "ip_addresses")
	c.EmailAddresses = field.NewField(table, "email_addresses")
	c.URIs = field.NewField(table, "uris")
	c.NotBefore = field.NewTime(table, "not_before")
	c.NotAfter = field.NewTime(table, "not_after")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *certificate) WithContext(ctx context.Context) *certificateDo {
	return c.certificateDo.WithContext(ctx)
}

func (c certificate) TableName() string { return c.certificateDo.TableName() }

func (c certificate) Alias() string { return c.certificateDo.Alias() }

func (c certificate) Columns(cols ...field.Expr) gen.Columns { return c.certificateDo.Columns(cols...) }

func (c *certificate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *certificate) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 22)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["certificate"] = c.Certificate
	c.fieldMap["private_key"] = c.PrivateKey
	c.fieldMap["version"] = c.Version
	c.fieldMap["iss_country"] = c.IssCountry
	c.fieldMap["iss_province"] = c.IssProvince
	c.fieldMap["iss_org"] = c.IssOrg
	c.fieldMap["iss_cn"] = c.IssCN
	c.fieldMap["iss_org_unit"] = c.IssOrgUnit
	c.fieldMap["sub_country"] = c.SubCountry
	c.fieldMap["sub_org"] = c.SubOrg
	c.fieldMap["sub_province"] = c.SubProvince
	c.fieldMap["sub_cn"] = c.SubCN
	c.fieldMap["dns_names"] = c.DNSNames
	c.fieldMap["ip_addresses"] = c.IPAddresses
	c.fieldMap["email_addresses"] = c.EmailAddresses
	c.fieldMap["uris"] = c.URIs
	c.fieldMap["not_before"] = c.NotBefore
	c.fieldMap["not_after"] = c.NotAfter
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c certificate) clone(db *gorm.DB) certificate {
	c.certificateDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c certificate) replaceDB(db *gorm.DB) certificate {
	c.certificateDo.ReplaceDB(db)
	return c
}

type certificateDo struct{ gen.DO }

func (c certificateDo) Debug() *certificateDo {
	return c.withDO(c.DO.Debug())
}

func (c certificateDo) WithContext(ctx context.Context) *certificateDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c certificateDo) ReadDB() *certificateDo {
	return c.Clauses(dbresolver.Read)
}

func (c certificateDo) WriteDB() *certificateDo {
	return c.Clauses(dbresolver.Write)
}

func (c certificateDo) Session(config *gorm.Session) *certificateDo {
	return c.withDO(c.DO.Session(config))
}

func (c certificateDo) Clauses(conds ...clause.Expression) *certificateDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c certificateDo) Returning(value interface{}, columns ...string) *certificateDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c certificateDo) Not(conds ...gen.Condition) *certificateDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c certificateDo) Or(conds ...gen.Condition) *certificateDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c certificateDo) Select(conds ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c certificateDo) Where(conds ...gen.Condition) *certificateDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c certificateDo) Order(conds ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c certificateDo) Distinct(cols ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c certificateDo) Omit(cols ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c certificateDo) Join(table schema.Tabler, on ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c certificateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *certificateDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c certificateDo) RightJoin(table schema.Tabler, on ...field.Expr) *certificateDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c certificateDo) Group(cols ...field.Expr) *certificateDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c certificateDo) Having(conds ...gen.Condition) *certificateDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c certificateDo) Limit(limit int) *certificateDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c certificateDo) Offset(offset int) *certificateDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c certificateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *certificateDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c certificateDo) Unscoped() *certificateDo {
	return c.withDO(c.DO.Unscoped())
}

func (c certificateDo) Create(values ...*model.Certificate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c certificateDo) CreateInBatches(values []*model.Certificate, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c certificateDo) Save(values ...*model.Certificate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c certificateDo) First() (*model.Certificate, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Certificate), nil
	}
}

func (c certificateDo) Take() (*model.Certificate, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Certificate), nil
	}
}

func (c certificateDo) Last() (*model.Certificate, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Certificate), nil
	}
}

func (c certificateDo) Find() ([]*model.Certificate, error) {
	result, err := c.DO.Find()
	return result.([]*model.Certificate), err
}

func (c certificateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Certificate, err error) {
	buf := make([]*model.Certificate, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c certificateDo) FindInBatches(result *[]*model.Certificate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c certificateDo) Attrs(attrs ...field.AssignExpr) *certificateDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c certificateDo) Assign(attrs ...field.AssignExpr) *certificateDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c certificateDo) Joins(fields ...field.RelationField) *certificateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c certificateDo) Preload(fields ...field.RelationField) *certificateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c certificateDo) FirstOrInit() (*model.Certificate, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Certificate), nil
	}
}

func (c certificateDo) FirstOrCreate() (*model.Certificate, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Certificate), nil
	}
}

func (c certificateDo) FindByPage(offset int, limit int) (result []*model.Certificate, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c certificateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c certificateDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c certificateDo) Delete(models ...*model.Certificate) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *certificateDo) withDO(do gen.Dao) *certificateDo {
	c.DO = *do.(*gen.DO)
	return c
}
