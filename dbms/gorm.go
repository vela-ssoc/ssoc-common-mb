package dbms

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	orm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	MaxOpenConn int               `json:"max_open_conn" yaml:"max_open_conn"`                          // 最大连接数
	MaxIdleConn int               `json:"max_idle_conn" yaml:"max_idle_conn"`                          // 最大空闲连接数
	MaxLifeTime time.Duration     `json:"max_life_time" yaml:"max_life_time"`                          // 连接最大存活时长
	MaxIdleTime time.Duration     `json:"max_idle_time" yaml:"max_idle_time"`                          // 空闲连接最大时长
	DSN         string            `json:"dsn"           yaml:"dsn"`                                    // 数据源
	User        string            `json:"user"          yaml:"user"   validate:"required_without=DSN"` // 数据库用户名
	Passwd      string            `json:"passwd"        yaml:"passwd" validate:"required_without=DSN"` // 密码
	Net         string            `json:"net"           yaml:"net"`                                    // 连接协议
	Addr        string            `json:"addr"          yaml:"addr"   validate:"required_without=DSN"` // 连接地址
	DBName      string            `json:"dbname"        yaml:"dbname" validate:"required_without=DSN"` // 库名
	Params      map[string]string `json:"params"        yaml:"params"`                                 // 参数
}

// FormatDSN 生成数据库连接
func (db Config) FormatDSN() string {
	if dsn := db.DSN; dsn != "" {
		return dsn
	}

	protocol := db.Net
	if protocol == "" {
		protocol = "tcp"
	}
	cfg := &mysql.Config{
		User:   db.User,
		Passwd: db.Passwd,
		Net:    protocol,
		Addr:   db.Addr,
		DBName: db.DBName,
		Params: db.Params,
	}

	return cfg.FormatDSN()
}

// Open 连接数据库
func Open(cfg Config, lgi logger.Interface) (*gorm.DB, *sql.DB, error) {
	dsn := cfg.FormatDSN()

	db, err := gorm.Open(orm.Open(dsn), &gorm.Config{Logger: lgi})
	if err != nil {
		return nil, nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	// ----------[ 设置连接参数 ]----------
	sdb.SetMaxIdleConns(cfg.MaxIdleConn)
	sdb.SetMaxOpenConns(cfg.MaxOpenConn)
	sdb.SetConnMaxLifetime(cfg.MaxLifeTime)
	sdb.SetConnMaxIdleTime(cfg.MaxIdleTime)

	// 雪花 ID 生成器
	sn := newSnow()
	_ = db.Callback().
		Create().
		Before("gorm:create").
		Register("generate_id", sn.plugin)

	return db, sdb, nil
}
