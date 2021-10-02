//+build wireinject

// dosomething package
package dosomething

import (
	"github.com/o-aloqaily/go-project-starter/internal/apiclient"
	"github.com/o-aloqaily/go-project-starter/internal/config"
	"github.com/o-aloqaily/go-project-starter/internal/interfaces"
	"github.com/o-aloqaily/go-project-starter/pkg/logger"
	"github.com/o-aloqaily/go-project-starter/pkg/responder"
	"github.com/o-aloqaily/go-project-starter/pkg/validator"

	"github.com/google/wire"
)

func InitializeDoSomethingAPI() interfaces.API {
	wire.Build(
		logger.NewLogger,
		config.NewConfig,
		apiclient.NewClient,
		responder.NewResponder,
		validator.NewValidator,
		NewService,
		NewAPI,
	)
	return &api{}
}
