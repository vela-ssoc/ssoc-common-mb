package dyncond

import "gorm.io/gen/field"

type Orders []*Order

type Wheres []*Where

type Enum struct {
	Key  string
	Desc string
}

type Enums []*Enum

type Where struct {
	UniqueID  string     // ID
	ShortID   string     // 短 ID
	TableName string     // 表名
	FieldName string     // 字段名
	FullName  string     // 表名+字段名
	DataType  string     // 数据类型
	Comment   string     // 注释名字
	Operators Operators  // 操作符
	Enums     Enums      // 枚举值
	Expr      field.Expr // gorm 表达式
}

type Order struct {
	UniqueID  string          // ID
	ShortID   string          // 短 ID
	TableName string          // 表名
	FieldName string          // 字段名
	FullName  string          // 表名+字段名
	DataType  string          // 数据类型
	Comment   string          // 注释名字
	Expr      field.OrderExpr // gorm 表达式
}
