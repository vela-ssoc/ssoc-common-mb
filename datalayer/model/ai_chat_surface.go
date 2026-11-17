package model

import (
	"time"

	"github.com/tmc/a2ui"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AIChatSurface struct {
	ID            bson.ObjectID    `bson:"_id,omitempty"             json:"id"`
	ThreadID      string           `bson:"thread_id"                 json:"thread_id"`
	SurfaceID     string           `bson:"surface_id"                json:"surface_id"`
	Version       string           `bson:"version"                   json:"version"`
	CatalogID     string           `bson:"catalog_id,omitempty"      json:"catalog_id,omitempty"`
	Components    []a2ui.Component `bson:"components"                json:"components"`
	DataModel     bson.M           `bson:"data_model,omitempty"      json:"data_model,omitempty"`
	SendDataModel bool             `bson:"send_data_model,omitempty" json:"send_data_model,omitempty"`
	Metadata      bson.M           `bson:"metadata,omitempty"        json:"metadata,omitempty"`
	Revision      int64            `bson:"revision"                  json:"revision"`
	Operator      Operator         `bson:"operator"                  json:"operator"`
	CreatedAt     time.Time        `bson:"created_at,omitempty"      json:"created_at"`
	UpdatedAt     time.Time        `bson:"updated_at,omitempty"      json:"updated_at"`
}

func (AIChatSurface) CollectionInfo() (string, string) {
	return "ai_chat_surface", "AI 对话画布"
}
