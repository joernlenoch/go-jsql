package jsql

import (
  "reflect"
)

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	if reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil() {
		return true
	}

	if reflect.ValueOf(i).Kind() == reflect.Slice && reflect.ValueOf(i).IsNil() {
		return true
	}

	return false
}
