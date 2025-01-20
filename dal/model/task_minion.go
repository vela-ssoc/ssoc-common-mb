package model

import (
	"encoding/json"
	"time"
)

type TaskMinion struct {
	ID           int64            `json:"id,string"`
	MinionID     int64            `json:"minion_id,string"`
	TaskID       int64            `json:"task_id,string"`
	PublishID    int64            `json:"publish_id,string"`
	TaskName     string           `json:"task_name"`
	ManagerPhase *TaskMinionPhase `json:"manager_phase"`
	BrokerPhase  *TaskMinionPhase `json:"broker_phase"`
	Result       json.RawMessage  `json:"result"`
	Succeed      bool             `json:"succeed"`
	Finished     bool             `json:"finished"`
	CreatedBy    Operator         `json:"created_by"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

func (TaskMinion) TableName() string {
	return "task_minion"
}

type TaskMinionPhase struct {
	CreatedAt time.Time `json:"created_at"` // 阶段触发时间
	Succeed   bool      `json:"succeed"`    // 该阶段是否执行成功
	Reason    string    `json:"reason"`     // 失败原因
}
