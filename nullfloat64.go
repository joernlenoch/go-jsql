package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strconv"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func NewNullFloat64(s interface{}) NullFloat64 {
	if val, ok := s.(float64); ok {
		return NullFloat64{
			sql.NullFloat64{
				Valid:   true,
				Float64: val,
			},
		}
	}

	return NullFloat64{
		sql.NullFloat64{
			Valid: false,
		},
	}
}

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
