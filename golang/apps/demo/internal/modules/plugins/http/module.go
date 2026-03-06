package http

import (
	"context"

	"github.com/gin-contrib/pprof"
	"github.com/google/wire"
	components2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/components"
	observability2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/observability"
)

var Bindings = wire.NewSet(
	NewModule,
	components2.NewGinEngine,
	components2.NewKernel,
	components2.NewConfig,
	components2.NewSlogginConfig,

	observability2.NewHealthRouter,
	observability2.NewOTELMiddleware,
)

type Module struct {
	kernel         *components2.Kernel
	healthRouter   *observability2.HealthRouter
	otelMiddleware *observability2.OTELMiddleware
}

func NewModule(healthRouter *observability2.HealthRouter, kernel *components2.Kernel, otelMiddleware *observability2.OTELMiddleware) *Module {
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
