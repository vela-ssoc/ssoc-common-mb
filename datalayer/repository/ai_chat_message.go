package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AIChatMessage interface {
	Collection[model.AIChatMessage]
	Messages(ctx context.Context, userID bson.ObjectID, sessID string, opts ...options.Lister[options.FindOptions]) ([]*model.AIChatMessage, error)
}

func NewAIChatMessage(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) AIChatMessage {
	coll := NewCollection[model.AIChatMessage](db, opts...)

	return &aiChatMessage{
		Collection: coll,
	}
}

type aiChatMessage struct {
	Collection[model.AIChatMessage]
}

func (c aiChatMessage) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "session_id", Value: 1}, {Key: "created_at", Value: -1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}

func (c aiChatMessage) Messages(ctx context.Context, userID bson.ObjectID, sessID string, opts ...options.Lister[options.FindOptions]) ([]*model.AIChatMessage, error) {
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "session_id", Value: sessID}}
	return c.Find(ctx, filter, opts...)
}
