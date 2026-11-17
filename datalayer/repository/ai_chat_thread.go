package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AIChatThread interface {
	Collection[model.AIChatThread]
}

func NewAIChatThread(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) AIChatThread {
	coll := NewCollection[model.AIChatThread](db, opts...)

	return &aiChatThread{
		Collection: coll,
	}
}

type aiChatThread struct {
	Collection[model.AIChatThread]
}

func (c aiChatThread) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "operator.id", Value: 1}, {Key: "thread_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "operator.id", Value: 1}, {Key: "thread_id", Value: 1}, {Key: "updated_at", Value: -1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
