package profiler

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/grafana/pyroscope-go"
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

func pyroscopeConfig(app string, addr string, logger *slog.Logger) pyroscope.Config {
	return pyroscope.Config{
		ApplicationName: app,
		ServerAddress:   addr,
		Logger:          &pyroscopeLogger{logger: logger},
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
}
