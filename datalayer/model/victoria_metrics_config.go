package model

import (
	"encoding/base64"
	"net/url"
	"path"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VictoriaMetricsConfig struct {
	ID        bson.ObjectID          `bson:"_id,omitempty"        json:"id"`
	Enabled   bool                   `bson:"enabled"              json:"enabled"`
	Address   string                 `bson:"address"              json:"address"`
	Tenant    *VictoriaMetricsTenant `bson:"tenant"               json:"tenant,omitempty"`
	Username  string                 `bson:"username"             json:"username"`
	Password  string                 `bson:"password"             json:"password"`
	Headers   HTTPHeader             `bson:"headers,omitempty"    json:"headers"`
	CreatedAt time.Time              `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time              `bson:"updated_at,omitempty" json:"updated_at"`
}

func (VictoriaMetricsConfig) CollectionInfo() (string, string) {
	return "victoria_metrics_config", "VictoriaMetrics 数据库配置"
}

type VictoriaMetricsTenant struct {
	AccountID uint32  `bson:"account_id" json:"account_id"`
	ProjectID *uint32 `bson:"project_id" json:"project_id"`
}

func (v VictoriaMetricsTenant) TenantID() string {
	id := strconv.FormatUint(uint64(v.AccountID), 10)
	if pid := v.ProjectID; pid != nil {
		id = id + ":" + strconv.FormatUint(uint64(*pid), 10)
	}

	return id
}

func (c VictoriaMetricsConfig) BasicAuth() string {
	if c.Username == "" && c.Password == "" {
		return ""
	}
	auth := c.Username + ":" + c.Password

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c VictoriaMetricsConfig) Target() (*url.URL, HTTPHeader, error) {
	addr, err := url.Parse(c.Address)
	if err != nil {
		return nil, nil, err
	}

	header := c.Headers.Clone()
	if basic := c.BasicAuth(); basic != "" {
		header["Authorization"] = basic
	}

	return addr, header, nil
}

func (c VictoriaMetricsConfig) SelectPrefixPath() string {
	return c.prefixPath("select")
}

func (c VictoriaMetricsConfig) InsertPrefixPath() string {
	return c.prefixPath("insert")
}

func (c VictoriaMetricsConfig) prefixPath(typ string) string {
	if c.Tenant == nil { // 非租户模式没有前缀
		return ""
	}
	tid := c.Tenant.TenantID()

	// /insert/1/prometheus OR /select/1/prometheus
	return path.Join(typ, tid, "prometheus")
}
