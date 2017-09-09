package jsql

import (
	"bytes"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
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

func (nt NullTime) IsExpired() bool {
	if !nt.Valid {
		return false
	}

	return time.Now().After(nt.Time)
}

func (nt NullTime) Before(t interface{}) bool {
	if !nt.Valid {
		return false
	}

	if ntt, ok := t.(time.Time); ok {
		return nt.Time.Before(ntt)
	}

	if ntt, ok := t.(NullTime); ok {
		if !ntt.Valid {
			return true
		}

		return nt.Time.Before(ntt.Time)
	}

	log.Print("[Warning] The given time is not a valid time")
	return false
}

func (nt NullTime) After(t interface{}) bool {
	if !nt.Valid {
		return false
	}

	if ntt, ok := t.(time.Time); ok {
		return nt.Time.After(ntt)
	}

	if ntt, ok := t.(NullTime); ok {
		if !ntt.Valid {
			return true
		}
		return nt.Time.After(ntt.Time)
	}

	log.Print("[Warning] The given time is not a valid time")
	return false
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
