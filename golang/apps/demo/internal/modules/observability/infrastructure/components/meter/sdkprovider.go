package meter

import (
	sdkmeter "go.opentelemetry.io/otel/sdk/metric"
)

func NewOTELSDKProvider() *sdkmeter.MeterProvider {
	return sdkmeter.NewMeterProvider()
}
