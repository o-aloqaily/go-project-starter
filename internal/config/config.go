// Package config is used to load configurations
// from env vars to a config struct
package config

import (
	"os"

	"github.com/o-aloqaily/go-project-starter/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/qiangxue/go-env"
)

// Config holds the configuration variables loaded from env vars
type Config struct {
	DbDsn      string
	ApiBaseURL string
}

// NewConfig returns a pointer to a loaded configurations object
// It checks for APP_ENV variable in env vars, if the variable was not found
// or is not set to production, it loads the env vars from the config/.env file
func NewConfig(log logger.Logger) *Config {

	c := new(Config)

	// If we're in the development environment, load env variables from the .env file
	// In production, environment variables will be injected into the app using k8s secrets
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load("../../config/.env")
		if err != nil {
			log.Fatal("Error loading .env file: ", err.Error())
		}
	}

	// Populate the config object with values from env vars
	if err := env.New("APP_", log.Infof).Load(c); err != nil {
		log.Fatal("Error unmarshalling environment variables", err.Error())
	}

	return c
}
