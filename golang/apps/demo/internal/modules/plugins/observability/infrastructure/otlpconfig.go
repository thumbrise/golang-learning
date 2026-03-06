package infrastructure

import "github.com/thumbrise/demo/golang-demo/internal/contracts"

type OTLPConfig struct {
	URL string `env:"OBSERVABILITY_OTLP_URL" env-required:"true"`
}

func NewOTLPConfig(loader contracts.EnvLoader) OTLPConfig {
	cfg := OTLPConfig{}

	loader.MustLoad(&cfg)

	return cfg
}
