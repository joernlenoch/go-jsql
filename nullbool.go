package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type NullBool struct {
	sql.NullBool
}

func NewNullBool(i interface{}) NullBool {
	n, _ := TryNullBool(i)
	return n
}

// Create a new NullFloat.
// - nil and numeric values are considered correct
func TryNullBool(i interface{}) (NullBool, error) {
	nb := NullBool{}
	return nb, nb.TrySet(i)
}

func (nb *NullBool) Set(i interface{}) {
	nb.TrySet(i)
}

func (nb *NullBool) TrySet(i interface{}) error {

	if i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil()) {
		nb.Bool = false
		nb.Valid = false
		return nil
	}

	// If the given data is a NullArray object, copy the data directly
	if copy, ok := i.(*NullBool); ok {
		nb.Valid = copy.Valid
		nb.Bool = copy.Bool
		return nil
	} else if copy, ok := i.(NullBool); ok {
		nb.Valid = copy.Valid
		nb.Bool = copy.Bool
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
		nb.Bool = false
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

func (nb NullBool) Add(nt2 NullBool) NullBool {

	val1 := false
	if nb.Valid {
		val1 = nb.Bool
	}

	val2 := false
	if nt2.Valid {
		val2 = nt2.Bool
	}

	return NullBool{
		NullBool: sql.NullBool{
			Valid: nb.Valid || nt2.Valid,
			Bool:  val1 || val2,
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
	nb.Bool = false
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
