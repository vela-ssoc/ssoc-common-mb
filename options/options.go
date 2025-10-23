package options

func Eval[T any](opts ...func(before T) (after T)) T {
	var v T
	for _, opt := range opts {
		if opt != nil {
			v = opt(v)
		}
	}

	return v
}
