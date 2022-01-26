package jsql_test

import (
  "github.com/joernlenoch/go-jsql"
  "testing"
)

func TestIsNil(t *testing.T) {

	if !jsql.IsNil(nil) {
		t.Errorf("nil must be nil")
		t.Fail()
	}

	// Check nil slices.
	var x []byte
	x = nil

	if !jsql.IsNil(x) {
		t.Errorf("nil slice must be nil")
		t.Fail()
	}

	x = []byte("")
	if jsql.IsNil(x) {
		t.Errorf("empty slice must not be nil")
		t.Fail()
	}

	x = []byte("null")
	if jsql.IsNil(x) {
		t.Errorf("slice with text null must not be nil")
		t.Fail()
	}

	// Check basic pointers on base types.
	var y *string
	y = nil

	if !jsql.IsNil(y) {
		t.Errorf("nil pointer must be nil")
		t.Fail()
	}

	yt := ""
	y = &yt

	if jsql.IsNil(y) {
		t.Errorf("filled pointer must not be nil")
		t.Fail()
	}

	yt2 := "null"
	y = &yt2

	if jsql.IsNil(y) {
		t.Errorf("pointer with value 'null' must not be nil")
		t.Fail()
	}

	// Check structs
	type simpleStruct struct {
		someValue int
	}

	if jsql.IsNil(simpleStruct{}) {
		t.Errorf("zero struct must not be nil")
		t.Fail()
	}

	if jsql.IsNil(&simpleStruct{}) {
		t.Errorf("pointer to zero struct must not be nil")
		t.Fail()
	}

	var nst simpleStruct
	if jsql.IsNil(nst) {
		t.Errorf("zero struct must not be nil")
		t.Fail()
	}

	var st *simpleStruct

	if !jsql.IsNil(st) {
		t.Errorf("nil pointer to struct must be nil")
		t.Fail()
	}

	st = &simpleStruct{}
	if jsql.IsNil(st) {
		t.Errorf("pointer to zero struct must not be nil")
		t.Fail()
	}
}
