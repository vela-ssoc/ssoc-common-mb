package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BrokerConnectHistory interface {
	Collection[model.BrokerConnectHistory]
}

func NewBrokerConnectHistory(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) BrokerConnectHistory {
	coll := NewCollection[model.BrokerConnectHistory](db, opts...)

	return &brokerConnectHistory{
		Collection: coll,
	}
}

type brokerConnectHistory struct {
	Collection[model.BrokerConnectHistory]
}

func (c *brokerConnectHistory) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "broker_id", Value: 1}, {Key: "disconnected_at", Value: -1}}},
		{Keys: bson.D{{Key: "slug", Value: 1}, {Key: "disconnected_at", Value: -1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
