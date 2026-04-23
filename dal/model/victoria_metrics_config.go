package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VictoriaMetricsConfig struct {
	ID        bson.ObjectID `json:"id"         gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string        `json:"name"       gorm:"column:name;size:20;comment:名字"`
	Enabled   bool          `json:"enabled"    gorm:"column:enabled;comment:是否启用"`
	Method    string        `json:"method"     gorm:"column:method;size:10;comment:请求方法"`
	URL       string        `json:"url"        gorm:"column:url;size:255;notnull;comment:"`
	Username  string        `json:"username"   gorm:"column:username;size:255;comment:认证用户名"`
	Password  string        `json:"password"   gorm:"column:username;size:255;comment:认证密码"`
	CreatedAt time.Time     `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}
