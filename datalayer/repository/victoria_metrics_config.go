package repository

import (
	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type VictoriaMetricsConfig interface {
	Collection[model.VictoriaMetricsConfig]
}

func NewVictoriaMetricsConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) VictoriaMetricsConfig {
	coll := NewCollection[model.VictoriaMetricsConfig](db, opts...)

	return &victoriaMetricsConfig{
		Collection: coll,
	}
}

type victoriaMetricsConfig struct {
	Collection[model.VictoriaMetricsConfig]
}
