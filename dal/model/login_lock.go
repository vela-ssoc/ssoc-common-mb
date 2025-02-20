package model

import "time"

// LoginLock 登录锁定记录表
type LoginLock struct {
	ID        int64     `json:"id"         gorm:"column:id;primaryKey"`    // 表 ID
	Username  string    `json:"username"   gorm:"column:username;size:20"` // 用户名
	CreatedAt time.Time `json:"created_at" gorm:"column:updated_at;notnull;default:now(3);comment:创建时间"`
}

func (LoginLock) TableName() string {
	return "login_lock"
}
