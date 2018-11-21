package jsql_test

import (
	"github.com/joernlenoch/go-jsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullFloat64_TrySet(t *testing.T) {
	ns2 := jsql.NullFloat64{}

	// Set a normal object
	ns1 := jsql.NewNullFloat64(1.2)

	assert.NoError(t, ns2.TrySet(ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, float64(1.2), ns2.Float64)

	// Set an pointer to an object
	ns1.Set(-5.4)

	assert.NoError(t, ns2.TrySet(&ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, float64(-5.4), ns2.Float64)
}
