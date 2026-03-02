package core

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

var Bindings = wire.NewSet(
	NewModule,
	NewKernel,
	app.NewConfig,
	bootstrap.NewEventLogger,
	bootstrap.NewBootstrapper,
	bootstrap.NewRunner,
	app.NewLoader,
	wire.Bind(
		new(contracts.EnvLoader),
		new(*app.Loader),
	),
	wire.Bind(
		new(contracts.CmdRegistrar),
		new(*Kernel),
	),
)

type Module struct{}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return "core"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
