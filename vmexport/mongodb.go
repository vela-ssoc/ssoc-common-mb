package vmexport

import (
	"context"
	"io"
	"log/slog"
	"strconv"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDB struct {
	db  *mongo.Database
	log *slog.Logger
}

func NewMongoDB(db *mongo.Database, log *slog.Logger) *MongoDB {
	return &MongoDB{db: db, log: log}
}

func (m *MongoDB) WriteMetric(ctx context.Context, w io.Writer) {
	m.dbStats(ctx, w)
	m.serverStatus(ctx, w)
}

func (m *MongoDB) dbStats(ctx context.Context, w io.Writer) {
	stats := m.runCommand(ctx, "dbStats")
	dbname, ok := stats["db"].(string)
	if !ok || dbname == "" {
		return
	}
	label := "{db=" + strconv.Quote(dbname) + "}"

	// # HELP mongodb_index_size_bytes Total size of all indexes in bytes
	// # TYPE mongodb_index_size_bytes gauge
	// mongodb_index_size_bytes{db="siem"} 36880384
	if num, yes := stats["indexSize"].(float64); yes {
		name := "mongodb_index_size_bytes" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}

	// # HELP mongodb_storage_size_bytes Total disk space allocated for documents in bytes
	// # TYPE mongodb_storage_size_bytes gauge
	// mongodb_storage_size_bytes{db="siem"} 250179584
	if num, yes := stats["storageSize"].(float64); yes {
		name := "mongodb_storage_size_bytes" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}

	//
	// # HELP mongodb_data_size_bytes Total uncompressed size of BSON data in bytes
	// # TYPE mongodb_data_size_bytes gauge
	// mongodb_data_size_bytes{db="siem"} 1137288527
	if num, yes := stats["dataSize"].(float64); yes {
		name := "mongodb_data_size_bytes" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}
	//
	// # HELP mongodb_objects_total Total number of documents in the database
	// # TYPE mongodb_objects_total gauge
	// mongodb_objects_total{db="siem"} 706054
	if num, yes := stats["objects"].(float64); yes {
		name := "mongodb_objects_total" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}
	//
	// # HELP mongodb_fs_used_bytes Used size of the underlying file system in bytes
	// # TYPE mongodb_fs_used_bytes gauge
	// mongodb_fs_used_bytes{db="siem"} 28819521536
	if num, yes := stats["fsUsedSize"].(float64); yes {
		name := "mongodb_fs_used_bytes" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}
	//
	// # HELP mongodb_fs_total_bytes Total size of the underlying file system in bytes
	// # TYPE mongodb_fs_total_bytes gauge
	// mongodb_fs_total_bytes{db="siem"} 65670217728
	if num, yes := stats["fsTotalSize"].(float64); yes {
		name := "mongodb_fs_total_bytes" + label
		metrics.WriteGaugeFloat64(w, name, num)
	}
}

func (m *MongoDB) serverStatus(ctx context.Context, w io.Writer) {
	stats := m.runCommand(ctx, "serverStatus")
	for _, counter := range m.convertD(stats["opcounters"]) {
		op := counter.Key
		if cnt, yes := counter.Value.(int64); yes {
			name := "mongodb_opcounters_total{op=" + strconv.Quote(op) + "}"
			metrics.WriteCounterUint64(w, name, uint64(cnt))
		}
	}
	for _, counter := range m.convertD(stats["mem"]) {
		key := counter.Key
		if key != "resident" && key != "virtual" {
			continue
		}

		if cnt, yes := counter.Value.(int32); yes {
			name := "mongodb_mem_" + key + "_megabytes"
			metrics.WriteCounterUint64(w, name, uint64(cnt))
		}
	}
	for _, counter := range m.convertD(stats["connections"]) {
		key := counter.Key
		cnt, yes := counter.Value.(int32)
		if !yes {
			continue
		}

		if key == "current" || key == "active" {
			name := "mongodb_connections_" + key
			metrics.WriteGaugeUint64(w, name, uint64(cnt))
		} else if key == "totalCreated" {
			name := "mongodb_connections_created_total"
			metrics.WriteGaugeUint64(w, name, uint64(cnt))
		}
	}
}

func (m *MongoDB) runCommand(parent context.Context, command string) map[string]any {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()

	ret := make(bson.M)
	res := m.db.RunCommand(ctx, bson.D{{Key: command, Value: 1}})
	if err := res.Decode(&ret); err != nil {
		m.log.Warn("命令执行出错", "command", command, "err", err)
	}

	return ret
}

func (*MongoDB) convertD(v any) bson.D {
	if ds, ok := v.(bson.D); ok {
		return ds
	}

	return nil
}
