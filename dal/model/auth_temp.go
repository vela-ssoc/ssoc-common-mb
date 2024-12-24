package model

import "time"

type AuthTemp struct {
	ID        int64     `json:"id"         gorm:"column:id;primaryKey"`
	UID       string    `json:"uid"        gorm:"column:uid;type:varchar(255);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;default:now(3);comment:创建时间"`
}

func (AuthTemp) TableName() string {
	return "auth_temp"
}

func (t AuthTemp) Expired(current time.Time, du time.Duration) bool {
	return t.CreatedAt.Add(du).Before(current)
}
