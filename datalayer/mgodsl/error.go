package mgodsl

import "fmt"

type DSLError struct {
	Err error
}

func (d *DSLError) Error() string {
	return d.Err.Error()
}

func (d *DSLError) Unwrap() error {
	return d.Err
}

func newDSLErrorf(msg string, args ...any) *DSLError {
	return &DSLError{Err: fmt.Errorf(msg, args...)}
}

type FuncError struct {
	Receiver any
	Method   string
}

func (fe *FuncError) Error() string {
	//TODO implement me
	panic("implement me")
}
