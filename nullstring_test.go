package jsql_test

import (
	"encoding/json"
	"testing"

	"github.com/joernlenoch/go-jsql"
)

func TestNullString_UnmarshalJSON(t *testing.T) {

	s := jsql.NullString{}

	json.Unmarshal([]byte("null"), &s)
	if s.Valid || s.String != "" {
		t.Fatal("null is not correctly unmarshalled", s)
	}

	json.Unmarshal([]byte(`"null"`), &s)
	if !s.Valid || s.String != "null" {
		t.Fatal("string 'null' is not correctly unmarshalled")
	}

	json.Unmarshal([]byte(`"漢字"`), &s)
	if !s.Valid || s.String != "漢字" {
		t.Fatal("'kanji' is not correctly unmarshalled")
	}

	var nested struct {
		Something jsql.NullString `json:"something"`
	}

	// Possitive testing nested unmarshalling
	err := json.Unmarshal([]byte(`{ "something": "val" }`), &nested)
	if err != nil {
		t.Fatal(err)
	}

	if nested.Something.Valid == false || nested.Something.String != "val" {
		t.Fatal("The nested unmarshalling did not work.", nested)
	}

	// Trying to unmarshal an nested nil value
	err = json.Unmarshal([]byte(`{ "something": null }`), &nested)
	if err != nil {
		t.Fatal(err)
	}

	if nested.Something.Valid != false || nested.Something.String != "" {
		t.Fatal("The nested unmarshalling did not work.", nested)
	}

}

func TestNullString_MarshalJSON_Valid(t *testing.T) {

	s := jsql.NewNullString("")

	data, err := s.MarshalJSON()

	if err != nil {
		t.Fatal(err)
		return
	}

	if string(data) != `""` {
		t.Fail()
	}

	s.String = "漢字"

	data, _ = s.MarshalJSON()
	if string(data) != `"漢字"` {
		t.Fail()
	}
}

// Test the behaviour of marshalling an invalid behaviour
func TestNullString_MarshalJSON_Invalid(t *testing.T) {

	s := jsql.NewNullString(nil)

	data, _ := s.MarshalJSON()
	if string(data) != "null" {
		t.Fail()
	}
}

func TestNullString_Scan(t *testing.T) {
	var s jsql.NullString
	if err := s.Scan(nil); err != nil {
		t.Fatal(err)
	}

	if s.Valid || s.String != "" {
		t.Fatal("nil is not correctly scanned", s)
	}

	if err := s.Scan(""); err != nil {
		t.Fatal(err)
	}

	if !s.Valid || s.String != "" {
		t.Fatal("empty is not correctly scanned", s)
	}

	if err := s.Scan("漢字"); err != nil {
		t.Fatal(err)
	}

	if !s.Valid || s.String != "漢字" {
		t.Fatal("Kanji is not correctly scanned", s)
	}
}

func TestNullString_Value(t *testing.T) {

	s := jsql.NewNullString("asd")

	data, _ := s.Value()
	val := data.(string)

	if val != "asd" {
		t.Fail()
	}

	s.String = "漢字"
	data, _ = s.Value()
	val = data.(string)

	if val != "漢字" {
		t.Fail()
	}

}
