package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/logger"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	observabilitytracer "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	oteltracer "go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	ErrStartProfiler    = errors.New("start profiler")
	ErrShutdownProfiler = errors.New("shutdown profiler")
	ErrStartTracer      = errors.New("start tracer")
	ErrShutdownTracer   = errors.New("shutdown tracer")
)

type Bootloader struct {
	healthRouter        *routers.HealthRouter
	observabilityRouter *routers.ObservabilityRouter
	pprofRouter         *routers.PprofRouter
	profiler            *profiler.Profiler
	traceProvider       *sdktrace.TracerProvider
	tracerConfig        observabilitytracer.Config
	appConfig           app.Config
}

func NewBootloader(appConfig app.Config, healthRouter *routers.HealthRouter, observabilityRouter *routers.ObservabilityRouter, pprofRouter *routers.PprofRouter, profiler *profiler.Profiler, traceProvider *sdktrace.TracerProvider, tracerConfig observabilitytracer.Config) *Bootloader {
	return &Bootloader{appConfig: appConfig, healthRouter: healthRouter, observabilityRouter: observabilityRouter, pprofRouter: pprofRouter, profiler: profiler, traceProvider: traceProvider, tracerConfig: tracerConfig}
}

func (b *Bootloader) Name() string {
	return "observability"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(
			NewBootloader,

			profiler.NewConfig,
			profiler.NewProfiler,

			logger.NewLogger,

			observabilitytracer.NewConfig,
			observabilitytracer.NewTracer,
			sdktrace.NewTracerProvider,
			fx.Annotate(
				func() *sdktrace.TracerProvider {
					return &sdktrace.TracerProvider{}
				},
				fx.As(new(oteltracer.TracerProvider)),
			),

			routers.NewHealthRouter,
			routers.NewObservabilityRouter,
			routers.NewPprofRouter,
		),
	}
}

func (b *Bootloader) BeforeStart() error {
	err := b.profiler.Start()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrStartProfiler, err)
	}
	// TODO: инкапсулировать все компоненты обсервабилити
	// err := b.traceProvider.Start()
	// if err != nil {
	//	return fmt.Errorf("%w: %w", ErrStartTracer, err)
	//}

	b.healthRouter.Register()
	b.pprofRouter.Register()
	b.observabilityRouter.Register()

	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return observabilitytracer.ConfigureTracerProvider(ctx, b.tracerConfig, b.appConfig)
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	err := b.profiler.Shutdown()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownProfiler, err)
	}

	err = observabilitytracer.Shutdown(ctx, b.traceProvider)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownTracer, err)
	}

	return nil
}
