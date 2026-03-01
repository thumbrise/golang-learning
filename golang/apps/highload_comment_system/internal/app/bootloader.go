package app

import (
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(
		NewConfig,
		bootstrap.NewRunner,
		fx.Annotate(
			NewLoader,
			fx.As(new(contracts.EnvLoader)),
		),
	),
)
