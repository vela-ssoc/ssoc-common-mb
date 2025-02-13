package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type TaskExecuteItem struct {
	ID            int64           `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	TaskID        int64           `json:"task_id,string"   gorm:"column:task_id;index:idx_task_id_exec_id;uniqueIndex:idx_task_id_exec_id_minion_id;comment:任务ID"`
	ExecID        int64           `json:"exec_id,string"   gorm:"column:exec_id;index:idx_task_id_exec_id;uniqueIndex:idx_task_id_exec_id_minion_id;comment:执行ID"`
	MinionID      int64           `json:"minion_id,string" gorm:"column:minion_id;uniqueIndex:idx_task_id_exec_id_minion_id;comment:节点ID"`
	Inet          string          `json:"inet"             gorm:"column:inet;size:15;comment:节点IP"`
	BrokerID      int64           `json:"broker_id,string" gorm:"column:broker_id;index:idx_broker_id;comment:代理节点ID"`
	BrokerName    string          `json:"broker_name"      gorm:"column:broker_name;size:255;comment:代理节点名字"`
	ManagerStatus *TaskStepStatus `json:"manager_status"   gorm:"column:manager_status;serializer:json;comment:manager执行状态"`
	BrokerStatus  *TaskStepStatus `json:"broker_status"    gorm:"column:broker_status;serializer:json;comment:broker执行状态"`
	MinionStatus  *TaskStepStatus `json:"minion_status"    gorm:"column:minion_status;serializer:json;comment:agent执行状态"`
	Finished      bool            `json:"finished"         gorm:"column:finished;comment:是否执行完毕"`
	Succeed       bool            `json:"succeed"          gorm:"column:succeed;comment:是否执行成功"`
	Result        json.RawMessage `json:"result"           gorm:"column:result;comment:agent执行结果"`
	CreatedAt     time.Time       `json:"created_at"       gorm:"column:created_at;notnull;default:now(3);comment:创建时间"`
	UpdatedAt     time.Time       `json:"updated_at"       gorm:"column:updated_at;notnull;default:now(3);comment:更新时间"`
}

func (TaskExecuteItem) TableName() string {
	return "task_execute_item"
}

type TaskStepStatus struct {
	Succeed    bool      `json:"succeed"`     // 是否执行成功
	Reason     string    `json:"reason"`      // 如果失败则为失败原因
	ExecutedAt time.Time `json:"executed_at"` // 执行时间
}

func (tss *TaskStepStatus) Value() (driver.Value, error) {
	return json.Marshal(tss)
}

func (tss *TaskStepStatus) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, tss)
}
