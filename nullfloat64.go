package jsql

import (
  "bytes"
  "database/sql"
  "database/sql/driver"
  "encoding/json"
  "fmt"
  "strconv"
  "errors"
)

type NullFloat64 sql.NullFloat64

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
    if err := json.Unmarshal(b, &nt.Float64); err != nil {
      return err
    }
    nt.Valid = true
  }

  return nil
}

func (nt *NullFloat64) Scan(value interface{}) (error) {
  nt.Valid = false

  str, ok := value.(string)
  if !ok {
    return errors.New(fmt.Sprintf("Unable to parse value '%s'", value))
  }

  var err error
  nt.Float64, err = strconv.ParseFloat(str, 64)
  if err != nil {
    return err
  }

  nt.Valid = true
  return nil
}

func (nt NullFloat64) Value() (driver.Value, error) {

  if !nt.Valid {
    return nil, nil
  }
  return nt.Float64, nil
}
