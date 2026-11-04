package repository

import (
	"context"
	"crypto/tls"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Certificate interface {
	Collection[model.Certificate]
	Certificates(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) ([]*tls.Certificate, error)
}

func NewCertificate(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Certificate {
	repo := NewCollection[model.Certificate](db, opts...)

	return &certificate{
		Collection: repo,
	}
}

type certificate struct {
	Collection[model.Certificate]
}

// Certificates 所有证书
func (c *certificate) Certificates(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) ([]*tls.Certificate, error) {
	dats, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	rets := make([]*tls.Certificate, 0, len(dats))
	for _, dat := range dats {
		pair, err1 := tls.X509KeyPair([]byte(dat.PublicKey), []byte(dat.PrivateKey))
		if err1 != nil {
			return nil, err1
		}
		rets = append(rets, &pair)
	}

	return rets, nil
}
