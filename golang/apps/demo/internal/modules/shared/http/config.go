package http

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	Port string `env:"HTTP_PORT" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
