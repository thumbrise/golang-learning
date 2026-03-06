package observability

import (
	"context"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/meter"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/histogram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type HTTPMetrics struct {
	requestsTotal   metric.Int64Counter
	requestDuration metric.Float64Histogram
	requestSize     metric.Float64Histogram
	responseSize    metric.Float64Histogram
}

func NewHTTPMetrics(p *meter.Provider) *HTTPMetrics {
	m := &HTTPMetrics{}
	mtr := p.Meter()

	var err error

	m.requestsTotal, err = mtr.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
	)
	if err != nil {
		otel.Handle(err)
	}

	defBuckets := histogram.DefBuckets()

	m.requestDuration, err = mtr.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("Duration of HTTP requests in seconds"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(defBuckets...),
	)
	if err != nil {
		otel.Handle(err)
	}

	// 100b to 1gb
	sizeBuckets := histogram.ExponentialBuckets(100, 10, 8)

	m.requestSize, err = mtr.Float64Histogram(
		"http_request_size_bytes",
		metric.WithDescription("HTTP request size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(sizeBuckets...),
	)
	if err != nil {
		otel.Handle(err)
	}

	m.responseSize, err = mtr.Float64Histogram(
		"http_response_size_bytes",
		metric.WithDescription("HTTP response size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(sizeBuckets...),
	)
	if err != nil {
		otel.Handle(err)
	}

	return m
}

// AddRequest увеличивает счётчик запросов с заданными атрибутами.
func (m *HTTPMetrics) AddRequest(ctx context.Context, attrs ...attribute.KeyValue) {
	if m.requestsTotal != nil {
		m.requestsTotal.Add(ctx, 1, metric.WithAttributes(attrs...))
	}
}

// RecordDuration записывает длительность запроса.
func (m *HTTPMetrics) RecordDuration(ctx context.Context, duration time.Duration, attrs ...attribute.KeyValue) {
	if m.requestDuration != nil {
		m.requestDuration.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
	}
}

// RecordRequestSize записывает размер запроса.
func (m *HTTPMetrics) RecordRequestSize(ctx context.Context, sizeBytes float64, attrs ...attribute.KeyValue) {
	if m.requestSize != nil {
		m.requestSize.Record(ctx, sizeBytes, metric.WithAttributes(attrs...))
	}
}

// RecordResponseSize записывает размер ответа.
func (m *HTTPMetrics) RecordResponseSize(ctx context.Context, sizeBytes float64, attrs ...attribute.KeyValue) {
	if m.responseSize != nil {
		m.responseSize.Record(ctx, sizeBytes, metric.WithAttributes(attrs...))
	}
}
