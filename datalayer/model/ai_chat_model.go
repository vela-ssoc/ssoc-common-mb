package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AIChatModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Enabled   bool          `bson:"enabled"              json:"enabled"`
	BaseURL   string        `bson:"base_url"             json:"base_url"`
	APIKey    string        `bson:"api_key"              json:"api_key"`
	MaxStep   int           `bson:"max_step"             json:"max_step"`
	Models    []string      `bson:"models"               json:"models"` // 允许使用的模型
	Prompt    string        `bson:"prompt"               json:"prompt"` // 默认系统提示词
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (AIChatModel) CollectionInfo() (string, string) { return "ai_chat_model", "AI 对话模型配置" }

func (m AIChatModel) IfMaxStep(fallback int) int {
	if m.MaxStep > 0 {
		return m.MaxStep
	}

	return fallback
}
