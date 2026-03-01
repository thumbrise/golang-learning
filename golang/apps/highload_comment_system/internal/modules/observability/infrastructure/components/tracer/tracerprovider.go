package tracer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var ErrTraceExporter = errors.New("failed to create trace exporter")

// NewTracerProvider creates exporter struct
func NewTracerProvider(ctx context.Context, cfgTrace config.Observability, cfgApp config.App) (*sdktrace.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(cfgTrace.OTLPURL),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithURLPath("/v1/traces"),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTraceExporter, err)
	}

	// Создаем ресурс с атрибутами сервиса
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfgApp.Name),
			semconv.ServiceVersion(cfgApp.Version),
			semconv.DeploymentEnvironment(cfgApp.Environment),
		),
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithContainer(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Создаем TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp,
			sdktrace.WithMaxExportBatchSize(512),
			sdktrace.WithMaxQueueSize(2048),
			sdktrace.WithExportTimeout(30*time.Second),
		),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(
			sdktrace.TraceIDRatioBased(1.0),
		)),
	)

	// Настраиваем глобальный провайдер
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	otel.SetTracerProvider(tp)
	otel.SetErrorHandler(ErrorHandler{})

	return tp, nil
}

// Shutdown правильно завершает работу TracerProvider
func Shutdown(ctx context.Context, tp *sdktrace.TracerProvider) error {
	if tp == nil {
		return nil
	}
	return tp.Shutdown(ctx)
}

// ShutdownWithTimeout удобная обертка с таймаутом
func ShutdownWithTimeout(tp *sdktrace.TracerProvider, timeout time.Duration) error {
	if tp == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return tp.Shutdown(ctx)
}
