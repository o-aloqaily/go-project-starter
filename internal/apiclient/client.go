// Package apiclient provides an API Client for accessing some random API
package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/o-aloqaily/go-project-starter/internal/config"
	"github.com/o-aloqaily/go-project-starter/pkg/logger"
)

// DO NOT DELETE THE FOLLOWING COMMENT
// This is to generate a mock logger
//go:generate mockgen -source=./client.go -destination=./mock_client.go -package=apiclient

// Client is the interface of type client
type Client interface {
	CallSomething(req CallSomethingRequest) error
}

type client struct {
	// Base URL of the apiclient API.
	// Should be populated using the factory/constructor method from the service's config
	BaseURL string
	// AuthToken is the authorization header
	AuthToken string
	// log is the logger to be used for logging while using the apiclient API
	log logger.Logger
	// httpClient is the client to be used for sending http requests
	httpClient *http.Client
}

// CallSomethingRequest represents the request body to be sent to the API's endpoint
type CallSomethingRequest struct {
	Field1 string `json:"field1"`
}

type CallSomethingResponseOK struct {
	Field1 string `json:"field1"`
}

// CallSomethingResponseErr will be used to unmarshal error messages received from the API
type CallSomethingResponseErr struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
}

var errorMessage string = "The API returned a response with a status code indicating an error"

func (err CallSomethingResponseErr) Error() string {
	return err.ErrorMessage
}

// NewClient is the factory/constructor method of the api client type
func NewClient(c *config.Config, log logger.Logger) Client {
	return &client{
		BaseURL: c.ApiBaseURL,
		log:     log,
	}
}

// CallSomething calls some endpoint on the API.....
func (c *client) CallSomething(req CallSomethingRequest) error {

	// Construct request
	urlToEndpoint := c.BaseURL + "/<pathToEndpoint>"
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(req)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}

	// Create the request and add headers
	r, err := http.NewRequest("POST", urlToEndpoint, b)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}
	r.Header.Add("Content-type", "application/json")
	r.Header.Add("Authorization", "Basic "+c.AuthToken)

	// Send the request
	res, err := c.httpClient.Do(r)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}

	defer res.Body.Close()

	// Decode the response and return an error if any
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		c.log.Error(errorMessage)
		var errMsg CallSomethingResponseErr
		json.NewDecoder(res.Body).Decode(&errMsg)
		return errMsg
	}

	// Return nil if no errors occured
	return nil
}
