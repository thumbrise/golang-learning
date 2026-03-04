package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/logger"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/metrics"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	observabilitytracer "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

var (
	ErrStartProfiler    = errors.New("start profiler")
	ErrShutdownProfiler = errors.New("shutdown profiler")
	ErrStartTracer      = errors.New("start tracer")
	ErrShutdownTracer   = errors.New("shutdown tracer")
)

var Bindings = wire.NewSet(
	NewModule,

	profiler.NewConfig,
	profiler.NewProfiler,

	logger.NewLogger,

	metrics.NewRegistry,

	observabilitytracer.NewConfig,
	observabilitytracer.NewTracer,
	observabilitytracer.NewSDKTracerProvider,
	observabilitytracer.NewErrorHandler,
	wire.Bind(new(oteltracer.TracerProvider), new(*sdktrace.TracerProvider)),

	routers.NewHealthRouter,
	routers.NewObservabilityRouter,
	routers.NewPprofRouter,

	middlewares.NewObservabilityMiddleware,
)

type Module struct {
	profiler            *profiler.Profiler
	healthRouter        *routers.HealthRouter
	pprofRouter         *routers.PprofRouter
	observabilityRouter *routers.ObservabilityRouter
	tracerConfig        observabilitytracer.Config
	appConfig           app.Config
	sdkTraceProvider    *sdktrace.TracerProvider
	otelTracer          oteltracer.Tracer
	tracerErrorHandler  *observabilitytracer.ErrorHandler
}

func NewModule(appConfig app.Config, healthRouter *routers.HealthRouter, observabilityRouter *routers.ObservabilityRouter, pprofRouter *routers.PprofRouter, profiler *profiler.Profiler, traceProvider *sdktrace.TracerProvider, tracerConfig observabilitytracer.Config, tracerErrorHandler *observabilitytracer.ErrorHandler) *Module {
	return &Module{appConfig: appConfig, healthRouter: healthRouter, observabilityRouter: observabilityRouter, pprofRouter: pprofRouter, profiler: profiler, sdkTraceProvider: traceProvider, tracerConfig: tracerConfig, tracerErrorHandler: tracerErrorHandler}
}

func (m *Module) Name() string {
	return "observability"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	err := m.profiler.Start()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrStartProfiler, err)
	}

	m.healthRouter.Register()
	m.pprofRouter.Register()
	m.observabilityRouter.Register()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return observabilitytracer.ConfigureTracerProvider(ctx, m.tracerConfig, m.appConfig, m.tracerErrorHandler)
}

func (m *Module) Shutdown(ctx context.Context) error {
	err := m.profiler.Shutdown()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownProfiler, err)
	}

	err = observabilitytracer.Shutdown(ctx, m.sdkTraceProvider)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownTracer, err)
	}

	return nil
}
