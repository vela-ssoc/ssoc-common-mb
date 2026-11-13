package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AIChatSession struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	UserID    bson.ObjectID `bson:"user_id"              json:"user_id"`
	Title     string        `bson:"title"                json:"title"`
	SessionID string        `bson:"session_id"           json:"session_id"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (AIChatSession) CollectionInfo() (string, string) {
	return "ai_chat_session", "AI 对话 session"
}
