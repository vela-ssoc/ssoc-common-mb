package model

import "time"

type GridFile struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;type:varchar(255);not null;index;comment:名字"`
	Extension string    `json:"extension"  gorm:"column:extension;type:varchar(50);index;comment:扩展名"`
	MIME      string    `json:"mime"       gorm:"column:mime;type:varchar(50);comment:MIME"`
	MD5       string    `json:"md5"        gorm:"column:md5;type:char(32);comment:MD5"`
	SHA1      string    `json:"sha1"       gorm:"column:sha1;type:char(40);comment:SHA1"`
	SHA256    string    `json:"sha256"     gorm:"column:sha256;type:char(64);comment:SHA256"`
	Size      int64     `json:"size"       gorm:"column:size;comment:文件大小"`
	Shard     int       `json:"-"          gorm:"column:shard;comment:分片大小"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;default:now(3);comment:更新时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:now(3);comment:创建时间"`
}

func (GridFile) TableName() string {
	return "grid_file"
}
