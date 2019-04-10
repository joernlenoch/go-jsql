package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) Add(nt2 NullFloat64) NullFloat64 {

	val1 := float64(0)
	if nf.Valid {
		val1 = nf.Float64
	}

	val2 := float64(0)
	if nt2.Valid {
		val2 = nt2.Float64
	}

	return NullFloat64{
		NullFloat64: sql.NullFloat64{
			Valid:   nf.Valid || nt2.Valid,
			Float64: val1 + val2,
		},
	}
}

func (nf NullFloat64) ToValue() interface{} {
	if !nf.Valid {
		return nil
	}

	return nf.Float64
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {

	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float64)
}

func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	nf.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		// Try to unmarshal as float first, if it fails, try
		// to use a string and convert it
		if err := json.Unmarshal(b, &nf.Float64); err != nil {

			var str string
			if err := json.Unmarshal(b, &str); err != nil {
				return err
			}

			nf.Float64, err = strconv.ParseFloat(str, 64)
			if err != nil {
				return err
			}
		}
		nf.Valid = true
	}

	return nil
}

// Create a new NullFloat64 value.
// Invalid values will result in having an invalid NullFloat64.
func NewNullFloat64(i interface{}) NullFloat64 {
	nf, _ := TryNullFloat64(i)
	return nf
}

// Create a new NullFloat.
// - nil and numeric values are considered correct
func TryNullFloat64(i interface{}) (NullFloat64, error) {
	nf := NullFloat64{}
	return nf, nf.TrySet(i)
}

func (nf *NullFloat64) Set(i interface{}) {
	nf.TrySet(i)
}

func (nf *NullFloat64) TrySet(i interface{}) error {

	if i == nil {
		nf.Valid = false
		return nil
	}

	// If the given data is a NullArray object, copy the data directly
	if copy, ok := i.(*NullFloat64); ok {
		nf.Valid = copy.Valid
		nf.Float64 = copy.Float64
		return nil
	} else if copy, ok := i.(NullFloat64); ok {
		nf.Valid = copy.Valid
		nf.Float64 = copy.Float64
		return nil
	}

	var val float64
	var err error

	switch i.(type) {
	case int:
		val = float64(i.(int))
	case int8:
		val = float64(i.(int8))
	case int16:
		val = float64(i.(int16))
	case int32:
		val = float64(i.(int32))
	case int64:
		val = float64(i.(int64))
	case uint:
		val = float64(i.(uint))
	case uint8:
		val = float64(i.(uint8))
	case uint16:
		val = float64(i.(uint16))
	case uint32:
		val = float64(i.(uint32))
	case uint64:
		val = float64(i.(uint64))
	case float32:
		val = float64(i.(float32))
	case float64:
		val = i.(float64)
	default:
		val, err = strconv.ParseFloat(fmt.Sprint(i), 64)
	}

	if err != nil {
		nf.Valid = false
		return err
	}

	nf.Valid = true
	nf.Float64 = val

	return nil
}
