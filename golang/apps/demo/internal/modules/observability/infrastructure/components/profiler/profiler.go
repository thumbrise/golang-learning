package profiler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/grafana/pyroscope-go"
	"github.com/thumbrise/demo/golang-demo/internal/app"
)

type Profiler struct {
	cfgApp           app.Config
	cfgObservability Config
	logger           *slog.Logger
	profiler         *pyroscope.Profiler
}

func NewProfiler(
	cfgApp app.Config,
	cfgObservability Config,
	logger *slog.Logger,
) *Profiler {
	return &Profiler{
		cfgApp:           cfgApp,
		cfgObservability: cfgObservability,
		logger:           logger,
	}
}

var ErrPyroscopeStart = errors.New("failed to start pyroscope")

func (p *Profiler) Start(ctx context.Context) error {
	cfgPyroscope := pyroscopeConfig(p.cfgApp.Name, p.cfgObservability.PyroscopeURL, p.logger)

	profiler, err := pyroscope.Start(cfgPyroscope)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPyroscopeStart, err)
	}

	p.profiler = profiler

	p.logger.InfoContext(ctx, "Pyroscope profiler started",
		"url", p.cfgObservability.PyroscopeURL,
	)

	return nil
}

func (p *Profiler) Shutdown(ctx context.Context) error {
	if p.profiler != nil {
		p.logger.InfoContext(ctx, "Shutting down pyroscope profiler")

		return p.profiler.Stop()
	}

	return nil
}
