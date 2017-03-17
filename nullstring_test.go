package jsql_test

import (
  "testing"
  "github.com/j-lenoch/jsql"
)


func TestNullString_UnmarshalJSON(t *testing.T) {

  s := jsql.NullString{}
  s.UnmarshalJSON([]byte("null"))

  if s.Valid || s.String != "" {
    t.Fatal("null is not correctly unmarshalled", s)
  }

  s.UnmarshalJSON([]byte(`"null"`))

  if !s.Valid || s.String != "null" {
    t.Fatal("string 'null' is not correctly unmarshalled")
  }

  s.UnmarshalJSON([]byte(`"漢字"`))

  if !s.Valid || s.String != "漢字" {
    t.Fatal("'kanji' is not correctly unmarshalled")
  }

}

func TestNullString_MarshalJSON_Valid(t *testing.T) {

  s := jsql.NullString{
    Valid: true,
    String: "",
  }

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

  s := jsql.NullString{
    Valid: false,
    String: "asd",
  }

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

  s := jsql.NullString{
    Valid: true,
    String: "asd",
  }

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
