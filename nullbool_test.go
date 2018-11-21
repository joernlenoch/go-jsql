package jsql_test

import (
	"github.com/joernlenoch/go-jsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullBool_TrySet(t *testing.T) {
	ns2 := jsql.NullBool{}

	// Set a normal object
	ns1 := jsql.NewNullBool(true)

	assert.NoError(t, ns2.TrySet(ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, true, ns2.Bool)

	// Set an pointer to an object
	ns1.Set(false)

	assert.NoError(t, ns2.TrySet(&ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, false, ns2.Bool)
}
