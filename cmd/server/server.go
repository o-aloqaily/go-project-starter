package main

import (
	"github.com/o-aloqaily/go-project-starter/pkg/logger"
	"github.com/o-aloqaily/go-project-starter/pkg/router"
	"github.com/o-aloqaily/go-project-starter/internal/config"
)

// Server struct will hold shared resources like db connection, config variables, logger...
type Server struct {
	log        *logger.Logger
	config     *config.Config
	rootRouter router.Router
}

// NewServer returns a new server instance with the dependencies passed in
func NewServer(log logger.Logger, config *config.Config, rootRouter router.Router) *Server {
	return &Server{
		log:        &log,
		config:     config,
		rootRouter: rootRouter,
	}
}