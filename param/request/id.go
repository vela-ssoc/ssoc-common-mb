package request

import (
	"encoding/json"
	"strconv"

	"gorm.io/gen/field"
)

type Int64ID struct {
	ID int64 `json:"id,string" query:"id" form:"id" validate:"required,gt=0"`
}

type Int64IDOptional struct {
	ID int64 `json:"id,string" query:"id" form:"id"`
}

type Int64s []int64

func (is Int64s) MarshalJSON() (text []byte, err error) {
	size := len(is)
	str := make([]string, 0, size)
	for _, i := range is {
		s := strconv.FormatInt(i, 10)
		str = append(str, s)
	}

	return json.Marshal(str)
}

func (is *Int64s) UnmarshalJSON(raw []byte) error {
	var str []string
	if err := json.Unmarshal(raw, &str); err != nil {
		return err
	}

	dats := make([]int64, 0, len(str))
	for _, st := range str {
		num, err := strconv.ParseInt(st, 10, 64)
		if err != nil {
			return err
		}
		dats = append(dats, num)
	}
	*is = dats

	return nil
}

type IDName struct {
	ID   int64  `json:"id,string" gorm:"column:id"`
	Name string `json:"name"      gorm:"column:name"`
}

type IDNames []*IDName

func (dn IDNames) Map() map[int64]*IDName {
	hm := make(map[int64]*IDName, len(dn))
	for _, n := range dn {
		hm[n.ID] = n
	}
	return hm
}

type NameCount struct {
	Name  string `json:"name"  gorm:"column:name"`
	Count int    `json:"count" gorm:"column:count"`
}

type NameCounts []*NameCount

func (nc NameCounts) Aliases() (name field.String, count field.Int) {
	// 要与 NameCount gorm tag 保持一致
	name = field.NewString("", "name")
	count = field.NewInt("", "count")
	return name, count
}
