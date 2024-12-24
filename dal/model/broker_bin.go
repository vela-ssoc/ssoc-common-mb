package model

import "time"

// BrokerBin 存放 broker 节点的可执行程序。
type BrokerBin struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;type:varchar(255);not null;comment:名字"`
	FileID    int64     `json:"-"          gorm:"column:file_id;not null;comment:文件ID"`
	Size      int64     `json:"size"       gorm:"column:size;comment:文件大小"`
	Hash      string    `json:"hash"       gorm:"column:hash;type:varchar(128);comment:文件哈希"`
	Goos      string    `json:"goos"       gorm:"column:goos;type:varchar(10);not null;comment:操作系统"`
	Arch      string    `json:"arch"       gorm:"column:arch;type:varchar(15);not null;comment:系统架构"`
	Semver    Semver    `json:"semver"     gorm:"column:semver;type:varchar(100);not null;comment:版本"`
	Changelog string    `json:"changelog"  gorm:"column:changelog;type:text;comment:更新日志"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;default:now(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;default:now(3);comment:更新时间"`
}

func (BrokerBin) TableName() string {
	return "broker_bin"
}
