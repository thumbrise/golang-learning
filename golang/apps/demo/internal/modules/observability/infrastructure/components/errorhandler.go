package components

import (
	"context"
	"log/slog"
)

type ErrorHandler struct {
	logger *slog.Logger
}

func NewErrorHandler(logger *slog.Logger) *ErrorHandler {
	return &ErrorHandler{logger: logger}
}

func (l *ErrorHandler) Handle(err error) {
	l.logger.ErrorContext(context.Background(),
		"openTelemetry error",
		slog.String("err", err.Error()),
	)
}
