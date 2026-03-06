package bootstrap

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Bootstrapper struct {
	logger *EventLogger
}

func NewBootstrapper(logger *EventLogger) *Bootstrapper {
	return &Bootstrapper{logger: logger}
}

func (b *Bootstrapper) Bootstrap(ctx context.Context, modules []contracts.Module) error {
	for _, mm := range modules {
		m := mm
		b.logger.Log("Module", m.Name(), "got", nil)
	}

	for _, mm := range modules {
		m := mm

		err := m.BeforeStart(ctx)
		b.logger.Log("Module", m.Name(), "before start", err)
	}

	for _, mm := range modules {
		m := mm

		err := m.OnStart(ctx)
		b.logger.Log("Module", m.Name(), "on start", err)
	}

	return nil
}

func (b *Bootstrapper) Shutdown(ctx context.Context, modules []contracts.Module) error {
	for _, mm := range modules {
		m := mm

		err := m.Shutdown(ctx)
		b.logger.Log("Module", m.Name(), "shutdown", err)
	}

	return nil
}
