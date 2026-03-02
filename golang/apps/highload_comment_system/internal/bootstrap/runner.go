package bootstrap

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

type Runner struct {
	logger *EventLogger
}

func NewRunner(
	logger *EventLogger,
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

func (h *Runner) Run(ctx context.Context, processes []*Process) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	grp, ctx := errgroup.WithContext(ctx)
	h.startProcesses(ctx, processes, grp)

	grp.Go(func() error {
		h.logEvent("Process", "Gracefull Shutdown", "waiting for signal", nil)
		<-ctx.Done()
		h.logEvent("Process", "Gracefull Shutdown", "received signal to exit", nil)

		h.shutdownProcesses(ctx, processes)

		return nil
	})

	h.logEvent("Process", "ErrorGroup", "Wait", grp.Wait())

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

func (h *Runner) logEvent(kind, name, event string, err error) {
	h.logger.Log(kind, name, event, err)
}
