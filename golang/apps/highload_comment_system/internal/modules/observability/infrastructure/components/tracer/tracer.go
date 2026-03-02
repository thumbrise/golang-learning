package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer creates tracer
//
//nolint:ireturn // specific
func NewTracer(cfg app.Config) trace.Tracer {
	return otel.GetTracerProvider().Tracer(cfg.Name)
}
