package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Broker interface {
	Collection[model.Broker]
}

func NewBroker(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Broker {
	coll := NewCollection[model.Broker](db, opts...)

	return &broker{
		Collection: coll,
	}
}

type broker struct {
	Collection[model.Broker]
}

func (c *broker) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "secret", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
