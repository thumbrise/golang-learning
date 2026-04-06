package analytics

import (
	"context"

	"github.com/google/wire"
)

var Bindings = wire.NewSet(
	NewModule,
	NewMetrics,
)

type Module struct {
	metrics *Metrics
}

func NewModule(metrics *Metrics) *Module {
	return &Module{metrics: metrics}
}

func (m *Module) Name() string {
	return "analytics"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.metrics.GaugeUsersTotal()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
