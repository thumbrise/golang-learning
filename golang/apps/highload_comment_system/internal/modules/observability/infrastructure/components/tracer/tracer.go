package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/config"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct{}

// NewTracer creates tracer
func NewTracer(cfg config.App, provider trace.TracerProvider) trace.Tracer { //nolint:ireturn // otel returns interface, cant fix that
	return provider.Tracer(cfg.Name)
}
