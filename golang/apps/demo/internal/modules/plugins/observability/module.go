package observability

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure"
	components2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components"
	logger2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components/logger"
	meter2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components/meter"
	profiler2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components/profiler"
	tracer2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/infrastructure/components/tracer"
)

var (
	ErrConfigureRegistrar = errors.New("failed configure registrar")
	Bindings              = wire.NewSet(
		NewModule,

		infrastructure.NewOTLPConfig,

		components2.NewResource,
		components2.NewRegistrar,
		components2.NewErrorHandler,

		profiler2.NewConfig,
		profiler2.NewProfiler,

		logger2.NewLogger,
		logger2.NewOTELSDKProvider,
		logger2.NewProvider,
		logger2.NewExporter,

		meter2.NewOTELSDKProvider,
		meter2.NewProvider,
		meter2.NewExporter,

		tracer2.NewProvider,
		tracer2.NewOTELSampler,
		tracer2.NewOTELExporter,
		tracer2.NewOTELSDKProvider,
		tracer2.NewStdOutExporter,

		NewHTTPMetrics,
	)
)

type Module struct {
	tracerErrorHandler *components2.ErrorHandler
	registrar          *components2.Registrar
}

func NewModule(registrar *components2.Registrar, tracerErrorHandler *components2.ErrorHandler) *Module {
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
