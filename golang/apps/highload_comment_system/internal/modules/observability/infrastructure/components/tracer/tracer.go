package tracer

import (
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer creates tracer
func NewTracer(cfg app.Config, provider trace.TracerProvider) trace.Tracer {
	return otel.GetTracerProvider().Tracer(cfg.Name)
}
