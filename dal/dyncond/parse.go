package dyncond

import (
	"context"
	"strings"

	"github.com/vela-ssoc/ssoc-common-mb/dal/query"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Options struct {
	WhereCallback func(*Tables, *Where) *Where
	OrderCallback func(*Tables, *Order) *Order
}

func ParseModels(qry *query.Query, mods []any, opts Options) (*Tables, error) {
	if opts.WhereCallback == nil {
		opts.WhereCallback = func(_ *Tables, v *Where) *Where { return v }
	}
	if opts.OrderCallback == nil {
		opts.OrderCallback = func(_ *Tables, v *Order) *Order { return v }
	}

	queryCtx := qry.WithContext(context.Background())
	db := queryCtx.User.UnderlyingDB()
	tbl := &Tables{db: db}

	var wheres Wheres
	var orders Orders
	for _, mod := range mods {
		stmt := &gorm.Statement{DB: db}
		if err := stmt.Parse(mod); err != nil {
			return nil, err
		}
		sch := stmt.Schema
		tableName, fields := sch.Table, sch.Fields

		for _, f := range fields {
			info := parseFieldInfo(tableName, f, stmt)
			uid := strings.Join([]string{tableName, info.name}, ".")
			where := &Where{
				UniqueID:  uid,
				ShortID:   uid,
				TableName: tableName,
				FieldName: info.name,
				FullName:  info.fullname,
				DataType:  info.kind,
				Comment:   info.desc,
				Operators: info.operators,
				Expr:      info.expr,
			}
			if where.DataType == "bool" {
				where.Enums = Enums{{Key: "true", Desc: "是"}, {Key: "false", Desc: "否"}}
			}
			if cond := opts.WhereCallback(tbl, where); cond != nil {
				wheres = append(wheres, cond)
			}

			if expr, ok := info.expr.(field.OrderExpr); ok {
				order := &Order{
					UniqueID:  uid,
					ShortID:   uid,
					TableName: tableName,
					FieldName: info.name,
					FullName:  info.fullname,
					DataType:  info.kind,
					Comment:   info.desc,
					Expr:      expr,
				}
				if cond := opts.OrderCallback(tbl, order); cond != nil {
					orders = append(orders, cond)
				}
			}
		}
	}

	return newTables(db, wheres, orders), nil
}
