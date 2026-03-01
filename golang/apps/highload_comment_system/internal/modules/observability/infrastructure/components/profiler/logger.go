package profiler

import (
	"fmt"
	"log/slog"
)

// Logger для Pyroscope
type pyroscopeLogger struct {
	logger *slog.Logger
}

func (p *pyroscopeLogger) Infof(format string, args ...interface{}) {
	p.logger.Info(fmt.Sprintf(format, args...))
}

func (p *pyroscopeLogger) Debugf(format string, args ...interface{}) {
	p.logger.Debug(fmt.Sprintf(format, args...))
}

func (p *pyroscopeLogger) Errorf(format string, args ...interface{}) {
	p.logger.Error(fmt.Sprintf(format, args...))
}

func newPyroScopeLogger(logger *slog.Logger) *pyroscopeLogger {
	return &pyroscopeLogger{logger: logger}
}
