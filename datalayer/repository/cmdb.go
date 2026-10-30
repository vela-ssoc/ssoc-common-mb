package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CMDB interface {
	Collection[model.CMDB]
}

func NewCMDB(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) CMDB {
	repo := NewCollection[model.CMDB](db, opts...)

	return &cmdb{
		Collection: repo,
	}
}

type cmdb struct {
	Collection[model.CMDB]
}

func (c *cmdb) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "ci_id", Value: -1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "ci_type", Value: 1}}},
		{Keys: bson.D{{Key: "private_ip", Value: 1}}},
		{Keys: bson.D{{Key: "department", Value: 1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
