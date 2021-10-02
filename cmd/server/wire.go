//+build wireinject
// Wire file for automated dependancy injection

package main

import (
	"github.com/o-aloqaily/go-project-starter/pkg/logger"
	"github.com/o-aloqaily/go-project-starter/pkg/router"
	"github.com/o-aloqaily/go-project-starter/internal/config"

	"github.com/google/wire"
)

func InitializeServer() *Server {
	wire.Build(
		logger.NewLogger,
		config.NewConfig,
		NewServer,
		router.NewRouter,
	)
	return &Server{}
}
