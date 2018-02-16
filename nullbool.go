package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
  "strconv"
  "fmt"
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
  var val bool
  var err error

  switch i.(type) {
  case bool:
    val = i.(bool)
  default:
    val, err = strconv.ParseBool(fmt.Sprint(i))
  }

  if err != nil {
    return NullBool{
      sql.NullBool{
        Valid: false,
      },
    }, err
  }

  return NullBool{
    sql.NullBool{
      Valid:   true,
      Bool: val,
    },
  }, nil
}

func (nt NullBool) ToValue() interface{} {
  if !nt.Valid {
    return nil
  }

  return nt.Bool
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
