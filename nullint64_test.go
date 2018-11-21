package jsql_test

import (
	"github.com/joernlenoch/go-jsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullInt64_TrySet(t *testing.T) {
	ns2 := jsql.NullInt64{}

	// Set a normal object
	ns1 := jsql.NewNullInt64(1)

	assert.NoError(t, ns2.TrySet(ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, int64(1), ns2.Int64)

	// Set an pointer to an object
	ns1.Set(-5)

	assert.NoError(t, ns2.TrySet(&ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, int64(-5), ns2.Int64)
}
