package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User interface {
	Collection[model.User]
}

func NewUser(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) User {
	repo := NewCollection[model.User](db, opts...)

	return &user{
		Collection: repo,
	}
}

type user struct {
	Collection[model.User]
}

func (c *user) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "username", Value: 1}, {Key: "company", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "session.token", Value: 1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
