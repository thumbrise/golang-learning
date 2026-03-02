package core

import (
	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

var Bindings = wire.NewSet(
	app.NewConfig,
	bootstrap.NewRunner,
	app.NewLoader,
	wire.Bind(
		new(contracts.EnvLoader),
		new(*app.Loader),
	),
)
