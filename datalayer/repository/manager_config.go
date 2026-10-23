package repository

import (
	"context"
	"errors"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ManagerConfig interface {
	Collection[model.ManagerConfig]
	Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.ManagerConfig, error)
}

func NewManagerConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) ManagerConfig {
	coll := NewCollection[model.ManagerConfig](db, opts...)

	return &managerConfig{
		Collection: coll,
	}
}

type managerConfig struct {
	Collection[model.ManagerConfig]
}

func (c *managerConfig) Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.ManagerConfig, error) {
	dat, err := c.FindOne(ctx, bson.D{}, opts...)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		dat = new(model.ManagerConfig)
	}
	if dat.Server == nil {
		dat.Server = new(model.ManagerServerConfig)
	}
	if dat.Logger == nil {
		dat.Logger = new(model.LoggerConfig)
	}

	return dat, nil
}
