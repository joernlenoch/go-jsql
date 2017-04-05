package jsql

import (
  "bytes"
  "database/sql/driver"
  "encoding/json"
  "strconv"
  "errors"
  "fmt"
)

type NullArray struct {
  Valid bool
  Array []interface{}
}

func (nt NullArray) ToStringArray() []string {
  out := make([]string, 0, len(nt.Array))

  for i, val := range nt.Array {
    out[i] = fmt.Sprint(val)
  }

  return out
}

func (nt NullArray) ToInt64Array() ([]int64, error) {
  out := make([]int64, 0, len(nt.Array))

  var err error
  for i, val := range nt.Array {
    if i, ok := val.(int64); ok {
      out[i] = i
      continue
    }

    str := fmt.Sprint(val)
    out[i], err = strconv.ParseInt(str, 10, 64)
    if err != nil {
      return nil, err
    }
  }
  return out, nil
}

func (nt NullArray) ToFloat64Array() ([]float64, error) {
  out := make([]float64, 0, len(nt.Array))

  var err error
  for i, val := range nt.Array {
    if f, ok := val.(float64); ok {
      out[i] = f
      continue
    }
    str := fmt.Sprint(val)
    out[i], err = strconv.ParseFloat(str, 64)
    if err != nil {
      return nil, err
    }
  }
  return out, nil
}

func (nt NullArray) MarshalJSON() ([]byte, error) {
  if !nt.Valid {
    return []byte("null"), nil
  }
  return json.Marshal(nt.Array)
}

func (nt *NullArray) UnmarshalJSON(b []byte) error {
  nt.Valid = false

  if bytes.Equal(b, []byte("null")) {
    return nil
  }

  if len(b) >= 0 {
    if err := json.Unmarshal(b, &nt.Array); err != nil {
      return err
    }
    nt.Valid = true
  }

  return nil
}

// Retrieve value as JSON data
func (nt *NullArray) Scan(value interface{}) (err error) {
  nt.Valid = false

  if value == nil {
    return
  }

  data, ok := value.([]byte)
  if !ok {
    return errors.New("The given data is not a valid string")
  }

  return nt.UnmarshalJSON(data)
}

func (nt NullArray)  Value() (driver.Value, error) {

  if !nt.Valid {
    return nil, nil
  }

  return nt.MarshalJSON()
}
