package model

import "time"

// ExtensionMarket 插件市场。
type ExtensionMarket struct {
	ID          int64     `json:"id,string"    gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name        string    `json:"name"         gorm:"column:name;type:varchar(100);unique;comment:名字"`
	Intro       string    `json:"intro"        gorm:"column:intro;type:varchar(1000);comment:简介"`
	Category    string    `json:"category"     gorm:"column:category;type:varchar(10);comment:插件类型"` // service:服务插件 task:任务插件
	Version     int64     `json:"version"      gorm:"column:version;comment:版本号"`
	Content     string    `json:"content"      gorm:"column:content;type:text;notnull;comment:代码"`
	ContentSHA1 string    `json:"content_sha1" gorm:"column:content_sha1;type:char(40);notnull;comment:SHA1"`
	Changelog   string    `json:"changelog"    gorm:"column:changelog;type:text;comment:更新日志"`
	CreatedBy   Operator  `json:"created_by"   gorm:"column:created_by;type:json;notnull;serializer:json;comment:创建者"`
	UpdatedBy   Operator  `json:"updated_by"   gorm:"column:updated_by;type:json;notnull;serializer:json;comment:修改者"`
	CreatedAt   time.Time `json:"created_at"   gorm:"column:updated_at;notnull;default:now(3);comment:修改时间"`
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:created_at;notnull;default:now(3);comment:创建时间"`
}

func (ExtensionMarket) TableName() string {
	return "extension_market"
}
