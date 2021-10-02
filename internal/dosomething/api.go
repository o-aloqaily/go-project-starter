// Package dosomething.....
package dosomething

import (
	"encoding/json"
	"net/http"

	"github.com/o-aloqaily/go-project-starter/internal/errors"
	"github.com/o-aloqaily/go-project-starter/internal/interfaces"
	"github.com/o-aloqaily/go-project-starter/pkg/logger"
	"github.com/o-aloqaily/go-project-starter/pkg/responder"
	"github.com/o-aloqaily/go-project-starter/pkg/router"
	"github.com/o-aloqaily/go-project-starter/pkg/validator"
)

// Request represents the data in the request body of dosomething
type Request struct {
	Field1 string `validate:"required" json:"field1"`
	Field2 string `validate:"required" json:"field2"`
}

// Response represents the data in the request body of dosomething
type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

// type api implements the interface API (interfaces.API)
type api struct {
	service  Service
	log      logger.Logger
	validate validator.Validator
	respond  responder.Responder
}

// NewAPI is the factory method of the dosomething API
func NewAPI(s Service, log logger.Logger, validate validator.Validator,
	responder responder.Responder) interfaces.API {
	return &api{
		service:  s,
		log:      log,
		validate: validate,
		respond:  responder,
	}
}

func (a *api) Handlers() router.Router {
	r := router.NewRouter()

	// dosomething endpoint
	r.Post("/exampleendpoint", a.doSomething)

	return r
}

func (a *api) doSomething(res http.ResponseWriter, req *http.Request) {
	var r Request

	// Decode request's body
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		a.respond.JSON(res, errors.InternalServerError(), 500)
		return
	}

	// Validate request's body
	if err := a.validate.Struct(r); err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		a.respond.JSON(res, errors.BadRequest(validatorErrs.First()), 400)
		return
	}

	// Execute the service, send an error if any is returned by the service
	if err := a.service.DoSomething(r.Field1, r.Field2); err != nil {
		a.respond.JSON(res, err.(errors.ErrorResponse), err.(errors.ErrorResponse).Status)
		return
	}

	// Return a 200 response
	a.respond.JSON(res, Response{
		Status: 200,
		Data:   "Bla bla bla.",
	}, 200)
}
