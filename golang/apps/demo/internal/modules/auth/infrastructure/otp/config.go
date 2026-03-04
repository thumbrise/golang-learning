package otp

import "github.com/thumbrise/demo/golang-demo/internal/contracts"

type Config struct {
	Length      int    `env:"AUTH_OTP_LENGTH"       env-required:"true"`
	TTLMinutes  int    `env:"AUTH_OTP_TTL_MINUTES"  env-required:"true"`
	ForcedValue string `env:"AUTH_OTP_FORCED_VALUE"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
