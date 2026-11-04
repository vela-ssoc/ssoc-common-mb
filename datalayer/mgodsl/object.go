package mgodsl

import "time"

type object interface {
	invoke(method string, args []any) (object, error)
	unwrap() any
}

type timeObject struct {
	value time.Time
	funcs functionMap
}

func newTimeObject(v time.Time) *timeObject {
	obj := &timeObject{value: v}
	obj.funcs = functionMap{
		"add": obj.add,
	}

	return obj
}

func (t *timeObject) invoke(method string, args []any) (object, error) {
	return t.funcs.call(method, args)
}

func (t *timeObject) unwrap() any {
	return t.value
}

func (t *timeObject) add(str string) (*timeObject, error) {
	du, err := time.ParseDuration(str)
	if err != nil {
		return nil, err
	}

	return newTimeObject(t.value.Add(du)), nil
}

type ObjectCaller interface {
	Call(method string, args []any) (any, error)
	Value() error
}

type datetime struct {
	val time.Time
	fms FuncMap
}

func (d datetime) Call(method string, args []any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (d datetime) Value() error {
	//TODO implement me
	panic("implement me")
}

func (d datetime) name() {

}
