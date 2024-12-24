package sqldb

import (
	"database/sql"
	"log/slog"

	"gitee.com/opengauss/openGauss-connector-go-pq"
	msql "github.com/go-sql-driver/mysql"
	"github.com/ssoc/opengauss"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open 连接数据库，并返回是否是信创 OpenGauss，否就是 MySQL。
func Open(dsn string, driverLog *slog.Logger, opts ...gorm.Option) (*gorm.DB, bool, error) {
	if cfg, err := msql.ParseDSN(dsn); err == nil && cfg != nil {
		cfg.Logger = &mysqlLog{log: driverLog}
		dia := &mysql.Dialector{
			Config: &mysql.Config{DSN: dsn, DSNConfig: cfg},
		}
		db, err := gorm.Open(dia, opts...)

		return db, false, err
	}

	cfg, err := pq.ParseConfig(dsn)
	if err != nil {
		return nil, false, err
	}

	cfg.Logger = &gaussLog{log: driverLog}
	connector, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		return nil, true, err
	}
	conn := sql.OpenDB(connector)
	gcfg := opengauss.Config{
		DSN:  dsn,
		Conn: conn,
	}
	dia := opengauss.New(gcfg)
	db, err := gorm.Open(dia, opts...)

	return db, true, err
}
