package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CMDBConfig interface {
	Collection[model.CMDBConfig]
	Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.CMDBConfig, error)
}

func NewCMDBConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) CMDBConfig {
	repo := NewCollection[model.CMDBConfig](db, opts...)

	return &cmdbConfig{
		Collection: repo,
	}
}

type cmdbConfig struct {
	Collection[model.CMDBConfig]
}

func (c *cmdbConfig) Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.CMDBConfig, error) {
	return c.FindOne(ctx, bson.D{}, opts...)
}
