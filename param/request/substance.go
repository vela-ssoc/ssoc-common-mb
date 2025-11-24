package request

type SubstanceExclude struct {
	MinionID    int64 `json:"minion_id,string"    validate:"required,gt=0"`
	SubstanceID int64 `json:"substance_id,string" validate:"required,gt=0"`
}
