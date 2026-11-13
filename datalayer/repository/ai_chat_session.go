package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AIChatSession interface {
	Collection[model.AIChatSession]
}

func NewAIChatSession(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) AIChatSession {
	coll := NewCollection[model.AIChatSession](db, opts...)

	return &aiChatSession{
		Collection: coll,
	}
}

type aiChatSession struct {
	Collection[model.AIChatSession]
}

func (c aiChatSession) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "session_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "session_id", Value: 1}, {Key: "updated_at", Value: -1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
