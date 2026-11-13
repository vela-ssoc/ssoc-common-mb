package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	AIRoleUser      = "user"
	AIRoleAssistant = "assistant"
)

type AIChatMessage struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	UserID    bson.ObjectID `bson:"user_id"              json:"user_id"`
	SessionID string        `bson:"session_id"           json:"session_id"`
	Role      string        `bson:"role"                 json:"role"`
	Content   string        `bson:"content"              json:"content"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
}

func (AIChatMessage) CollectionInfo() (string, string) {
	return "ai_chat_message", "AI 对话消息列表"
}
