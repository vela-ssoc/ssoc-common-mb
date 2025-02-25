package dyncond

import (
	"reflect"

	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type fieldInfo struct {
	expr      field.Expr
	name      string
	tablename string
	fullname  string
	kind      string
	desc      string
	operators Operators
}

func parseFieldInfo(tableName string, f *schema.Field, stmt *gorm.Statement) fieldInfo {
	dbName, dataType := f.DBName, f.DataType
	info := fieldInfo{
		name:      dbName,
		tablename: tableName,
		kind:      string(dataType),
		desc:      f.Comment,
	}

	realType := getFieldRealType(f.FieldType)
	switch realType {
	case "string":
		info.expr = field.NewString(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Like, NotLike, Regex, NotRegex, Between, NotBetween, In, NotIn}
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		info.expr = field.NewInt(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Between, NotBetween, In, NotIn}
	case "float32", "float64":
		info.expr = field.NewFloat64(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Between, NotBetween, In, NotIn}
	case "bool":
		info.expr = field.NewBool(tableName, dbName)
		info.operators = Operators{Eq, Neq}
	case "time.Time":
		info.expr = field.NewTime(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Between, NotBetween, In, NotIn}
	case "bytes":
		info.expr = field.NewBytes(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Like, In}
	case "serializer":
		info.expr = field.NewSerializer(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Like, In}
	default:
		info.expr = field.NewField(tableName, dbName)
		info.operators = Operators{Eq, Neq, Gt, Gte, Lt, Lte, Like, In}
	}

	stmt.SQL.Reset()
	info.expr.Build(stmt)
	info.fullname = stmt.SQL.String()

	return info
}

// getFieldRealType  get basic type of field.
// https://github.com/go-gorm/gen/blob/v0.3.26/internal/generate/query.go#L75-L96
func getFieldRealType(f reflect.Type) string {
	serializerInterface := reflect.TypeOf((*schema.SerializerInterface)(nil)).Elem()
	if f.Implements(serializerInterface) || reflect.New(f).Type().Implements(serializerInterface) {
		return "serializer"
	}
	scanValuer := reflect.TypeOf((*field.ScanValuer)(nil)).Elem()
	if f.Implements(scanValuer) || reflect.New(f).Type().Implements(scanValuer) {
		return "field"
	}

	if f.Kind() == reflect.Ptr {
		f = f.Elem()
	}
	if f.String() == "time.Time" {
		return "time.Time"
	}
	if f.String() == "[]uint8" || f.String() == "json.RawMessage" {
		return "bytes"
	}

	return f.Kind().String()
}
