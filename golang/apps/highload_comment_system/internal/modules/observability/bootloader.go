package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
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

var Module = fx.Module("observability",
	fx.Provide(
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

		middlewares.NewObservabilityMiddleware,
	),
	fx.Invoke(func(
		profiler *profiler.Profiler,
		healthRouter *routers.HealthRouter,
		pprofRouter *routers.PprofRouter,
		observabilityRouter *routers.ObservabilityRouter,
	) error {
		err := profiler.Start()
		if err != nil {
			return fmt.Errorf("%w: %w", ErrStartProfiler, err)
		}
		// TODO: инкапсулировать все компоненты обсервабилити
		// err := b.traceProvider.Start()
		// if err != nil {
		//	return fmt.Errorf("%w: %w", ErrStartTracer, err)
		//}

		healthRouter.Register()
		pprofRouter.Register()
		observabilityRouter.Register()

		return nil
	}),
	fx.Invoke(func(
		lc fx.Lifecycle,
		tracerConfig observabilitytracer.Config,
		appConfig app.Config,
		profiler *profiler.Profiler,
		traceProvider *sdktrace.TracerProvider,
	) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return observabilitytracer.ConfigureTracerProvider(ctx, tracerConfig, appConfig)
			},
			OnStop: func(ctx context.Context) error {
				err := profiler.Shutdown()
				if err != nil {
					return fmt.Errorf("%w: %w", ErrShutdownProfiler, err)
				}

				err = observabilitytracer.Shutdown(ctx, traceProvider)
				if err != nil {
					return fmt.Errorf("%w: %w", ErrShutdownTracer, err)
				}

				return nil
			},
		})
	}),
)
