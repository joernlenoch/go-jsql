package jsql

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type NullArray struct {
	Valid bool
	Array []interface{}
}

func NewNullArray(i interface{}) NullArray {
	na, _ := TryNewNullArray(i)
	return na
}

func TryNewNullArray(i interface{}) (NullArray, error) {
	na := NullArray{}
	return na, na.TrySet(i)
}

func (na *NullArray) Set(i interface{}) {
	na.TrySet(i)
}

func (na *NullArray) TrySet(i interface{}) error {

	if i == nil {
		na.Array = nil
		na.Valid = false
		return nil
	}

	// If the given data is a NullArray object, copy the data directly
	if copy, ok := i.(*NullArray); ok {
		na.Valid = copy.Valid
		na.Array = copy.Array
		return nil
	} else if copy, ok := i.(NullArray); ok {
		na.Valid = copy.Valid
		na.Array = copy.Array
		return nil
	}

	raw := reflect.ValueOf(i)

	if raw.Kind() != reflect.Slice {
		na.Array = nil
		na.Valid = false
		return fmt.Errorf("expected a slice, got a %v", raw.Kind())
	}

	a := make([]interface{}, raw.Len())
	for i := 0; i < raw.Len(); i++ {
		a[i] = raw.Index(i).Interface()
	}

	na.Array = a
	na.Valid = true
	return nil
}

func (na NullArray) ToStringArray() []string {
	out := make([]string, len(na.Array))

	for i, val := range na.Array {
		out[i] = fmt.Sprint(val)
	}

	return out
}

func (na NullArray) ToInt64Array() ([]int64, error) {
	out := make([]int64, len(na.Array))

	var err error
	for i, val := range na.Array {
		if i, ok := val.(int64); ok {
			out[i] = i
			continue
		}

		str := fmt.Sprint(val)
		out[i], err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (na NullArray) ToFloat64Array() ([]float64, error) {
	out := make([]float64, len(na.Array))

	var err error
	for i, val := range na.Array {
		if f, ok := val.(float64); ok {
			out[i] = f
			continue
		}
		str := fmt.Sprint(val)
		out[i], err = strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToValue transform the current value into nil or array
func (na NullArray) ToValue() interface{} {
	if !na.Valid {
		return nil
	}

	return na.Array
}

func (na NullArray) MarshalJSON() ([]byte, error) {
	if !na.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(na.Array)
}

func (na *NullArray) UnmarshalJSON(b []byte) error {
	na.Array = nil
	na.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		if err := json.Unmarshal(b, &na.Array); err != nil {
			return err
		}
		na.Valid = true
	}

	return nil
}

// Retrieve value as JSON data
func (na *NullArray) Scan(value interface{}) error {
	na.Array = nil
	na.Valid = false

	if value == nil {
		return nil
	}

	data, ok := value.([]byte)
	if !ok {
		return errors.New("The given data is not a valid string")
	}

	if len(data) == 0 {
		return nil
	}

	return na.UnmarshalJSON(data)
}

func (na NullArray) Value() (driver.Value, error) {

	if !na.Valid {
		return nil, nil
	}

	data, err := na.MarshalJSON()
	log.Print("Get value ", na, string(data), err)
	return na.MarshalJSON()
}
