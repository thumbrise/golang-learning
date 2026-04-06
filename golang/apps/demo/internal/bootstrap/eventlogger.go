package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
)

type EventLogger struct {
	logger *slog.Logger
}

func NewEventLogger(logger *slog.Logger) *EventLogger {
	return &EventLogger{logger: logger}
}

func (el *EventLogger) Log(ctx context.Context, kind, name, event string, err error) {
	msg := fmt.Sprintf("%s %s: event %s", kind, name, event)
	if err != nil {
		msg = fmt.Sprintf("%s ERROR: %s", msg, err)
		el.logger.ErrorContext(ctx, msg)
	} else {
		el.logger.DebugContext(ctx, msg)
	}
}
