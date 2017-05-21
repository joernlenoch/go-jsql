package jsql

import (
  "bytes"
  "database/sql/driver"
  "encoding/json"
  "fmt"
  "time"

  "github.com/go-sql-driver/mysql"
  "github.com/juju/errors"
)

type NullTime mysql.NullTime

func (nt NullTime) MarshalJSON() ([]byte, error) {

  if !nt.Valid {
    return []byte("null"), nil
  }

  return json.Marshal(nt.Time)
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {

  nt.Time = time.Time{}
  nt.Valid = false

  if bytes.Equal(b, []byte("null")) {
    return nil
  }

  if len(b) >= 0 {

    // Try to extract the 'string'. If this failed we simply
    // use the base value as string.
    if err := json.Unmarshal(b, &nt.Time); err != nil {
      return err
    }

    nt.Valid = true
  }

  return nil
}


// Scan implements the Scanner interface.
// The value type must be time.Time or string / []byte (formatted time-string),
// otherwise Scan fails.
func (nt *NullTime) Scan(value interface{}) (err error) {

  // TODO Fix me - for some reason nested methods are not found.
  // This should work for now...
  f := mysql.NullTime{}
  if err := f.Scan(value); err != nil {
    return errors.New(fmt.Sprintf("Unable to parse time value: %s", err.Error()))
  }

  nt.Valid, nt.Time = f.Valid, f.Time
  return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
  if !nt.Valid {
    return nil, nil
  }
  return nt.Time, nil
}
