package jsql

import (
  "database/sql/driver"
  "encoding/json"
  "errors"
  "fmt"
  "bytes"
)

type (
  // Basically a clone of the sql.NullString, but with
  // additional functionality like JSON marshalling.
	NullString struct {
    String string
    Valid bool
  }
)

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

		if err := json.Unmarshal(b, &s.String); err != nil {
			return err
		}

    s.Valid = true
	}

	return nil
}

func (s *NullString) Scan(src interface{}) error {

  if src == nil {
    s.String = ""
    s.Valid = false
    return nil
  }

  switch src.(type) {
  case []byte:
    s.String = string(src.([]byte))
    s.Valid = true
  case string:
    s.String = src.(string)
    s.Valid = true
  default:
    return errors.New(fmt.Sprintf("The given data is not a valid []byte. (%#v)", src))
  }

  return nil
}

func (s NullString) Value() (driver.Value, error) {

  if !s.Valid {
    return driver.Value(nil), nil
  }

  return s.String, nil
}
