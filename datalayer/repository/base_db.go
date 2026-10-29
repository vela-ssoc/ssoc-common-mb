package repository

import (
	"context"
	"log/slog"
	"reflect"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BaseDB struct {
	db  *mongo.Database
	log *slog.Logger

	broker                Broker
	brokerConnectHistory  BrokerConnectHistory
	managerConfig         ManagerConfig
	oauthConfig           OAuthConfig
	reverseProxy          ReverseProxy
	victoriaMetricsConfig VictoriaMetricsConfig
	zipArchive            ZipArchive
	zipFile               ZipFile
}

func NewBaseDB(db *mongo.Database, log *slog.Logger) *BaseDB {
	return &BaseDB{
		db:  db,
		log: log,

		broker:                NewBroker(db),
		brokerConnectHistory:  NewBrokerConnectHistory(db),
		managerConfig:         NewManagerConfig(db),
		oauthConfig:           NewOAuthConfig(db),
		reverseProxy:          NewReverseProxy(db),
		victoriaMetricsConfig: NewVictoriaMetricsConfig(db),
		zipArchive:            NewZipArchive(db),
		zipFile:               NewZipFile(db),
	}
}

func (b *BaseDB) Database() *mongo.Database { return b.db }

func (b *BaseDB) StartSession(opts ...options.Lister[options.SessionOptions]) (*mongo.Session, error) {
	return b.db.Client().StartSession(opts...)
}

func (b *BaseDB) Broker() Broker                               { return b.broker }
func (b *BaseDB) BrokerConnectHistory() BrokerConnectHistory   { return b.brokerConnectHistory }
func (b *BaseDB) ManagerConfig() ManagerConfig                 { return b.managerConfig }
func (b *BaseDB) OAuthConfig() OAuthConfig                     { return b.oauthConfig }
func (b *BaseDB) ReverseProxy() ReverseProxy                   { return b.reverseProxy }
func (b *BaseDB) VictoriaMetricsConfig() VictoriaMetricsConfig { return b.victoriaMetricsConfig }
func (b *BaseDB) ZipArchive() ZipArchive                       { return b.zipArchive }
func (b *BaseDB) ZipFile() ZipFile                             { return b.zipFile }

func (b *BaseDB) CreateIndex(ctx context.Context) error {
	rv := reflect.ValueOf(b)
	for _, mv := range rv.Methods() {
		idx := b.reflectIndexer(mv)
		if idx == nil {
			continue
		}

		name, description := idx.CollectionInfo()
		args := []any{"name", name, "description", description}
		b.log.DebugContext(ctx, "准备创建索引", args...)

		indexes, err := idx.CreateIndex(ctx)
		if err != nil {
			args = append(args, "err", err)
			b.log.ErrorContext(ctx, "索引创建错误", args...)
			return err
		}

		if len(indexes) == 0 {
			b.log.Debug("该集合无需创建索引", args...)
		} else {
			args = append(args, "indexes", indexes)
			b.log.Debug("集合索引创建完毕", args...)
		}
	}

	return nil
}

func (b *BaseDB) reflectIndexer(mv reflect.Value) *collectionIndexProxy {
	if mt := mv.Type(); mt.NumIn() != 0 || mt.NumOut() != 1 {
		return nil
	}

	rets := mv.Call([]reflect.Value{})
	if len(rets) != 1 {
		return nil
	}
	ret := rets[0]
	if ret.IsNil() || !ret.IsValid() {
		return nil
	}

	val := ret.Interface()
	ic, ok := val.(CreateIndexer)
	if !ok {
		return nil
	}

	pxy := &collectionIndexProxy{index: ic}
	if inf, yes := val.(CollectionInfer); yes {
		pxy.name, pxy.comment = inf.CollectionInfo()
	} else {
		name := ret.Type().Name()
		if ni, match := val.(interface{ Name() string }); match {
			name = ni.Name()
		}
		pxy.name = name
	}

	return pxy
}

type collectionIndexProxy struct {
	name    string
	comment string
	index   CreateIndexer
}

func (c *collectionIndexProxy) CollectionInfo() (string, string) {
	return c.name, c.comment
}

func (c *collectionIndexProxy) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	return c.index.CreateIndex(ctx, opts...)
}
