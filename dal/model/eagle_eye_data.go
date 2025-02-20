package model

import (
	"encoding/json"
	"time"
)

type EagleEyeData struct {
	ID        int64           `json:"id,string"        gorm:"column:id;primaryKey"`
	MinionID  int64           `json:"minion_id,string" gorm:"column:minion_id"`
	Inet      string          `json:"inet"             gorm:"column:inet;size:20"`
	Path      string          `json:"path"             gorm:"column:path;size:255"`
	Data      json.RawMessage `json:"data"             gorm:"column:data"`
	CreatedAt time.Time       `json:"created_at"       gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at"       gorm:"column:updated_at"`
}

func (EagleEyeData) TableName() string {
	return "eagle_eye_data"
}
