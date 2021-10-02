package main

import (
	"net/http"

	"github.com/o-aloqaily/go-project-starter/internal/dosomething"
)

func main() {
	// Initialize server struct with logger, config and root router
	s := InitializeServer()

	// Mount the api endpoints for example service
	s.rootRouter.Mount("/example", dosomething.InitializeDoSomethingAPI().Handlers())

	// Start listening for incoming requests
	http.ListenAndServe(":8080", s.rootRouter)
}
