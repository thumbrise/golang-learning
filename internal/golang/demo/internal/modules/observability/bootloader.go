package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	"go.opentelemetry.io/otel/sdk/trace"
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
	traceProvider       *trace.TracerProvider
}

func NewBootloader(healthRouter *routers.HealthRouter, observabilityRouter *routers.ObservabilityRouter, pprofRouter *routers.PprofRouter, profiler *profiler.Profiler, traceProvider *trace.TracerProvider) *Bootloader {
	return &Bootloader{healthRouter: healthRouter, observabilityRouter: observabilityRouter, pprofRouter: pprofRouter, profiler: profiler, traceProvider: traceProvider}
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	err := b.profiler.Shutdown()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownProfiler, err)
	}

	err = b.traceProvider.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrShutdownTracer, err)
	}

	return nil
}

func (b *Bootloader) Boot(ctx context.Context) error {
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
