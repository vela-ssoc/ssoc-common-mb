package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RedisConfig interface {
	Collection[model.RedisConfig]
	Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.RedisConfig, error)
}

func NewRedisConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) RedisConfig {
	coll := NewCollection[model.RedisConfig](db, opts...)

	return &redisConfig{
		Collection: coll,
	}
}

type redisConfig struct {
	Collection[model.RedisConfig]
}

func (c *redisConfig) Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.RedisConfig, error) {
	return c.FindOne(ctx, bson.D{}, opts...)
}
