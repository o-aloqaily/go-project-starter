// Package router is a wrapper around the router package or logic being used for routing
package router

import (
	"github.com/go-chi/chi"
)

// Router is a wrapper around a router package or logic being used for routing
type Router interface {
	// Mount(s string, h http.Handler)
	chi.Router
}

type router struct {
	// // Implementing the same functions from chi into the router struct of ours
	// // If we need to switch to a different router later on we can implement
	// // the same functions provided by the chi router in this file
	// // such as Mount
	Router
}

// NewRouter is the factory method / constructor of the router
func NewRouter() Router {
	return &router{chi.NewRouter()}
}
