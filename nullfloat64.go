package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strconv"
  "fmt"
)

type NullFloat64 struct {
	sql.NullFloat64
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

  if i == nil {
    return NullFloat64{
      sql.NullFloat64{
        Valid:   false,
      },
    }, nil
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
    return NullFloat64{
      sql.NullFloat64{
        Valid: false,
      },
    }, err
  }

  return NullFloat64{
    sql.NullFloat64{
      Valid:   true,
      Float64: val,
    },
  }, nil
}

func (nt NullFloat64) ToValue() interface{} {
  if !nt.Valid {
    return nil
  }

  return nt.Float64
}


func (nt NullFloat64) MarshalJSON() ([]byte, error) {

	if !nt.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nt.Float64)
}

func (nt *NullFloat64) UnmarshalJSON(b []byte) error {
	nt.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		// Try to unmarshal as float first, if it fails, try
		// to use a string and convert it
		if err := json.Unmarshal(b, &nt.Float64); err != nil {

			var str string
			if err := json.Unmarshal(b, &str); err != nil {
				return err
			}

			nt.Float64, err = strconv.ParseFloat(str, 64)
			if err != nil {
				return err
			}
		}
		nt.Valid = true
	}

	return nil
}
