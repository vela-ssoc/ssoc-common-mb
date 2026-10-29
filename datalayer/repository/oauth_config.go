package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type OAuthConfig interface {
	Collection[model.OAuthConfig]
	Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.OAuthConfig, error)
}

func NewOAuthConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) OAuthConfig {
	coll := NewCollection[model.OAuthConfig](db, opts...)

	return &oauthConfig{
		Collection: coll,
	}
}

type oauthConfig struct {
	Collection[model.OAuthConfig]
}

func (c *oauthConfig) Get(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*model.OAuthConfig, error) {
	return c.FindOne(ctx, bson.D{}, opts...)
}
