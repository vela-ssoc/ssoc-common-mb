package dyncond

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type WhereInput struct {
	ShortID  string
	Operator Operator
	Values   []string
}

type WhereInputs []*WhereInput

func newTables(db *gorm.DB, wheres Wheres, orders Orders) *Tables {
	wheresMap := make(map[string]*Where, 16)
	ordersMap := make(map[string]*Order, 16)

	for _, where := range wheres {
		fieldName, uid := where.FieldName, where.UniqueID
		if dup := wheresMap[fieldName]; dup == nil {
			where.ShortID = fieldName
			wheresMap[fieldName] = where
		} else {
			delete(wheresMap, fieldName)
			dup.ShortID = dup.UniqueID
			wheresMap[dup.UniqueID] = dup
			wheresMap[uid] = where
		}
	}
	for _, order := range orders {
		fieldName, uid := order.FieldName, order.UniqueID
		if dup := ordersMap[fieldName]; dup == nil {
			order.ShortID = fieldName
			ordersMap[fieldName] = order
		} else {
			delete(ordersMap, fieldName)
			dup.ShortID = dup.UniqueID
			ordersMap[dup.UniqueID] = dup
			ordersMap[uid] = order
		}
	}

	return &Tables{
		wheres:    wheres,
		orders:    orders,
		wheresMap: wheresMap,
		ordersMap: ordersMap,
		db:        db,
	}
}

type Tables struct {
	wheres    Wheres
	orders    Orders
	wheresMap map[string]*Where
	ordersMap map[string]*Order
	db        *gorm.DB
}

func (tbl *Tables) Columns() (Wheres, Orders) {
	return tbl.wheres, tbl.orders
}

func (tbl *Tables) CompileWhere(inputs WhereInputs, or bool) ([]gen.Condition, []field.Expr, error) {
	size := len(inputs)
	if size == 0 {
		return nil, nil, nil
	}

	ands := make([]gen.Condition, 0, size)
	ors := make([]field.Expr, 0, size)
	exprs := make([]field.Expr, 0, size)
	for _, input := range inputs {
		cond, expr, err := tbl.compileWhere(input)
		if err != nil {
			return nil, nil, err
		}
		if cond == nil {
			continue
		}
		ands = append(ands, cond)
		ors = append(ors, cond)
		exprs = append(exprs, expr)
	}
	if or {
		expr := field.Or(exprs...)
		return []gen.Condition{expr}, exprs, nil
	}

	return ands, exprs, nil
}

func (tbl *Tables) EqualsExpr(a, b field.Expr) bool {
	stmt := &gorm.Statement{DB: tbl.db}
	a.Build(stmt)
	an := stmt.SQL.String()
	stmt.SQL.Reset()

	b.Build(stmt)
	bn := stmt.SQL.String()

	return an == bn
}

func (tbl *Tables) compileWhere(input *WhereInput) (field.Expr, field.Expr, error) {
	fd := tbl.wheresMap[input.ShortID]
	if fd == nil {
		return nil, nil, nil
	}
	expr := fd.Expr
	switch f := expr.(type) {
	case field.String:
		return tbl.fieldString(f, input), expr, nil
	case field.Bool:
		return tbl.fieldBool(f, input), expr, nil
	case field.Time:
		return tbl.fieldTime(f, input), expr, nil
	case field.Int:
		return tbl.fieldInt(f, input), expr, nil
	case field.Float64:
		return tbl.fieldFloat64(f, input), expr, nil
	case field.Field:
		return tbl.fieldField(f, input), expr, nil
	default:
		return nil, nil, nil
	}
}

func (tbl *Tables) fieldString(f field.String, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.stringN(0)
	if !ok || arg0 == "" {
		switch op {
		case Eq, Neq:
		default:
			return nil
		}
	}

	switch op {
	case Eq:
		if !ok {
			return f.IsNull()
		}
		return f.Eq(arg0)
	case Neq:
		if !ok {
			return f.IsNotNull()
		}
		return f.Neq(arg0)
	case Gt:
		return f.Gt(arg0)
	case Gte:
		return f.Gte(arg0)
	case Lt:
		return f.Lt(arg0)
	case Lte:
		return f.Lte(arg0)
	case Like, NotLike:
		if !strings.Contains(arg0, "%") {
			arg0 = "%" + arg0 + "%"
		}
		if op == Like {
			return f.Like(arg0)
		}
		return f.NotLike(arg0)
	case Regex:
		return f.Regexp(arg0)
	case NotRegex:
		return f.NotRegxp(arg0)
	case Between, NotBetween:
		arg1, exist := values.stringN(1)
		if !exist {
			return nil
		}
		if op == Between {
			return f.Between(arg0, arg1)
		}
		return f.NotBetween(arg0, arg1)
	case In:
		return f.In(values...)
	case NotIn:
		return f.NotIn(values...)
	default:
		return nil
	}
}

