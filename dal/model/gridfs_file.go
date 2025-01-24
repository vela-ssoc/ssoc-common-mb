package model

import "time"

type GridfsFile struct {
	ID        int64     `json:"id"         gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;size:255;comment:文件名"`
	Size      int64     `json:"size"       gorm:"column:size;comment:文件大小"`
	Checksum  string    `json:"checksum"   gorm:"column:name;size:50;comment:哈希校验码"`
	CreatedAt time.Time `json:"created_at" gorm:"column:updated_at;notnull;default:now(3);comment:创建时间"`
}

func (GridfsFile) TableName() string {
	return "gridfs_file"
}
