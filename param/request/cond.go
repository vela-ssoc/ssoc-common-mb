package request

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/vela-ssoc/ssoc-common-mb/dal/dyncond"
	"github.com/vela-ssoc/ssoc-common-mb/param/errcode"
)

type KeywordConditions struct {
	Keywords
	Conditions
}

type PageKeywordConditions struct {
	Pages
	Keywords
	Conditions
}

type PageConditions struct {
	Pages
	Conditions
}

type Conditions struct {
	CondWhereInputs
	CondOrderInputs
}

type CondOrderInputs struct {
	Order []*CondOrderInput `json:"order" form:"order" query:"order" validate:"lte=10,dive"`
}

type CondOrderInput struct {
	Key  string `json:"key"`
	Desc bool   `json:"desc"`
}

// UnmarshalBind 从 query 参数中解析排序规则。
//
// 1. 支持常规的 JSON 格式：
//   - {"key":"age"}                 ORDER BY age
//   - {"key":"age","desc":false}    ORDER BY age
//   - {"key":"age","desc":true}     ORDER BY age DESC
//
// 2. 支持 <key>:<desc> 格式（因为 JSON 格式放在 query 中太冗长）：
//
//	<desc> 是未定义的值，或者是空（此时冒号可以省略），均按照升序排序：
//		- age:hello    ORDER BY age
//		- age:balabala ORDER BY age
//		- age:         ORDER BY age
//		- age	       ORDER BY age
//	<desc> 可以是枚举：asc、desc：
//		- age:asc      ORDER BY age
//		- age:desc     ORDER BY age DESC
//	<desc> 可以是数字，小于 0 时代表降序，否则代表升序：
//		- age:0        ORDER BY age
//		- age:123      ORDER BY age
//		- age:-1       ORDER BY age DESC
//		- age:-456     ORDER BY age DESC
//
// 错误格式：
//   - :desc
//   - :巴拉巴拉
func (o *CondOrderInput) UnmarshalBind(str string) error {
	if data := []byte(str); json.Valid(data) {
		return json.Unmarshal(data, o)
	}

	idx := strings.LastIndex(str, ":")
	if idx < 0 {
		o.Key = str
		return nil
	} else if idx == 0 {
		return errcode.ErrFieldNameRequired
	}

	o.Key = str[:idx]
	value := str[idx+1:]
	if sorted := strings.ToLower(str[idx+1:]); sorted == "asc" || sorted == "desc" {
		o.Desc = sorted == "desc"
		return nil
	}

	num, _ := strconv.ParseInt(value, 10, 64)
	o.Desc = num < 0

	return nil
}

type CondWhereInput struct {
	Key      string `json:"key"      validate:"required,lte=100"`
	Operator string `json:"operator" validate:"oneof=eq neq gt gte lt lte like notlike regex notregex between notbetween in notin null"`
	Value    string `json:"value"    validate:"gt=0,lte=1000"`
}

func (c *CondWhereInput) UnmarshalBind(str string) error {
	if data := []byte(str); json.Valid(data) {
		return json.Unmarshal(data, c)
	}

	// created_at:gte:2024-10-14T06:31:20.620Z
	// size:between:1024,4096
	// hobbies:in:sing,dance,rap

	sn := strings.SplitN(str, ":", 3)
	if len(sn) != 3 {
		return errcode.ErrFieldNameRequired
	}
	name, opera, val := sn[0], sn[1], sn[2]
	c.Key, c.Operator = name, opera
	c.Value = val

	return nil
}

type CondWhereInputs struct {
	Or      bool              `json:"or"      form:"or"      query:"or"`
	Filters []*CondWhereInput `json:"filters" form:"filters" query:"filters" validate:"lte=200,dive"`
}

func (cws CondWhereInputs) Inputs() dyncond.WhereInputs {
	size := len(cws.Filters)
	inputs := make(dyncond.WhereInputs, 0, size)
	for _, cw := range cws.Filters {
		shortID := cw.Key
		if shortID == "" {
			continue
		}
		op, exist := dyncond.LookupOperator(cw.Operator)
		if !exist {
			continue
		}

		input := &dyncond.WhereInput{
			ShortID:  shortID,
			Operator: op,
			Values:   strings.Split(cw.Value, ","),
		}
		inputs = append(inputs, input)
	}

	return inputs
}
