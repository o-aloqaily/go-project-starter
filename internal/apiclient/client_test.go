// Package apiclient provides an API Client for accessing some random API
package apiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/o-aloqaily/go-project-starter/internal/config"
	"github.com/o-aloqaily/go-project-starter/pkg/logger"

	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := &config.Config{
		DbDsn:      "dbdsnmock",
		ApiBaseURL: "apibaseurl",
	}
	assert.NotNil(t, NewClient(mockConfig, logger.NewMockLogger(ctrl)))
}

func TestCallSomething_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := logger.NewMockLogger(ctrl)

	apiReq := CallSomethingRequest{Field1: "blablabla"}
	// mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.WriteHeader(200)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := client{
		BaseURL:    server.URL,
		AuthToken:  "dummyauthtoken",
		log:        mockObj,
		httpClient: server.Client(),
	}
	err := c.CallSomething(apiReq)
	assert.Nil(t, err)
}

func TestCallSomething_Err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := logger.NewMockLogger(ctrl)
	mockObj.EXPECT().Error(errorMessage)

	apiReq := CallSomethingRequest{Field1: "blablabla"}
	// mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.WriteHeader(400)
		json.NewEncoder(rw).Encode(&CallSomethingResponseErr{
			ErrorMessage: "dummy",
			ErrorCode:    "111",
		})
	}))

	// Close the server when test finishes
	defer server.Close()

	c := client{
		BaseURL:    server.URL,
		AuthToken:  "dummyauthtoken",
		log:        mockObj,
		httpClient: server.Client(),
	}
	err := c.CallSomething(apiReq)
	assert.NotNil(t, err)
}
