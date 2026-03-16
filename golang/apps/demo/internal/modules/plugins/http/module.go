package http

import (
	"context"

	"github.com/gin-contrib/pprof"
	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/observability"
)

var Bindings = wire.NewSet(
	NewModule,

	components.NewGinEngine,
	components.NewKernel,
	components.NewConfig,
	components.NewSlogginConfig,

	observability.NewHealthRouter,
	observability.NewOTELMiddleware,
	observability.NewHTTPMetrics,
	observability.NewOtelRecorder,
)

type Module struct {
	kernel         *components.Kernel
	healthRouter   *observability.HealthRouter
	otelMiddleware *observability.OTELMiddleware
}

func NewModule(healthRouter *observability.HealthRouter, kernel *components.Kernel, otelMiddleware *observability.OTELMiddleware) *Module {
	return &Module{healthRouter: healthRouter, kernel: kernel, otelMiddleware: otelMiddleware}
}

func (m *Module) Name() string {
	return "http"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.kernel.Gin().Use(m.otelMiddleware.Handler(ctx))
	m.healthRouter.Register()
	pprof.Register(m.kernel.Gin())

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
