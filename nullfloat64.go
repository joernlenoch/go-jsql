package jsql

import (
  "bytes"
  "database/sql"
  "database/sql/driver"
  "encoding/json"
  "fmt"
  "strconv"
  "errors"
  "reflect"
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

func (nt *NullFloat64) Scan(value interface{}) (error) {
  nt.Valid = false

  if value == nil {
    return nil
  }

  var ok bool
  nt.Float64, ok = value.(float64)
  if !ok {

    // [Try to parse the data from byte or string]
    data, ok := value.([]byte)
    if !ok {
      return errors.New(fmt.Sprintf("Unable to parse value '%s' (%s)", value, reflect.TypeOf(value)))
    }

    var err error
    nt.Float64, err = strconv.ParseFloat(string(data), 64)
    if err != nil {
      return err
    }
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