func (tbl *Tables) fieldInt(f field.Int, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.intN(0)
	if !ok {
		switch op {
		case Eq, Neq:
		default:
			return nil
		}
	}

	switch op {
	case Eq:
		if ok {
			return f.Eq(arg0)
		}
		return f.IsNull()
	case Neq:
		if ok {
			return f.IsNotNull()
		}
		return f.Neq(arg0)
	case Gt:
		return f.Gt(arg0)
	case Gte:
		return f.Gte(arg0)
	case Lt:
		return f.Lt(arg0)
	case Lte:
		return f.Lte(arg0)
	case Like:
		return f.Like(arg0)
	case NotLike:
		return f.NotLike(arg0)
	case Between, NotBetween:
		arg1, exist := values.intN(1)
		if !exist {
			return nil
		}
		if op == Between {
			return f.Between(arg0, arg1)
		}
		return f.NotBetween(arg0, arg1)
	case In, NotIn:
		args := values.ints()
		if len(args) == 0 {
			return nil
		}
		if op == In {
			return f.In(args...)
		}
		return f.NotIn(args...)
	default:
		return nil
	}
}

func (tbl *Tables) fieldFloat64(f field.Float64, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.float64N(0)
	if !ok && op != Eq && op != Neq {
		return nil
	}

	switch op {
	case Eq:
		if ok {
			return f.Eq(arg0)
		}
		return f.IsNull()
	case Neq:
		if ok {
			return f.Neq(arg0)
		}
		return f.IsNotNull()
	case Gt:
		return f.Gt(arg0)
	case Gte:
		return f.Gte(arg0)
	case Lt:
		return f.Lt(arg0)
	case Lte:
		return f.Lte(arg0)
	case Like:
		return f.Like(arg0)
	case NotLike:
		return f.NotLike(arg0)
	case Between, NotBetween:
		arg1, exist := values.float64N(1)
		if !exist {
			return nil
		}
		if op == Between {
			return f.Between(arg0, arg1)
		}
		return f.NotBetween(arg0, arg1)
	case In, NotIn:
		args := values.float64s()
		if len(args) == 0 {
			return nil
		}
		if op == In {
			return f.In(args...)
		}
		return f.NotIn(args...)
	default:
		return nil
	}
}

func (tbl *Tables) fieldTime(f field.Time, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.timeN(0)
	if !ok && op != Eq && op != Neq {
		return nil
	}

	switch op {
	case Eq:
		if ok {
			return f.Eq(arg0)
		}
		return f.IsNull()
	case Neq:
		if ok {
			return f.Neq(arg0)
		}
		return f.IsNotNull()
	case Gt:
		return f.Gt(arg0)
	case Gte:
		return f.Gte(arg0)
	case Lt:
		return f.Lt(arg0)
	case Lte:
		return f.Lte(arg0)
	case Between, NotBetween:
		arg1, exist := values.timeN(1)
		if !exist {
			return nil
		}
		if op == Between {
			return f.Between(arg0, arg1)
		}
		return f.NotBetween(arg0, arg1)
	case In, NotIn:
		args := values.times()
		if len(args) == 0 {
			return nil
		}
		if op == In {
			return f.In(args...)
		}
		return f.NotIn(args...)
	default:
		return nil
	}
}

func (tbl *Tables) fieldBool(f field.Bool, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.boolN(0)
	if !ok && op != Eq && op != Neq {
		return nil
	}
	switch op {
	case Eq:
		if ok {
			return f.Is(arg0)
		}
		return f.IsNull()
	case Neq:
		if ok {
			return f.Is(!arg0)
		}
		return f.IsNotNull()
	default:
		return nil
	}
}

func (tbl *Tables) fieldField(f field.Field, input *WhereInput) field.Expr {
	values, op := stringValues(input.Values), input.Operator
	arg0, ok := values.valueN(0)
	if !ok {
		return nil
	}
	switch op {
	case Eq:
		return f.Eq(arg0)
	case Neq:
		return f.Neq(arg0)
	case Gt:
		return f.Gt(arg0)
	case Gte:
		return f.Gte(arg0)
	case Lt:
		return f.Lt(arg0)
	case Lte:
		return f.Lte(arg0)
	case Like:
		return f.Like(arg0)
	case In, NotIn:
		vals := values.values()
		if len(vals) == 0 {
			return nil
		}
		return f.In(vals...)
	default:
		return nil
	}
}
