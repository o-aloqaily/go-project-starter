// Package validator is a validation package used to validate fields in different structures
package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	assert.NotNil(t, NewValidator())
}

// No need to write more unit tests for the validator
// as it's just a wrapper around the go-validator package
