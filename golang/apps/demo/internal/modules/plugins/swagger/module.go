package swagger

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/swagger/endpoints/http"
)

var Bindings = wire.NewSet(
	NewModule,
	http.NewSwaggerRouter,
)

type Module struct {
	swaggerRouter *http.SwaggerRouter
}

func NewModule(swaggerRouter *http.SwaggerRouter) *Module {
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
