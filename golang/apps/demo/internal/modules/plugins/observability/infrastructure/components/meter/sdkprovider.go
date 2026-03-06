package meter

import (
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmeter "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewOTELSDKProvider(res *resource.Resource, exp *otlpmetricgrpc.Exporter) *sdkmeter.MeterProvider {
	return sdkmeter.NewMeterProvider(
		sdkmeter.WithResource(res),
		sdkmeter.WithReader(sdkmeter.NewPeriodicReader(exp,
			sdkmeter.WithInterval(10*time.Second),
		)),
	)
}
