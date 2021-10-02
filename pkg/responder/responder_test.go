// Package responder is a helper utility for sending all types of responds
package responder

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {
	assert.NotNil(t, NewResponder())
}

func TestJSON(t *testing.T) {
	type mockResponse struct {
		MockField string `json:"field"`
	}
	responder := NewResponder()

	recorder := httptest.NewRecorder()
	responder.JSON(recorder, mockResponse{MockField: "test"}, http.StatusOK)
	assert.Equal(t, recorder.Code, 200)
	assert.JSONEq(t, `{"field": "test"}`, recorder.Body.String())
}
