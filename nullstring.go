package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
  "errors"
)

type (

	// Basically a clone of the sql.NullString, but with
	// additional functionality like JSON marshalling.
	NullString struct {
		sql.NullString
	}
)

func NewNullString(i interface{}) NullString {
	n, _ := TryNullString(i)
	return n
}

// Create a new NullFloat.
// - nil and numeric values are considered correct
func TryNullString(i interface{}) (NullString, error) {

  if i == nil {
    return NullString{
      sql.NullString{
        Valid:   false,
      },
    }, nil
  }

  var val string
  var err error

  switch i.(type) {
  case string:
    val = i.(string)
  case []byte:
    val = string(i.([]byte))
  default:
    err = errors.New(fmt.Sprintf("given value '%s' is not en explicit string: please cast it to ensure that this behaviour is expected", i))
  }

  if err != nil {
    return NullString{
      sql.NullString{
        Valid: false,
      },
    }, err
  }

  return NullString{
    sql.NullString{
      Valid:   true,
      String: val,
    },
  }, nil
}

func (nt NullString) ToValue() interface{} {
  if !nt.Valid {
    return nil
  }

  return nt.String
}

// NullString MarshalJSON interface redefinition
func (s NullString) MarshalJSON() ([]byte, error) {

	if !s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(b []byte) error {

	s.String = ""
	s.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {

		// Try to extract the 'string'. If this failed we simply
		// use the base value as string.
		if err := json.Unmarshal(b, &s.String); err != nil {
			s.String = string(b)
		}

		s.Valid = true
	}

	return nil
}
