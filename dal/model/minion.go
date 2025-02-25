package model

import (
	"database/sql"
	"time"
)

type MinionStatus uint8

const (
	// MSInactive 未激活
	MSInactive MinionStatus = iota + 1
	// MSOffline 离线
	MSOffline
	// MSOnline 在线
	MSOnline
	// MSDelete 已删除
	MSDelete
)

func (ms MinionStatus) String() string {
	switch ms {
	case MSInactive:
		return "未激活"
	case MSOffline:
		return "离线"
	case MSOnline:
		return "在线"
	case MSDelete:
		return "已删除"
	default:
		return "未知"
	}
}

// Minion 节点表
type Minion struct {
	ID         int64        `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Inet       string       `json:"inet"             gorm:"column:inet;size:20;not null;unique;comment:IPv4"`
	Inet6      string       `json:"inet6"            gorm:"column:inet6;size:64;comment:IPv6"`
	MAC        string       `json:"mac"              gorm:"column:mac;size:17;comment:MAC地址"`
	Goos       string       `json:"goos"             gorm:"column:goos;size:10;comment:操作系统"`
	Arch       string       `json:"arch"             gorm:"column:arch;size:10;comment:系统架构"`
	Edition    string       `json:"edition"          gorm:"column:edition;size:50;comment:Agent版本"`
	Status     MinionStatus `json:"status"           gorm:"column:status;comment:节点状态"`
	Uptime     sql.NullTime `json:"uptime"           gorm:"column:uptime;comment:上线时间"`
	BrokerID   int64        `json:"broker_id,string" gorm:"column:broker_id;comment:代理节点ID"`
	BrokerName string       `json:"broker_name"      gorm:"column:broker_name;size:255;comment:代理节点"`
	Unload     bool         `json:"unload"           gorm:"column:unload;comment:是否静默模式"`
	Unstable   bool         `json:"unstable"         gorm:"column:unstable;comment:是否内测版本"`
	Customized string       `json:"customized"       gorm:"column:customized;size:50;comment:定制版本"`
	OrgPath    string       `json:"org_path"         gorm:"column:org_path;size:255;comment:部门路径"`
	Identity   string       `json:"identity"         gorm:"column:identity;type:text;comment:堡垒机用户"`
	Category   string       `json:"category"         gorm:"column:category;size:255;comment:部门信息"`
	OpDuty     string       `json:"op_duty"          gorm:"column:op_duty;size:255;comment:运维负责人"`
	Comment    string       `json:"comment"          gorm:"column:comment;type:text;comment:节点描述"`
	IBu        string       `json:"ibu"              gorm:"column:ibu;size:100;comment:部门"`
	IDC        string       `json:"idc"              gorm:"column:idc;size:50;comment:IDC"`
	CreatedAt  time.Time    `json:"created_at"       gorm:"column:created_at;notnull;default:now(3);comment:更新时间"`
	UpdatedAt  time.Time    `json:"updated_at"       gorm:"column:updated_at;notnull;default:now(3);comment:创建时间"`
}

// TableName implement gorm schema.Tabler
func (Minion) TableName() string {
	return "minion"
}

func (m Minion) Invalid() bool {
	return m.Status != MSOffline && m.Status != MSOnline
}

// Minions []*Minion
type Minions []*Minion

// BrokerMap 整理为 key: brokerID; value: minionIDs
func (ms Minions) BrokerMap() map[int64][]int64 {
	ret := make(map[int64][]int64, 16)
	for _, m := range ms {
		minionIDs := ret[m.BrokerID]
		if minionIDs == nil {
			ss := make([]int64, 0, 32)
			ss = append(ss, m.ID)
			ret[m.BrokerID] = ss
		} else {
			ret[m.BrokerID] = append(minionIDs, m.ID)
		}
	}
	return ret
}
