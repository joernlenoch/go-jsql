package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type NullBool struct {
	sql.NullBool
}

func NewNullBool(i interface{}) *NullBool {
	n, _ := TryNullBool(i)
	return n
}

// Create a new NullFloat.
// - nil and numeric values are considered correct
func TryNullBool(i interface{}) (*NullBool, error) {
	nb := &NullBool{}
	return nb, nb.TrySet(i)
}

func (nb *NullBool) TrySet(i interface{}) error {

	if i == nil {
		nb.Valid = false
		return nil
	}

	var val bool
	var err error

	switch i.(type) {
	case bool:
		val = i.(bool)
	default:
		val, err = strconv.ParseBool(fmt.Sprint(i))
	}

	if err != nil {
		nb.Valid = false
		return err
	}

	nb.Valid = true
	nb.Bool = val
	return nil
}

func (nb NullBool) ToValue() interface{} {
	if !nb.Valid {
		return nil
	}

	return nb.Bool
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
