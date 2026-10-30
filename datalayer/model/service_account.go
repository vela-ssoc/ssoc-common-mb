package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ServiceAccount struct {
	ID          bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Name        string        `bson:"name"                 json:"name"`
	Description string        `bson:"description"          json:"description"`
	Token       string        `bson:"token"                json:"token"`
	Enabled     string        `bson:"enabled"              json:"enabled"`
	ExpiresAt   time.Time     `bson:"expires_at"           json:"expires_at,omitzero"`
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}
