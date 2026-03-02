package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"golang.org/x/sync/errgroup"
)

type Runner struct {
	logger  *slog.Logger
	modules []contracts.Module
}

func NewRunner(
	logger *slog.Logger,
	modules []contracts.Module,
) *Runner {
	return &Runner{
		logger:  logger,
		modules: modules,
	}
}

type StartFunc func(ctx context.Context) error
type ShutdownFunc func(ctx context.Context) error
type Process struct {
	Name     string
	Start    StartFunc
	Shutdown ShutdownFunc
}

func (h *Runner) Run(ctx context.Context, processes []*Process) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	h.bootstrapModules(ctx)

	grp, ctx := errgroup.WithContext(ctx)

	for _, pp := range processes {
		p := pp

		h.logEvent("Process", p.Name, "long run", nil)

		grp.Go(func() error {
			err := p.Start(ctx)
			if err != nil {
				h.logEvent("Process", p.Name, "long run", err)
			}

			return err
		})
	}

	grp.Go(func() error {
		h.logger.Info("waiting for signal")
		<-ctx.Done()
		h.logger.Info("received signal to exit")
		for _, pp := range processes {
			p := pp

			h.logEvent("Process", p.Name, "shutdown", nil)
		}

		ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		h.shutdownModules(ctxShutdown)

		return nil
	})

	return grp.Wait()
}

func (h *Runner) bootstrapModules(ctx context.Context) {
	for _, mm := range h.modules {
		m := mm
		h.logEvent("Module", m.Name(), "got", nil)
	}

	for _, mm := range h.modules {
		m := mm

		err := m.BeforeStart(ctx)
		h.logEvent("Module", m.Name(), "before start", err)
	}

	for _, mm := range h.modules {
		m := mm

		err := m.OnStart(ctx)
		h.logEvent("Module", m.Name(), "on start", err)
	}
}

func (h *Runner) shutdownModules(ctx context.Context) {
	for _, mm := range h.modules {
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
