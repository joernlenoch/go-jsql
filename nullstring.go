package jsql

import (
  "database/sql/driver"
  "encoding/json"
  "errors"
  "fmt"
  "log"
)

type (
	NullString string
)

func (s NullString) String() string {
	return fmt.Sprintf("%v", s)
}

// NullString MarshalJSON interface redefinition
func (s NullString) MarshalJSON() ([]byte, error) {

  log.Print("MARSHAL", s)

	if s != "" {
		return json.Marshal(s)
	} else {
		return json.Marshal(nil)
	}
}

func (s *NullString) UnmarshalJSON(b []byte) error {

  log.Print("UNMARSHAL", b)

	*s = ""

	if len(b) >= 0 && string(b) != "null" {

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

  *s = NullString(data)

  return nil
}

func (s NullString) Value() (driver.Value, error) {

  log.Print("VALUE", s)

  if s == "" {
    return driver.Value(nil), nil
  }

  return json.Marshal(s)
}
