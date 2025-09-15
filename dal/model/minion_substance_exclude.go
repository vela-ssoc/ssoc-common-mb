package model

type MinionSubstanceExclude struct {
	ID          int64 `json:"id,string"           gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID    int64 `json:"minion_id,string"    gorm:"column:minion_id;uniqueIndex:uk_minion_id_substance_id;"`
	SubstanceID int64 `json:"substance_id,string" gorm:"column:substance_id;uniqueIndex:uk_minion_id_substance_id;"`
}

func (MinionSubstanceExclude) TableName() string {
	return "minion_substance_exclude"
}
