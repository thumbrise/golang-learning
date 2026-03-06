package tracer

import sdktrace "go.opentelemetry.io/otel/sdk/trace"

func NewOTELSampler() sdktrace.Sampler {
	return sdktrace.ParentBased(
		sdktrace.TraceIDRatioBased(1.0),
	)
}
