package app

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	Name        string `env:"APP_NAME"        env-required:"true"`
	Version     string `default:"1.0.0"       env:"APP_VERSION"`
	Environment string `env:"APP_ENVIRONMENT" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
