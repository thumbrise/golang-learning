package config

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Observability struct {
	OTLPURL      string `env:"OBSERVABILITY_OTLP_URL"      env-required:"true"`
	PyroscopeURL string `env:"OBSERVABILITY_PYROSCOPE_URL" env-required:"true"`
}

func NewObservability(loader contracts.EnvLoader) Observability {
	cfg := Observability{}

	loader.MustLoad(&cfg)

	return cfg
}
