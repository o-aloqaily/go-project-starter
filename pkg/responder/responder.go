// Package responder is a helper utility for sending all types of responds
package responder

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Responder is the interface of the responder type
type Responder interface {
	// JSON marshals 'v' to JSON, sets the Content-Type header as application/json
	// and sets the http responde status as s
	JSON(w http.ResponseWriter, v interface{}, s int)
}

type responder struct{}

// NewResponder is the factory method / constructor of the type Responder
func NewResponder() Responder {
	return &responder{}
}

func (r *responder) JSON(w http.ResponseWriter, v interface{}, s int) {
	// Write response type header
	w.Header().Set("Content-Type", "application/json")
	// Write http response status
	w.WriteHeader(s)
	// Marshal and send the response as JSON
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}
