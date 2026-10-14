package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ReverseProxy interface {
	Collection[model.ReverseProxy]
}

func NewReverseProxy(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) ReverseProxy {
	coll := NewCollection[model.ReverseProxy](db, opts...)

	return &reverseProxy{
		Collection: coll,
	}
}

type reverseProxy struct {
	Collection[model.ReverseProxy]
}

func (c *reverseProxy) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "location", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
