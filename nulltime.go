package jsql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type NullTime struct {
	mysql.NullTime
}

func NewNullTime(i interface{}) NullTime {
	nt := NullTime{}
	nt.Set(i)
	return nt
}

func (nt *NullTime) Set(i interface{}) {
	nt.TrySet(i)
}

func (nt *NullTime) TrySet(i interface{}) error {

	if IsNil(i) {
		nt.Time = time.Time{}
		nt.Valid = false
		return nil
	}

	// If the given data is a NullArray object, copy the data directly
	if copy, ok := i.(*NullTime); ok {
		nt.Valid = copy.Valid
		nt.Time = copy.Time
		return nil
	} else if copy, ok := i.(NullTime); ok {
		nt.Valid = copy.Valid
		nt.Time = copy.Time
		return nil
	}

	if val, ok := i.(*time.Time); ok {
		nt.Valid = true
		nt.Time = *val
		return nil
	} else if val, ok := i.(time.Time); ok {
		nt.Valid = true
		nt.Time = val
		return nil
	}

	nt.Time = time.Time{}
	nt.Valid = false
	return fmt.Errorf("expected time value, got: %T", i)
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

// ToValue transform the current value into nil or time.Time
func (nt NullTime) ToValue() interface{} {
	if !nt.Valid {
		return nil
	}

	return nt.Time
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
