package components

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components/profiler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmeter "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	ErrStartProfiler    = errors.New("start profiler")
	ErrShutdownProfiler = errors.New("shutdown profiler")
	ErrShutdownLogger   = errors.New("shutdown otel logger")
	ErrShutdownMeter    = errors.New("shutdown otel meter")
	ErrShutdownTracer   = errors.New("shutdown otel tracer")
)

type Registrar struct {
	cfg               app.Config
	sdkMeterProvider  *sdkmeter.MeterProvider
	sdkTracerProvider *sdktrace.TracerProvider
	sdkLoggerProvider *sdklog.LoggerProvider
	profiler          *profiler.Profiler
	errHandler        *ErrorHandler
}

func NewRegistrar(cfg app.Config, errHandler *ErrorHandler, profiler *profiler.Profiler, sdkLoggerProvider *sdklog.LoggerProvider, sdkMeterProvider *sdkmeter.MeterProvider, sdkTracerProvider *sdktrace.TracerProvider) *Registrar {
	return &Registrar{cfg: cfg, errHandler: errHandler, profiler: profiler, sdkLoggerProvider: sdkLoggerProvider, sdkMeterProvider: sdkMeterProvider, sdkTracerProvider: sdkTracerProvider}
}

func (t *Registrar) Configure(ctx context.Context) error {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	otel.SetTracerProvider(t.sdkTracerProvider)
	otel.SetMeterProvider(t.sdkMeterProvider)
	otel.SetErrorHandler(t.errHandler)
	// Experimental
	global.SetLoggerProvider(t.sdkLoggerProvider)

	err := t.profiler.Start(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrStartProfiler, err)
	}

	return nil
}

func (t *Registrar) Shutdown(ctx context.Context) error {
	var errs error

	if err := t.profiler.Shutdown(ctx); err != nil {
		errs = errors.Join(
			errs,
			fmt.Errorf("%w: %w", ErrShutdownProfiler, err),
		)
	}

	if err := t.sdkTracerProvider.Shutdown(ctx); err != nil {
		errs = errors.Join(
			errs,
			fmt.Errorf("%w: %w", ErrShutdownTracer, err),
		)
	}

	if err := t.sdkMeterProvider.Shutdown(ctx); err != nil {
		errs = errors.Join(
			errs,
			fmt.Errorf("%w: %w", ErrShutdownMeter, err),
		)
	}

	if err := t.sdkLoggerProvider.Shutdown(ctx); err != nil {
		errs = errors.Join(
			errs,
			fmt.Errorf("%w: %w", ErrShutdownLogger, err),
		)
	}

	return errs
}
