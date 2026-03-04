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
		var err error = nil
		h.logger.Log("Process", "Graceful Shutdown", "waiting for signal", err)
		<-ctx.Done()

		var err2 error = nil
		h.logger.Log("Process", "Graceful Shutdown", "received signal to exit", err2)

		h.shutdownProcesses(ctx, processes)

		return nil
	})

	h.logger.Log("Process", "ErrorGroup", "Wait", grp.Wait())

	return nil
}

func (h *Runner) startProcesses(ctx context.Context, processes []*Process, grp *errgroup.Group) {
	for _, pp := range processes {
		p := pp

		var err error = nil
		h.logger.Log("Process", p.Name, "start", err)

		grp.Go(func() error {
			err := p.Start(ctx)
			if err != nil {
				h.logger.Log("Process", p.Name, "long run", err)
			}

			return err
		})
	}
}

func (h *Runner) shutdownProcesses(ctx context.Context, processes []*Process) {
	for _, pp := range processes {
		p := pp
		if p.Shutdown == nil {
			return
		}

		err := p.Shutdown(ctx)
		h.logger.Log("Process", p.Name, "shutdown", err)
	}
}
