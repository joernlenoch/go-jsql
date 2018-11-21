package jsql_test

import (
	"github.com/joernlenoch/go-jsql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNullTime_TrySet(t *testing.T) {
	ns2 := jsql.NullTime{}

	// Set a normal object
	data := time.Now()
	ns1 := jsql.NewNullTime(data)

	assert.NoError(t, ns2.TrySet(ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, data, ns2.Time)

	// Set an pointer to an object
	data2 := time.Unix(1, 0)
	ns1.Set(data2)

	assert.NoError(t, ns2.TrySet(&ns1))
	assert.True(t, ns2.Valid)
	assert.Equal(t, data2, ns2.Time)
}
