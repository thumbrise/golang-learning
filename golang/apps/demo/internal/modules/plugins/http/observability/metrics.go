package observability

import (
	"context"
	"errors"
	"time"

	otelhttp "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/otel"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/histogram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/semconv/v1.39.0/httpconv"
)

type HTTPMetrics struct {
	requestsTotal   metric.Int64Counter
	requestDuration httpconv.ServerRequestDuration
	requestSize     httpconv.ServerRequestBodySize
	responseSize    httpconv.ServerResponseBodySize
}

func NewHTTPMetrics() *HTTPMetrics {
	m := &HTTPMetrics{}
	mtr := otelhttp.Meter

	var (
		errs error
		err  error
	)

	m.requestsTotal, err = mtr.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
	)
	errs = errors.Join(errs, err)

	m.requestDuration, err = httpconv.NewServerRequestDuration(
		mtr,
		metric.WithExplicitBucketBoundaries(histogram.DefBuckets()...),
	)
	errs = errors.Join(errs, err)

	// 100b to 1gb
	sizeBuckets := histogram.ExponentialBuckets(100, 10, 8)

	m.requestSize, err = httpconv.NewServerRequestBodySize(
		mtr,
		metric.WithExplicitBucketBoundaries(sizeBuckets...),
	)
	errs = errors.Join(errs, err)

	m.responseSize, err = httpconv.NewServerResponseBodySize(
		mtr,
		metric.WithExplicitBucketBoundaries(sizeBuckets...),
	)

	errs = errors.Join(errs, err)
	if errs != nil {
		otel.Handle(errs)
	}

	return m
}

// AddRequest увеличивает счётчик запросов с заданными атрибутами.
func (m *HTTPMetrics) AddRequest(ctx context.Context, attrs attribute.Set) {
	m.requestsTotal.Add(ctx, 1, metric.WithAttributeSet(attrs))
}

// RecordDuration записывает длительность запроса.
func (m *HTTPMetrics) RecordDuration(ctx context.Context, duration time.Duration, attrs attribute.Set) {
	m.requestDuration.RecordSet(
		ctx,
		duration.Seconds(),
		attrs,
	)
}

// RecordRequestSize записывает размер запроса.
func (m *HTTPMetrics) RecordRequestSize(ctx context.Context, sizeBytes int64, attrs attribute.Set) {
	m.requestSize.RecordSet(ctx, sizeBytes, attrs)
}

// RecordResponseSize записывает размер ответа.
func (m *HTTPMetrics) RecordResponseSize(ctx context.Context, sizeBytes int64, attrs attribute.Set) {
	m.responseSize.RecordSet(ctx, sizeBytes, attrs)
}
