package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ZipArchive interface {
	Collection[model.ZipArchive]

	// FindByName 通过文件名查询文件信息。
	FindByName(ctx context.Context, name string, opts ...options.Lister[options.FindOneOptions]) (*model.ZipArchive, error)
}

func NewZipArchive(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) ZipArchive {
	coll := NewCollection[model.ZipArchive](db, opts...)

	return &zipArchive{
		Collection: coll,
	}
}

type zipArchive struct {
	Collection[model.ZipArchive]
}

func (c *zipArchive) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}

func (c *zipArchive) FindByName(ctx context.Context, name string, opts ...options.Lister[options.FindOneOptions]) (*model.ZipArchive, error) {
	filter := bson.D{{Key: "name", Value: name}}
	return c.FindOne(ctx, filter, opts...)
}
