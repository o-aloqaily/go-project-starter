// Package router is a wrapper around the router package or logic being used for routing
package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	assert.NotNil(t, NewRouter())
}
