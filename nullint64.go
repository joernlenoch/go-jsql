package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) ToValue() interface{} {
	if !ni.Valid {
		return nil
	}

	return ni.Int64
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {

	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	ni.Int64 = 0
	ni.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		if err := json.Unmarshal(b, &ni.Int64); err != nil {
			return err
		}
		ni.Valid = true
	}

	return nil
}

func (ni NullInt64) Add(nt2 NullInt64) NullInt64 {

	val1 := int64(0)
	if ni.Valid {
		val1 = ni.Int64
	}

	val2 := int64(0)
	if nt2.Valid {
		val2 = nt2.Int64
	}

	return NullInt64{
		NullInt64: sql.NullInt64{
			Valid: ni.Valid || nt2.Valid,
			Int64: val1 + val2,
		},
	}
}

func NewNullInt64(s interface{}) NullInt64 {
	n, _ := TryNullInt64(s)
	return n
}

// TryNullInt64 tries to create a new object
func TryNullInt64(i interface{}) (NullInt64, error) {
	ni := NullInt64{}
	return ni, ni.TrySet(i)
}

func (ni *NullInt64) Set(i interface{}) {
	ni.TrySet(i)
}

func (ni *NullInt64) TrySet(i interface{}) error {

	if IsNil(i) {
		ni.Int64 = 0
		ni.Valid = false
		return nil
	}

	// If the given data is a NullArray object, copy the data directly
	if copy, ok := i.(*NullInt64); ok {
		ni.Valid = copy.Valid
		ni.Int64 = copy.Int64
		return nil
	} else if copy, ok := i.(NullInt64); ok {
		ni.Valid = copy.Valid
		ni.Int64 = copy.Int64
		return nil
	}

	var val int64
	var err error

	switch i.(type) {
	case int:
		val = int64(i.(int))
	case int8:
		val = int64(i.(int8))
	case int16:
		val = int64(i.(int16))
	case int32:
		val = int64(i.(int32))
	case int64:
		val = i.(int64)
	case uint:
		val = int64(i.(uint))
	case uint8:
		val = int64(i.(uint8))
	case uint16:
		val = int64(i.(uint16))
	case uint32:
		val = int64(i.(uint32))
	case uint64:
		val = int64(i.(uint64))
	case float32:
		val = int64(i.(float32))
	case float64:
		val = int64(i.(float64))
	default:
		val, err = strconv.ParseInt(fmt.Sprint(i), 10, 64)
	}

	if err != nil {
		ni.Int64 = 0
		ni.Valid = false
		return err
	}

	ni.Int64 = val
	ni.Valid = true
	return nil
}
