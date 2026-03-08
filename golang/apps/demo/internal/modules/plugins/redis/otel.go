package redis

import (
	"errors"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	sdkmeter "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	ErrInstrumentTracing = errors.New("instrument tracing")
	ErrInstrumentMetrics = errors.New("instrument metrics")
)

type OTELRegistrar struct {
	client         *redis.Client
	tracerProvider *sdktrace.TracerProvider
	meterProvider  *sdkmeter.MeterProvider
}

func NewOTELRegistrar(mtp *sdkmeter.MeterProvider, rdb *redis.Client, trp *sdktrace.TracerProvider) *OTELRegistrar {
	return &OTELRegistrar{meterProvider: mtp, client: rdb, tracerProvider: trp}
}

func (or *OTELRegistrar) Register() error {
	err := redisotel.InstrumentTracing(or.client, redisotel.WithTracerProvider(or.tracerProvider))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInstrumentTracing, err)
	}

	err = redisotel.InstrumentMetrics(or.client, redisotel.WithMeterProvider(or.meterProvider))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInstrumentMetrics, err)
	}

	return nil
}
