package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProgressTask struct {
	ID          bson.ObjectID          `bson:"_id,omitempty"         json:"id"`
	Type        string                 `bson:"type"                  json:"type"`
	Name        string                 `bson:"name"                  json:"name"`
	Description string                 `bson:"description,omitempty" json:"description,omitzero"`
	Statistics  ProgressTaskStatistics `bson:"progress"              json:"progress"`
	Revision    int64                  `bson:"revision"              json:"revision"`
	CreatedBy   Operator               `bson:"created_by"            json:"created_by"`
	CreatedAt   time.Time              `bson:"created_at,omitempty"  json:"created_at"`
	StartedAt   time.Time              `bson:"started_at,omitempty"  json:"started_at,omitzero"`
	UpdatedAt   time.Time              `bson:"updated_at,omitempty"  json:"updated_at"`
	FinishedAt  time.Time              `bson:"finished_at,omitempty" json:"finished_at,omitzero"`
}

func (ProgressTask) CollectionInfo() (string, string) {
	return "progress_task", "任务进度"
}

type ProgressTaskStatistics struct {
	Total    int64 `bson:"total"    json:"total"`    // 总数量
	Success  int64 `bson:"success"  json:"success"`  // 成功数量
	Failed   int64 `bson:"failed"   json:"failed"`   // 失败数量
	Canceled int64 `bson:"canceled" json:"canceled"` // 取消数量
}
