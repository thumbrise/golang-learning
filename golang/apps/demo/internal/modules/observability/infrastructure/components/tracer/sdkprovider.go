package tracer

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func NewOTELSDKProvider(
	res *resource.Resource,
	exp *otlptrace.Exporter,
	sampler sdktrace.Sampler,
	expstdout *stdouttrace.Exporter,
) *sdktrace.TracerProvider {
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(
			exp,
			batchOptions()...,
		),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sampler),
	)
}
