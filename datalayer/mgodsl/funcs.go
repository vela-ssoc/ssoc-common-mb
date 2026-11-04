package mgodsl

type FuncMap map[string]any

func (fm FuncMap) Call(method string, args []any) (any, error) {
	fn := fm[method]
	if fn == nil {

	}

	return nil, nil
}
