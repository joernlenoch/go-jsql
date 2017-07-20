package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type NullBool struct {
	sql.NullBool
}

func NewNullBool(b interface{}) NullBool {

	if val, ok := b.(bool); ok {
		return NullBool{
			sql.NullBool{
				Bool:  val,
				Valid: true,
			},
		}
	}

	return NullBool{
		sql.NullBool{
			Valid: false,
		},
	}
}

func (nb NullBool) MarshalJSON() ([]byte, error) {

	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

func (nb *NullBool) UnmarshalJSON(b []byte) error {
	nb.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		if err := json.Unmarshal(b, &nb.Bool); err != nil {
			return err
		}
		nb.Valid = true
	}

	return nil
}
