package model

import (
	"database/sql/driver"
	"encoding/json"
)

type RawMessage json.RawMessage

func (r RawMessage) Value() (driver.Value, error) {
	return r, nil
}
