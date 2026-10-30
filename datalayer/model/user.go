package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Company   string        `bson:"company"              json:"company"`             // 集团-group ,证券-security ,天天基金-ttfund ,期货-futures , 东财基金-dcfund
	Username  string        `bson:"username"             json:"username"`            // 用户名
	Nickname  string        `bson:"nickname"             json:"nickname"`            // 昵称
	AvatarURL string        `bson:"avatar_url"           json:"avatar_url"`          // 用户头像
	Enabled   bool          `bson:"enabled"              json:"enabled"`             // 是否启用
	OAuth     bool          `bson:"oauth"                json:"oauth"`               // 是否 OAuth 认证用户
	Password  UserPassword  `bson:"password"             json:"password"`            // 登录密码（本地用户登录需要）
	Session   UserSession   `bson:"session"              json:"session"`             // Session
	TOTP      UserTOTP      `bson:"totp"                 json:"totp"`                // TOTP
	ExpiresAt time.Time     `bson:"expires_at"           json:"expires_at,omitzero"` // 过期时间（空表示永不过期）
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`          // 创建时间
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`          // 更新时间
}

func (User) CollectionInfo() (string, string) { return "user", "用户" }

type UserTOTP struct {
	Secret     string    `bson:"secret"      json:"-"`           // TOTP 密钥
	Verified   bool      `bson:"verified"    json:"verified"`    // 是否已验证（即扫码绑定成功）
	VerifiedAt time.Time `bson:"verified_at" json:"verified_at"` // 验证成功的时间
}

type UserSession struct {
	Token    string    `bson:"token,omitempty" json:"-"`         // Token
	IssuedAt time.Time `bson:"issued_at"       json:"issued_at"` // 签发时间
	LastedAt time.Time `bson:"lasted_at"       json:"lasted_at"` // 最近活动时间（非实时更新）
}

type UserPassword struct {
	Secret    string    `bson:"secret"     json:"-"`
	ChangedAt time.Time `bson:"changed_at" json:"changed_at"`
}

type Operator struct {
	ID       bson.ObjectID `bson:"id"                json:"id"`
	Company  string        `bson:"company"           json:"company"`
	Username string        `bson:"username"          json:"username"`
	Nickname string        `bson:"nickname"          json:"nickname"`
	Service  bool          `bson:"service,omitempty" json:"service,omitzero"`
}
