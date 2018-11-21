package jsql_test

import (
	"github.com/joernlenoch/go-jsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullArray_TrySet(t *testing.T) {
	ns2 := jsql.NullArray{}

	// Set a normal object
	data := []string{"a", "b"}
	ns1 := jsql.NewNullArray(data)

	assert.NoError(t, ns2.TrySet(ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, []interface{}{"a", "b"}, ns2.Array)

	// Set an pointer to an object
	data2 := []int{1, 2}
	ns1.Set(data2)

	assert.NoError(t, ns2.TrySet(&ns1))
	assert.True(t, ns2.Valid)
	assert.EqualValues(t, []interface{}{1, 2}, ns2.Array)
}
