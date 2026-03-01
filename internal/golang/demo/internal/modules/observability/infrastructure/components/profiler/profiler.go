package profiler

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/grafana/pyroscope-go"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/config"
)

type Profiler struct {
	cfgApp           config.App
	cfgObservability config.Observability
	logger           *slog.Logger
	profiler         *pyroscope.Profiler
}

func NewProfiler(
	cfgApp config.App,
	cfgObservability config.Observability,
	logger *slog.Logger,
) *Profiler {
	return &Profiler{
		cfgApp:           cfgApp,
		cfgObservability: cfgObservability,
		logger:           logger,
	}
}

var ErrPyroscopeStart = errors.New("failed to start pyroscope")

func (p *Profiler) Start() error {
	cfgPyroscope := p.cfg()

	profiler, err := pyroscope.Start(cfgPyroscope)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPyroscopeStart, err)
	}

	p.profiler = profiler

	p.logger.Info("Pyroscope profiler started",
		"url", p.cfgObservability.PyroscopeURL,
		"app", p.cfgApp.Name,
	)

	return nil
}

func (p *Profiler) cfg() pyroscope.Config {
	cfgPyroscope := pyroscope.Config{
		ApplicationName: p.cfgApp.Name,
		ServerAddress:   p.cfgObservability.PyroscopeURL,
		Logger:          newPyroScopeLogger(p.logger),
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
		UploadRate: 10 * time.Second,
	}

	return cfgPyroscope
}

func (p *Profiler) Shutdown() error {
	if p.profiler != nil {
		p.logger.Info("Shutting down pyroscope profiler")

		return p.profiler.Stop()
	}

	return nil
}
