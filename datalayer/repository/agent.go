package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Agent interface {
	Collection[model.Agent]
}

func NewAgent(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Agent {
	coll := NewCollection[model.Agent](db, opts...)

	return &agent{
		Collection: coll,
	}
}

type agent struct {
	Collection[model.Agent]
}

func (c *agent) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "idv1", Value: 1}}},
		{Keys: bson.D{{Key: "machine_id", Value: 1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
