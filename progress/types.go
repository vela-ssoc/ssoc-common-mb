package progress

type Update struct {
	Current *int64
	Total   *int64
	Percent *float64
	Message *string
	Data    map[string]any
}
