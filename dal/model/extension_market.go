package model

import "time"

// ExtensionMarket 插件市场。
type ExtensionMarket struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;type:varchar(100);comment:名字"`
	Author    string    `json:"author"     gorm:"column:author;type:varchar(50);comment:作者"`
	Semver    string    `json:"semver"     gorm:"column:semver;type:varchar(50);comment:版本号"`
	Category  string    `json:"category"   gorm:"column:category;type:varchar(10);comment:插件类型"` // service:服务插件 task:任务插件
	Intro     string    `json:"intro"      gorm:"column:intro;type:varchar(1000);comment:简介"`
	Document  string    `json:"document"   gorm:"column:document;type:text;comment:文档"`
	Content   string    `json:"content"    gorm:"column:content;type:text;comment:代码"`
	CreatedAt time.Time `json:"created_at" gorm:"column:updated_at;not null;default:now(3);comment:更新时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:created_at;not null;default:now(3);comment:创建时间"`
}

func (ExtensionMarket) TableName() string {
	return "extension_market"
}
