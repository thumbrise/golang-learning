package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/logger"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/meter"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/profiler"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/tracer"
)

var (
	ErrConfigureRegistrar = errors.New("failed configure registrar")
	Bindings              = wire.NewSet(
		NewModule,

		components.NewOTLPConfig,

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
		meter.NewExporter,

		tracer.NewProvider,
		tracer.NewOTELSampler,
		tracer.NewOTELExporter,
		tracer.NewOTELSDKProvider,
		tracer.NewStdOutExporter,

		NewHTTPMetrics,
	)
)

type Module struct {
	tracerErrorHandler *components.ErrorHandler
	registrar          *components.Registrar
}

func NewModule(registrar *components.Registrar, tracerErrorHandler *components.ErrorHandler) *Module {
	return &Module{registrar: registrar, tracerErrorHandler: tracerErrorHandler}
}

func (m *Module) Name() string {
	return "observability"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	err := m.registrar.Configure(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfigureRegistrar, err)
	}

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return m.registrar.Shutdown(ctx)
}
