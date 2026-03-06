package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/logger"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/meter"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	observabilitytracer "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

var (
	ErrConfigureRegistrar = errors.New("failed configure registrar")
	Bindings              = wire.NewSet(
		NewModule,

		infrastructure.NewOTLPConfig,

		components.NewResource,
		components.NewRegistrar,
		components.NewErrorHandler,

		profiler.NewConfig,
		profiler.NewProfiler,

		logger.NewLogger,
		logger.NewOTELSDKProvider,
		logger.NewProvider,
		logger.NewExporter,

		meter.NewOTELSDKProvider,
		meter.NewProvider,

		observabilitytracer.NewProvider,
		observabilitytracer.NewOTELSampler,
		observabilitytracer.NewOTELExporter,
		observabilitytracer.NewOTELSDKProvider,
		observabilitytracer.NewStdOutExporter,

		wire.Bind(new(oteltracer.TracerProvider), new(*sdktrace.TracerProvider)),

		routers.NewHealthRouter,
		routers.NewObservabilityRouter,
		routers.NewPprofRouter,

		middlewares.NewObservabilityMiddleware,
	)
)

type Module struct {
	healthRouter        *routers.HealthRouter
	pprofRouter         *routers.PprofRouter
	observabilityRouter *routers.ObservabilityRouter
	tracerErrorHandler  *components.ErrorHandler
	registrar           *components.Registrar
}

func NewModule(healthRouter *routers.HealthRouter, observabilityRouter *routers.ObservabilityRouter, pprofRouter *routers.PprofRouter, telemetry *components.Registrar, tracerErrorHandler *components.ErrorHandler) *Module {
	return &Module{healthRouter: healthRouter, observabilityRouter: observabilityRouter, pprofRouter: pprofRouter, registrar: telemetry, tracerErrorHandler: tracerErrorHandler}
}

func (m *Module) Name() string {
	return "observability"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	err := m.registrar.Configure(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfigureRegistrar, err)
	}

	m.observabilityRouter.Register()
	m.healthRouter.Register()
	m.pprofRouter.Register()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return m.registrar.Shutdown(ctx)
}
