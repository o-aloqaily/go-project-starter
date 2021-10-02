// Package validator is a validation package used to validate fields in different structures
package validator

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validatorpkg "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	once sync.Once

	// Singleton instance fo the validator type. This helps in caching struct info.
	v        Validator
	validate *validatorpkg.Validate

	// Error translator
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

// Validator is the interface of validator, provides
type Validator interface {
	Struct(s interface{}) error
}

// Implementation of the interface "Validator"
type validator struct {
	Validator
}

// NewValidator is the factory method (constructor) of type validator
func NewValidator() Validator {
	// Instantiate only a single instance of the validator type
	once.Do(func() {
		validate = validatorpkg.New()
		// Validator will be using go-playground/validator functions for validation
		english := en.New()
		uni = ut.New(english, english)
		trans, _ = uni.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(validate, trans)
		v = &validator{validate}
	})
	// Return the singleton validator
	return v
}

func (va *validator) Struct(s interface{}) error {
	ve := validate.Struct(s)
	if ve != nil {
		return validationErrors{fieldErrors: ve.(validatorpkg.ValidationErrors)}
	}
	return nil
}
