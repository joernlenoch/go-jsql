package jsql

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
)

type NullTime struct {
	mysql.NullTime
}

func NewNullTime(s interface{}) NullTime {
	if val, ok := s.(time.Time); ok {
		return NullTime{
			NullTime: mysql.NullTime{
				Valid: true,
				Time:  val,
			},
		}
	}

	return NullTime{
		NullTime: mysql.NullTime{
			Valid: false,
		},
	}
}

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
