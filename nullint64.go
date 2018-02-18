package jsql

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/go-errors"
)

type NullInt64 struct {
	sql.NullInt64
}

func NewNullInt64(s interface{}) NullInt64 {
	n, _ := TryNullInt64(s)
	return n
}

// Create a new NullFloat.
// - nil and numeric values are considered correct
func TryNullInt64(i interface{}) (NullInt64, error) {

  if i == nil {
    return NullInt64{
      sql.NullInt64{
        Valid:   false,
      },
    }, nil
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
    return NullInt64{
      sql.NullInt64{
        Valid: false,
      },
    }, err
  }

  return NullInt64{
    sql.NullInt64{
      Valid:   true,
      Int64: val,
    },
  }, nil
}

func (nt NullInt64) ToValue() interface{} {
  if !nt.Valid {
    return nil
  }

  return nt.Int64
}

func (nt NullInt64) MarshalJSON() ([]byte, error) {

	if !nt.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nt.Int64)
}

func (nt *NullInt64) UnmarshalJSON(b []byte) error {
	nt.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {
		if err := json.Unmarshal(b, &nt.Int64); err != nil {
			return err
		}
		nt.Valid = true
	}

	return nil
}

func (nt *NullInt64) Scan(value interface{}) error {
	nt.Valid = false

	if value == nil {
		return nil
	}

	var ok bool
	nt.Int64, ok = value.(int64)
	if !ok {

		str, ok := value.(string)
		if !ok {
			return errors.New(fmt.Sprintf("Unable to parse value '%s'", value))
		}

		var err error
		nt.Int64, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
	}

	nt.Valid = true
	return nil
}

func (nt NullInt64) Value() (driver.Value, error) {

	if !nt.Valid {
		return nil, nil
	}
	return nt.Int64, nil
}
