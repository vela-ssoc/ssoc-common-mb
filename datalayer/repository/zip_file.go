package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ZipFile interface {
	Collection[model.ZipFile]
	FindFile(ctx context.Context, archiveID bson.ObjectID, fullName string, opts ...options.Lister[options.FindOneOptions]) (*model.ZipFile, error)
}

func NewZipFile(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) ZipFile {
	coll := NewCollection[model.ZipFile](db, opts...)

	return &zipFile{
		Collection: coll,
	}
}

type zipFile struct {
	Collection[model.ZipFile]
}

func (c *zipFile) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "archive_id", Value: 1}, {Key: "full_name", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return c.Indexes().CreateMany(ctx, indexes, opts...)
}

func (c *zipFile) FindFile(ctx context.Context, archiveID bson.ObjectID, fullName string, opts ...options.Lister[options.FindOneOptions]) (*model.ZipFile, error) {
	filter := bson.D{{Key: "archive_id", Value: archiveID}, {Key: "full_name", Value: fullName}}
	return c.FindOne(ctx, filter, opts...)
}
