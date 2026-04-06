package observability

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

type OtelRecorder struct {
	httpMetrics *HTTPMetrics
}

func NewOtelRecorder(httpMetrics *HTTPMetrics) *OtelRecorder {
	return &OtelRecorder{httpMetrics: httpMetrics}
}

func (o *OtelRecorder) Start(_ context.Context, c *gin.Context, span trace.Span) {
	span.SetAttributes(
		semconv.HTTPMethod(c.Request.Method),
		semconv.HTTPTarget(c.Request.URL.Path),
		semconv.URLScheme(c.Request.URL.Scheme),
		semconv.UserAgentOriginal(c.Request.UserAgent()),
		semconv.HTTPRequestBodySize(int(c.Request.ContentLength)),
		attribute.String("http.handler", c.HandlerName()),
		attribute.String("http.host", c.Request.Host),
	)
}

func (o *OtelRecorder) End(ctx context.Context, c *gin.Context, span trace.Span, duration time.Duration) {
	status := c.Writer.Status()
	respSize := c.Writer.Size()

	attrsMetrics := attribute.NewSet(
		semconv.HTTPMethod(c.Request.Method),
		semconv.HTTPRoute(c.Request.Pattern),
		semconv.HTTPStatusCode(status),
		attribute.String("http.handler", c.HandlerName()),
	)

	o.httpMetrics.AddRequest(ctx, attrsMetrics)
	o.httpMetrics.RecordDuration(ctx, duration, attrsMetrics)
	o.httpMetrics.RecordRequestSize(ctx, c.Request.ContentLength, attrsMetrics)
	o.httpMetrics.RecordResponseSize(ctx, int64(respSize), attrsMetrics)

	span.SetAttributes(
		semconv.HTTPResponseStatusCode(status),
		semconv.HTTPResponseBodySize(respSize),
		attribute.Float64(semconv.HTTPServerRequestDurationName, duration.Seconds()),
	)

	switch {
	case status >= 200 && status < 400:
		span.SetStatus(codes.Ok, "OK")
	case status >= 400 && status < 500:
		span.RecordError(c.Err())
		span.SetStatus(codes.Error, "CLIENT_ERROR")
	case status >= 500:
		span.RecordError(c.Err())
		span.SetStatus(codes.Error, "SERVER_ERROR")
	}
}
