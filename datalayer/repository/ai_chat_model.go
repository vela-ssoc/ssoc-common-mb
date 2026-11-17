package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AIChatModel interface {
	Collection[model.AIChatModel]
	Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.AIChatModel, error)
}

func NewAIChatModel(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) AIChatModel {
	coll := NewCollection[model.AIChatModel](db, opts...)

	return &aiChatModel{
		Collection: coll,
	}
}

type aiChatModel struct {
	Collection[model.AIChatModel]
}

func (c aiChatModel) Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.AIChatModel, error) {
	return c.FindOne(ctx, bson.D{}, opts...)
}
