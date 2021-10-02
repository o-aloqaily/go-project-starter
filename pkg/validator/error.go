// Package validator is a validation package used to validate fields in different structures
package validator

import (
	validatorpkg "github.com/go-playground/validator/v10"
)

// ValidationErrors will expose functions to retrieve the validation errors
type ValidationErrors interface {
	// To implement go's built-in error interface
	Error() string
	// First returns the first validation error translated into human-readable format
	// This is useful for api's / services to use when returning a message to the client
	First() string
}

// validationErrors will contain a slice of all the field validation errros
type validationErrors struct {
	fieldErrors []validatorpkg.FieldError
}

// First returns the first validation error translated into human-readable format
// This is useful for api's / services to use when returning a message to the client
func (fe validationErrors) First() string {
	// use the translator from the validation.go file
	return fe.fieldErrors[0].Translate(trans)
}

func (fe validationErrors) Error() string {
	// use the translator from the validation.go file
	return fe.fieldErrors[0].Translate(trans)
}
