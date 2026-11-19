package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProgressTaskStatus string

const (
	TaskStatusRunning  ProgressTaskStatus = "running"  // 执行中
	TaskStatusSuccess  ProgressTaskStatus = "success"  // 全部成功
	TaskStatusFailure  ProgressTaskStatus = "failure"  // 任务失败
	TaskStatusCanceled ProgressTaskStatus = "canceled" // 已取消
)

type ProgressTaskDetail struct {
	ID         bson.ObjectID      `bson:"_id,omitempty"         json:"id"`
	Target     ProgressTaskTarget `bson:"target"                json:"target"`
	Status     ProgressTaskStatus `bson:"status"                json:"status"`
	Message    string             `bson:"message,omitempty"     json:"message,omitempty"`
	Input      any                `bson:"input,omitempty"       json:"input"`
	Result     any                `bson:"result,omitempty"      json:"result"`
	StartedAt  time.Time          `bson:"started_at,omitempty"  json:"started_at,omitzero"`
	UpdatedAt  time.Time          `bson:"updated_at"            json:"updated_at"`
	FinishedAt time.Time          `bson:"finished_at,omitempty" json:"finished_at,omitzero"`
}

type ProgressTaskTarget struct {
	ID   bson.ObjectID `bson:"id"   json:"id"`
	Type string        `bson:"type" json:"type"`
	Name string        `bson:"name" json:"name"`
}
