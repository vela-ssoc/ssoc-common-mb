package sqldb

import (
	"database/sql"
	"log/slog"

	"gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/go-sql-driver/mysql"
	"github.com/vela-ssoc/opengauss"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open 连接数据库，并检测是否是信创 OpenGauss。
func Open(dsn string, driverLog *slog.Logger, opts ...gorm.Option) (*gorm.DB, bool, error) {
	// 区分 MySQL 还是 OpenGauss，用的是 parse dsn 的方式。
	// 注意：用该方式区分必须要先用 [mysql.ParseDSN] 才较为稳妥。
	if cfg, err := mysql.ParseDSN(dsn); err == nil && cfg != nil {
		cfg.Logger = &mysqlLog{log: driverLog}
		dia := &gormysql.Dialector{
			Config: &gormysql.Config{DSN: dsn, DSNConfig: cfg},
		}
		db, exx := gorm.Open(dia, opts...)
		if exx != nil {
			return nil, false, exx
		}
		exx = autoPrimaryKey(db)

		return db, false, exx
	}

	cfg, err := pq.ParseConfig(dsn)
	if err != nil {
		return nil, false, err
	}

	cfg.Logger = &gaussLog{log: driverLog}
	connector, exx := pq.NewConnectorConfig(cfg)
	if exx != nil {
		return nil, true, exx
	}
	conn := sql.OpenDB(connector)
	gcfg := opengauss.Config{DSN: dsn, Conn: conn}
	dia := opengauss.New(gcfg)
	db, exx := gorm.Open(dia, opts...)
	if exx != nil {
		return nil, false, exx
	}
	exx = autoPrimaryKey(db)

	return db, true, err
}

func autoPrimaryKey(db *gorm.DB) error {
	// 雪花 ID 生成器
	sn := newSnow()
	return db.Callback().
		Create().
		Before("gorm:create").
		Register("generate_id", sn.autoID)
}
