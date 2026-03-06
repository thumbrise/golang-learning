package swagger

import (
	"context"

	"github.com/google/wire"
)

var Bindings = wire.NewSet(
	NewModule,
	NewSwaggerRouter,
)

type Module struct {
	swaggerRouter *SwaggerRouter
}

func NewModule(swaggerRouter *SwaggerRouter) *Module {
	return &Module{swaggerRouter: swaggerRouter}
}

func (m *Module) Name() string {
	return "swagger"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.swaggerRouter.Register()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
