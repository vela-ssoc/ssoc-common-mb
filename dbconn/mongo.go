package dbconn

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/connstring"
)

func OpenMDB(uri string) (*mongo.Database, error) {
	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		return nil, err
	}
	if dbname := cs.Database; dbname == "" {
		return nil, errors.New("database name is required")
	}

	opts := options.Client().ApplyURI(uri)
	cli, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	dbname := cs.Database
	db := cli.Database(dbname)

	return db, nil
}
