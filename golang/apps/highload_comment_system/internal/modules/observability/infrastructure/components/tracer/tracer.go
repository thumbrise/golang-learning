package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer creates tracer
//
//nolint:ireturn // spec
func NewTracer(cfg app.Config) trace.Tracer {
	return otel.GetTracerProvider().Tracer(cfg.Name)
}
