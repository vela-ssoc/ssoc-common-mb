package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ZipArchive struct {
	ID          bson.ObjectID `bson:"_id,omitempty"        json:"id"`          // 数据库 ID
	Name        string        `bson:"name"                 json:"name"`        // 文件名（不带 .zip 后缀）
	Description string        `bson:"description"          json:"description"` // 说明描述信息
	Method      int           `bson:"method"               json:"method"`      // 加密算法 0-不加密 1-ZipCrypto 2-AES-128 3-AES-192 4-AES-256
	Password    string        `bson:"password"             json:"password"`    // 加密密码（下载时生效）
	Version     int64         `bson:"version"              json:"version"`     // 版本计数器：每次子文件增删改时自增
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at"`  // 创建时间
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at"`  // 最近一次修改时间
}

func (ZipArchive) CollectionInfo() (string, string) { return "zip_archive", "ZIP 文件元数据" }
