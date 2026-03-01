package tracer

import (
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/config"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct{}

// NewTracer creates tracer
func NewTracer(cfg config.App, provider trace.TracerProvider) trace.Tracer { //nolint:ireturn // otel returns interface, cant fix that
	return provider.Tracer(cfg.Name)
}
