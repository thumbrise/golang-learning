package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"golang.org/x/sync/errgroup"
)

type Runner struct {
	logger *slog.Logger
}

func NewRunner(
	logger *slog.Logger,
) *Runner {
	return &Runner{
		logger: logger,
	}
}

type (
	StartFunc    func(ctx context.Context) error
	ShutdownFunc func(ctx context.Context) error
	Process      struct {
		Name     string
		Start    StartFunc
		Shutdown ShutdownFunc
	}
)

func (h *Runner) Run(ctx context.Context, processes []*Process, modules []contracts.Module) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	h.bootstrapModules(ctx, modules)

	grp, ctx := errgroup.WithContext(ctx)
	h.startProcesses(ctx, processes, grp)

	grp.Go(func() error {
		h.logger.Info("waiting for signal")
		<-ctx.Done()
		h.logger.Info("received signal to exit")

		h.shutdownProcesses(ctx, processes)

		return nil
	})

	h.logEvent("Process", "ErrorGroup", "Wait", grp.Wait())

	h.shutdownModules(ctx, modules)

	return nil
}

func (h *Runner) startProcesses(ctx context.Context, processes []*Process, grp *errgroup.Group) {
	for _, pp := range processes {
		p := pp

		h.logEvent("Process", p.Name, "start", nil)

		grp.Go(func() error {
			err := p.Start(ctx)
			if err != nil {
				h.logEvent("Process", p.Name, "long run", err)
			}

			return err
		})
	}
}

func (h *Runner) shutdownProcesses(ctx context.Context, processes []*Process) {
	for _, pp := range processes {
		p := pp

		err := p.Shutdown(ctx)
		h.logEvent("Process", p.Name, "shutdown", err)
	}
}

func (h *Runner) bootstrapModules(ctx context.Context, modules []contracts.Module) {
	for _, mm := range modules {
		m := mm
		h.logEvent("Module", m.Name(), "got", nil)
	}

	for _, mm := range modules {
		m := mm

		err := m.BeforeStart(ctx)
		h.logEvent("Module", m.Name(), "before start", err)
	}

	for _, mm := range modules {
		m := mm

		err := m.OnStart(ctx)
		h.logEvent("Module", m.Name(), "on start", err)
	}
}

func (h *Runner) shutdownModules(ctx context.Context, modules []contracts.Module) {
	for _, mm := range modules {
		m := mm

		err := m.Shutdown(ctx)
		h.logEvent("Module", m.Name(), "shutdown", err)
	}
}

func (h *Runner) logEvent(kind, name, event string, err error) {
	msg := fmt.Sprintf("%s %s: event %s", kind, name, event)
	if err != nil {
		msg = fmt.Sprintf("%s ERROR: %s", msg, err)
		h.logger.Error(msg)
	} else {
		h.logger.Info(msg)
	}
}
