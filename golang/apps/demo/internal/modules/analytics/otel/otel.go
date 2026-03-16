package otel

import "go.opentelemetry.io/otel"

const name = "analytics"

var Meter = otel.Meter(name)
