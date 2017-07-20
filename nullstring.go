package jsql

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
)

type (

	// Basically a clone of the sql.NullString, but with
	// additional functionality like JSON marshalling.
	NullString struct {
		sql.NullString
	}
)

func NewNullString(s interface{}) NullString {
	return NullString{
		NullString: sql.NullString{
			String: fmt.Sprint(s),
			Valid:  s != nil,
		},
	}
}

// NullString MarshalJSON interface redefinition
func (s NullString) MarshalJSON() ([]byte, error) {

	if !s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(b []byte) error {

	s.String = ""
	s.Valid = false

	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	if len(b) >= 0 {

		// Try to extract the 'string'. If this failed we simply
		// use the base value as string.
		if err := json.Unmarshal(b, &s.String); err != nil {
			s.String = string(b)
		}

		s.Valid = true
	}

	return nil
}
