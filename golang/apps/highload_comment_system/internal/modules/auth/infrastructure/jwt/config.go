package jwt

import "github.com/thumbrise/demo/golang-demo/internal/contracts"

type Config struct {
	Secret            string `env:"AUTH_JWT_SECRET"              env-required:"true"`
	Issuer            string `env:"AUTH_JWT_ISSUER"              env-required:"true"`
	AccessTTLMinutes  int    `env:"AUTH_JWT_ACCESS_TTL_MINUTES"  env-required:"true"`
	RefreshTTLMinutes int    `env:"AUTH_JWT_REFRESH_TTL_MINUTES" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
