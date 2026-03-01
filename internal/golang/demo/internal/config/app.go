package config

import (
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/contracts"
)

type App struct {
	Name        string `env:"APP_NAME"  env-required:"true"`
	Version     string `default:"1.0.0" env:"APP_VERSION"`
	Environment string `env:"APP_ENVIRONMENT"  env-required:"true"`
}

func NewApp(loader contracts.EnvLoader) App {
	cfg := App{}

	loader.MustLoad(&cfg)

	return cfg
}
