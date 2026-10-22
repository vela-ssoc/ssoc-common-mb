package pterm

import "encoding/json"

type responseData struct {
	Type string `json:"type"`
	Data string `json:"data,omitzero"`
}

type ResizeData struct {
	Cols int `json:"cols" validate:"gt=0,lte=10000"`
	Rows int `json:"rows" validate:"gt=0,lte=10000"`
}

type Input struct {
	Type string          `json:"type" validate:"required"`
	Data json.RawMessage `json:"data"`
}

func (in Input) Unmarshal(v any) error {
	return json.Unmarshal(in.Data, v)
}

func (in Input) AsString() (string, error) {
	var str string
	err := in.Unmarshal(&str)
	return str, err
}

func (in Input) AsResize() (*ResizeData, error) {
	rd := new(ResizeData)
	if err := in.Unmarshal(rd); err != nil {
		return nil, err
	}

	return rd, nil
}
