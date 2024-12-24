package model

import "time"

// AlertServer 告警服务器配置
type AlertServer struct {
	ID        int64     `json:"-"          gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Mode      string    `json:"mode"       gorm:"column:mode;type:varchar(10);not null;comment:发送模式"` // 枚举：siem dong
	Name      string    `json:"name"       gorm:"column:name;type:varchar(20);not null;index:idx_name,unique;comment:名字"`
	URL       string    `json:"url"        gorm:"column:url;type:varchar(255);not null;comment:服务器地址"`
	Token     string    `json:"token"      gorm:"column:token;type:varchar(255);not null;comment:认证令牌"`
	Account   string    `json:"account"    gorm:"column:account;type:varchar(20);comment:咚咚账号"`
	CreatedAt time.Time `json:"created_at" gorm:"column:updated_at;not null;default:now(3);comment:更新时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:created_at;not null;default:now(3);comment:创建时间"`
}

func (AlertServer) TableName() string {
	return "alert_server"
}
