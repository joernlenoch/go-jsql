package jsql

import (
  "github.com/go-sql-driver/mysql"
  "encoding/json"
  "bytes"
  "time"
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
