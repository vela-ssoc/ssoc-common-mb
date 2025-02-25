package errcode

import "github.com/xgfone/ship/v5"

var ErrFieldNameRequired = ship.ErrBadRequest.Newf("字段名必须填写")

const (
	FmtSchemaSQL = formatError("自定义搜索语句错误: %v")
)

type formatError string

func (f formatError) Fmt(a ...any) error {
	return ship.ErrBadRequest.Newf(string(f), a...)
}
