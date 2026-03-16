package otel

import "go.opentelemetry.io/otel"

const name = "http"

var (
	Tracer = otel.Tracer(name)
	Meter  = otel.Meter(name)
)
