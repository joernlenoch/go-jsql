package jsql

import (
	"database/sql"
	"encoding/json"
)

type (
	NullString struct {
		sql.NullString
	}
)

// NullString MarshalJSON interface redefinition
func (r NullString) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.String)
	} else {
		return json.Marshal(nil)
	}
}

func (r *NullString) UnmarshalJSON(b []byte) error {

	r.Valid = false
	r.String = ""

	if len(b) >= 0 && string(b) != "null" {

		if err := json.Unmarshal(b, &r.String); err != nil {
			return err
		}

		r.Valid = true
	}

	return nil
}
