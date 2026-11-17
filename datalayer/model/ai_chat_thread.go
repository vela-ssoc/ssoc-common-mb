package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AIChatThread struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	ThreadID  string        `bson:"thread_id"            json:"thread_id"`
	Title     string        `bson:"title"                json:"title"`
	Operator  Operator      `bson:"operator"             json:"operator"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (AIChatThread) CollectionInfo() (string, string) {
	return "ai_chat_thread", "AI 对话"
}
