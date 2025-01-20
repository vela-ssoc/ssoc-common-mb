package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type TaskPublish struct {
	ID           int64           `json:"id,string"      gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	TaskID       int64           `json:"task_id,string" gorm:"column:task_id;index;comment:任务ID"`
	Name         string          `json:"name"           gorm:"column:name;type:varchar(100);comment:名字"`
	Intro        string          `json:"intro"          gorm:"column:intro;type:varchar(1000);comment:简介"`
	Progress     TaskProgress    `json:"progress"       gorm:"column:progress;type:json;comment:执行进度"`
	Finished     bool            `json:"finished"       gorm:"column:finished;comment:是否结束" `
	Code         string          `json:"code"           gorm:"column:code;type:text;notnull;comment:执行代码"`
	CodeSHA1     string          `json:"code_sha1"      gorm:"column:code_sha1;type:char(40);notnull;comment:执行代码SHA1"`
	ContentQuote *ExtensionQuote `json:"content_quote"  gorm:"column:content_quote;type:json;serializer:json;comment:插件引用"`
	Filters      []string        `json:"filters"        gorm:"column:filters;type:json;serializer:json;comment:过滤节点"`
	Excludes     []string        `json:"excludes"       gorm:"column:excludes;serializer:json;comment:排除节点"`
	CreatedBy    Operator        `json:"created_by"     gorm:"column:created_by;type:json;notnull;serializer:json;comment:创建者"`
	CreatedAt    time.Time       `json:"created_at"     gorm:"column:created_at;notnull;default:now(3);comment:创建时间"`
	UpdatedAt    time.Time       `json:"updated_at"     gorm:"column:updated_at;notnull;default:now(3);comment:更新时间"`
}

func (TaskPublish) TableName() string {
	return "task_publish"
}

type TaskProgress struct {
	Total   int64 `json:"total"`   // 节点总数
	Succeed int64 `json:"succeed"` // 成功节点数
	Failed  int64 `json:"failed"`  // 失败节点数
}

func (t *TaskProgress) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, t)
}

func (t TaskProgress) Value() (driver.Value, error) {
	return json.Marshal(t)
}
