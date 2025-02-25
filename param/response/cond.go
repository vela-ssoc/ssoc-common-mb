package response

import "github.com/vela-ssoc/vela-common-mb/dal/dyncond"

func ParseCond(tbl *dyncond.Tables) *Cond {
	wheres, orders := tbl.Columns()
	cond := new(Cond)
	for _, wc := range wheres {
		where := &CondWhere{
			Key:       wc.ShortID,
			Desc:      wc.Comment,
			Type:      wc.DataType,
			Operators: parseCondOperators(wc.Operators),
			Enums:     parseCondEnums(wc.Enums),
		}
		cond.Wheres = append(cond.Wheres, where)
	}
	cond.Conditions = cond.Wheres

	for _, oc := range orders {
		order := &CondKeyDesc{Key: oc.ShortID, Desc: oc.Comment}
		cond.Orders = append(cond.Orders, order)
	}

	return cond
}

func parseCondEnums(enums dyncond.Enums) []*CondKeyDesc {
	size := len(enums)
	if size == 0 {
		return nil
	}

	ret := make([]*CondKeyDesc, 0, size)
	for _, enum := range enums {
		ret = append(ret, &CondKeyDesc{Key: enum.Key, Desc: enum.Desc})
	}

	return ret
}

func parseCondOperators(ops dyncond.Operators) []*CondKeyDesc {
	size := len(ops)
	if size == 0 {
		return nil
	}

	ret := make([]*CondKeyDesc, 0, size)
	for _, op := range ops {
		key, desc := op.Info()
		ret = append(ret, &CondKeyDesc{Key: key, Desc: desc})
	}

	return ret
}

type Cond struct {
	Conditions []*CondWhere   `json:"conditions,omitempty"`
	Wheres     []*CondWhere   `json:"wheres,omitempty"`
	Orders     []*CondKeyDesc `json:"orders,omitempty"`
}

type CondWhere struct {
	Key       string         `json:"key"`
	Desc      string         `json:"desc"`
	Type      string         `json:"type"`
	Operators []*CondKeyDesc `json:"operators,omitzero"`
	Enums     []*CondKeyDesc `json:"enums,omitzero"`
}

type CondKeyDesc struct {
	Key  string `json:"key"`
	Desc string `json:"desc"`
}
