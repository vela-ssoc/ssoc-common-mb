package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// TaskExtension 任务镜像。
type TaskExtension struct {
	ID            int64              `json:"id,string"                gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name          string             `json:"name"                     gorm:"column:name;type:varchar(100);unique;comment:名字"`
	Intro         string             `json:"intro"                    gorm:"column:intro;type:varchar(1000);comment:简介"`
	Code          string             `json:"code"                     gorm:"column:code;type:text;notnull;comment:执行代码"`
	CodeSHA1      string             `json:"code_sha1"                gorm:"column:code_sha1;type:char(40);notnull;comment:执行代码SHA1"`
	ContentQuote  *ExtensionQuote    `json:"content_quote"            gorm:"column:content_quote;serializer:json;comment:插件引用"`
	StepDone      bool               `json:"step_done"                gorm:"column:released;comment:步骤完成"`
	Enabled       bool               `json:"enabled"                  gorm:"column:enabled;comment:开启任务"`
	Cron          string             `json:"cron,omitempty"           gorm:"column:cron;size:50;comment:定时任务表达式"`
	SpecificTimes []time.Time        `json:"specific_times,omitempty" gorm:"column:specific_times;type:json;serializer:json;comment:定点任务时间"`
	Timeout       Duration           `json:"timeout"                  gorm:"column:timeout;serializer:json;comment:超时时间"`
	PushSize      int                `json:"push_size"                gorm:"column:push_size;comment:推送并发数"`
	Filters       []string           `json:"filters"                  gorm:"column:filters;serializer:json;comment:过滤节点"`
	Excludes      []string           `json:"excludes"                 gorm:"column:excludes;serializer:json;comment:排除节点"`
	ExecID        int64              `json:"exec_id,string"           gorm:"column:exec_id;comment:执行ID"`
	Status        *TaskExecuteStatus `json:"status"                   gorm:"column:status;type:json;comment:最近一次任务状态"`
	Finished      bool               `json:"finished"                 gorm:"column:finished;comment:最近一次任务是否结束"`
	CreatedBy     Operator           `json:"created_by"               gorm:"column:created_by;type:json;notnull;serializer:json;comment:创建者"`
	UpdatedBy     Operator           `json:"updated_by"               gorm:"column:updated_by;type:json;notnull;serializer:json;comment:更新者"`
	CreatedAt     time.Time          `json:"created_at"               gorm:"column:created_at;notnull;default:now(3);comment:创建时间"`
	UpdatedAt     time.Time          `json:"updated_at"               gorm:"column:updated_at;notnull;default:now(3);comment:更新时间"`
}

func (TaskExtension) TableName() string {
	return "task_extension"
}

type ExtensionQuote struct {
	ID          int64           `json:"id,string"    gorm:"column:id"`      // 插件商店 ID
	Name        string          `json:"name"         gorm:"column:name"`    // 插件名字
	Version     int64           `json:"version"      gorm:"column:version"` // 引用时的版本
	Data        json.RawMessage `json:"data"         gorm:"column:data;type:json;serializer:json;comment:模板渲染参数"`
	Content     string          `json:"content"      gorm:"column:content;type:text;notnull;comment:代码"`
	ContentSHA1 string          `json:"content_sha1" gorm:"column:content_sha1;type:char(40);notnull;comment:SHA1"`
	CreatedBy   Operator        `json:"created_by"   gorm:"column:created_by;serializer:json"` // 插件创建者
	UpdatedBy   Operator        `json:"updated_by"   gorm:"column:updated_by;serializer:json"` // 引用时插件的修改者
}

func (e *ExtensionQuote) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, e)
}

func (e ExtensionQuote) Value() (driver.Value, error) {
	return json.Marshal(e)
}

type TaskStatus struct {
	Total   int `json:"total"   gorm:"column:total"`
	Succeed int `json:"succeed" gorm:"column:succeed"`
	Failed  int `json:"failed"  gorm:"column:failed"`
}

func (ts TaskStatus) Finished() bool {
	return ts.Succeed+ts.Failed >= ts.Total
}

func (ts *TaskStatus) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, ts)
}

func (ts TaskStatus) Value() (driver.Value, error) {
	return json.Marshal(ts)
}

type name struct {
	Keyword  string   `json:"keyword"`
	Tags     []string `json:"tags"`
	MatchTag bool     `json:"match_tag"`
}

type DFF struct {
	Field string `json:"field"`
}
