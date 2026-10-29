package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OAuthConfig struct {
	ID                    bson.ObjectID `bson:"_id,omitempty"          json:"-"`
	Enabled               bool          `bson:"enabled"                json:"enabled"`
	ClientID              string        `bson:"client_id"              json:"client_id"`
	ClientSecret          string        `bson:"client_secret"          json:"client_secret"`
	AuthorizationEndpoint string        `bson:"authorization_endpoint" json:"authorization_endpoint"`
	TokenEndpoint         string        `bson:"token_endpoint"         json:"token_endpoint"`
	UserinfoEndpoint      string        `bson:"userinfo_endpoint"      json:"userinfo_endpoint"`
	RedirectURIs          []string      `bson:"redirect_uris"          json:"redirect_uris"`
	Scopes                []string      `bson:"scopes,omitempty"       json:"scopes,omitzero"`
	CreatedAt             time.Time     `bson:"created_at,omitempty"   json:"created_at"`
	UpdatedAt             time.Time     `bson:"updated_at,omitempty"   json:"updated_at"`
}

func (OAuthConfig) CollectionInfo() (string, string) { return "oauth_config", "OAuth 认证配置" }
