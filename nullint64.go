package jsql

import (
  "bytes"
  "database/sql"
  "database/sql/driver"
  "encoding/json"
  "strconv"
  "github.com/kataras/go-errors"
  "fmt"
)

type NullInt64 sql.NullInt64

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

func (nt *NullInt64) Scan(value interface{}) (error) {
  nt.Valid = false

  str, ok := value.(string)
  if !ok {
    return errors.New(fmt.Sprintf("Unable to parse value '%s'", value))
  }

  var err error
  nt.Int64, err = strconv.ParseInt(str, 10, 64)
  if err != nil {
    return err
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
