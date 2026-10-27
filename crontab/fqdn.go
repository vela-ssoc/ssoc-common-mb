package crontab

import (
	"crypto/sha1"
	"encoding/hex"
	"reflect"
	"runtime"
)

// QualifiedID 通过 FQDN 生成任务唯一 ID。
//
// 只要对象的名称或路径未改变，则该方法生成的 ID 能保证唯一、稳定。
// 由于是通过 FQDN 生成的名字，所以严禁一个 struct 生成多个实例。
// 起名困难症的福音。
func QualifiedID(v any) string {
	s := fqdn(v)
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func fqdn(v any) string {
	if v == nil {
		return "<nil>"
	}

	tof := reflect.TypeOf(v)
	kind := tof.Kind()
	if kind == reflect.Pointer {
		tof = tof.Elem()
	} else if kind == reflect.Func {
		vof := reflect.ValueOf(v).Pointer()
		if pc := runtime.FuncForPC(vof); pc != nil {
			return pc.Name()
		}

		return ""
	}

	pkg, name := tof.PkgPath(), tof.Name()
	if pkg == "" {
		pkg = "main"
	}

	return pkg + "." + name
}
