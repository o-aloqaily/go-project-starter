// Package dosomething defines and implements the dosomething function
package dosomething

import (
	"github.com/o-aloqaily/go-project-starter/internal/errors"
	"github.com/o-aloqaily/go-project-starter/internal/apiclient"
)

// Service encapsulates usecase logic for the dosomething package.
type Service interface {
	DoSomething(field1 string, field2 string) error
}

type service struct {
	api apiclient.Client
}

// NewService is the service's factory method (constructor)
func NewService(api apiclient.Client) Service {
	return &service{
		api: api,
	}
}

func (s *service) DoSomething(field1 string, field2 string) error {
	req := apiclient.CallSomethingRequest{
		Field1: field1,
	}
	err := s.api.CallSomething(req)
	// return err if any error occurs
	if err != nil {
		return errors.ErrorResponse{
			Status:  500,
			Message: err.(apiclient.CallSomethingResponseErr).ErrorMessage,
		}
	}
	return nil
}
