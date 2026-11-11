package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AgentConnectHistory interface {
	Collection[model.AgentConnectHistory]
}

func NewAgentConnectHistory(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) AgentConnectHistory {
	coll := NewCollection[model.AgentConnectHistory](db, opts...)

	return &agentConnectHistory{
		Collection: coll,
	}
}

type agentConnectHistory struct {
	Collection[model.AgentConnectHistory]
}

func (c *agentConnectHistory) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "agent_idv1", Value: 1}, {Key: "disconnected_at", Value: -1}}},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}
