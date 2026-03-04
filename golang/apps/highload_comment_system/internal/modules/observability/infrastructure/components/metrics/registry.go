package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func NewRegistry() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func M() {
	// m := otel.Meter("")
}
