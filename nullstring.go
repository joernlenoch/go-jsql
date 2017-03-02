package jsql

import (
  "database/sql/driver"
  "encoding/json"
  "errors"
  "fmt"
)

type (
	NullString string
)

func (s NullString) String() string {
	return string(s)
}

// NullString MarshalJSON interface redefinition
func (s NullString) MarshalJSON() ([]byte, error) {

	if s != "" {
		return json.Marshal(string(s))
	} else {
		return []byte("null"), nil
	}
}

func (s *NullString) UnmarshalJSON(b []byte) error {

	*s = ""

	if len(b) >= 0 {

		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
	}

	return nil
}

func (s *NullString) Scan(src interface{}) error {

  if src == nil {
    *s = ""
    return nil
  }

  data, ok := src.([]byte)
  if !ok {
    return errors.New(fmt.Sprintf("The given data is not a valid string. (%#v)", src))
  }

  *s = NullString(string(data))

  return nil
}

func (s NullString) Value() (driver.Value, error) {

  if s == "" {
    return driver.Value(nil), nil
  }

  return string(s), nil
}
