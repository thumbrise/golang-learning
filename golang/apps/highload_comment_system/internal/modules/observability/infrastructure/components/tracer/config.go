package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	OTLPURL string `env:"OBSERVABILITY_OTLP_URL" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
