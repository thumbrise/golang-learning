package observability

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

type (
	profilerContextKey string

	OTELMiddleware struct {
		cfgApp         app.Config
		profiler       *profiler.Profiler
		logger         *slog.Logger
		tracerProvider *tracer.Provider
		httpMetrics    *observability.HTTPMetrics
	}
)

const profilerContextValue profilerContextKey = "profiler"

func NewOTELMiddleware(cfgApp app.Config, httpMetrics *observability.HTTPMetrics, logger *slog.Logger, profiler *profiler.Profiler, tracerProvider *tracer.Provider) *OTELMiddleware {
	return &OTELMiddleware{cfgApp: cfgApp, httpMetrics: httpMetrics, logger: logger, profiler: profiler, tracerProvider: tracerProvider}
}

func (m *OTELMiddleware) Handler(_ context.Context) gin.HandlerFunc {
	trc := m.tracerProvider.Tracer()

	return func(c *gin.Context) {
		start := time.Now()

		// 1. Получаем путь обработчика (если доступен)
		handler := "unknown"
		if c.HandlerName() != "" {
			handler = c.HandlerName()
		}

		// 2. Измеряем размер запроса
		reqSize := c.Request.ContentLength
		if reqSize < 0 {
			reqSize = 0
		}
		// 3. Начинаем трейс (спан) для этого запроса
		ctx, span := trc.Start(
			c.Request.Context(),
			fmt.Sprintf("%s %s", c.Request.Method, c.FullPath()),
			trace.WithAttributes(
				semconv.HTTPMethod(c.Request.Method),
				semconv.HTTPTarget(c.FullPath()),
				semconv.UserAgentOriginal(c.Request.UserAgent()),
				semconv.URLScheme(c.Request.URL.Scheme),
				semconv.HTTPRequestBodySize(int(reqSize)),
				attribute.String("http.handler", handler),
				attribute.String("http.host", c.Request.Host),
			),
		)

		if m.profiler != nil {
			labels := map[string]string{
				string(semconv.HTTPTargetKey): c.Request.URL.Path,
				string(semconv.HTTPMethodKey): c.Request.Method,
				"handler":                     handler,
				"span_id":                     span.SpanContext().SpanID().String(),
				"trace_id":                    span.SpanContext().TraceID().String(),
			}

			ctx = context.WithValue(ctx, profilerContextValue, labels)
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		duration := time.Since(start)

		status := c.Writer.Status()
		respSize := c.Writer.Size()

		m.httpMetrics.AddRequest(ctx,
			semconv.HTTPMethod(c.Request.Method),
			semconv.URLPath(c.Request.URL.Path),
			semconv.HTTPStatusCode(status),
			attribute.String("http.handler", handler),
		)

		m.httpMetrics.RecordDuration(ctx,
			duration,
			semconv.HTTPMethod(c.Request.Method),
			semconv.URLPath(c.Request.URL.Path),
			attribute.String("http.handler", handler),
		)
		m.httpMetrics.RecordRequestSize(ctx,
			float64(reqSize),
			semconv.HTTPMethod(c.Request.Method),
			semconv.URLPath(c.Request.URL.Path),
		)

		m.httpMetrics.RecordResponseSize(ctx,
			float64(respSize),
			semconv.HTTPMethod(c.Request.Method),
			semconv.URLPath(c.Request.URL.Path),
		)

		span.SetAttributes(
			semconv.HTTPResponseStatusCode(status),
			semconv.HTTPResponseBodySize(respSize),
			attribute.Float64("http.duration_seconds", duration.Seconds()),
		)

		if status >= 400 {
			span.SetAttributes(
				attribute.Bool("error", true),
				attribute.String("error.type", "http_error"),
			)

			if status >= 500 {
				span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", status))
			}
		}

		span.End()

		// 13. Логируем медленные запросы
		m.logSlow(ctx,
			duration,
			string(semconv.HTTPMethodKey), c.Request.Method,
			string(semconv.URLPathKey), c.Request.URL.Path,
			"duration_ms", duration.Milliseconds(),
			"status", status,
			"handler", handler,
		)
	}
}

func (m *OTELMiddleware) logSlow(ctx context.Context, duration time.Duration, args ...any) {
	if duration.Seconds() < 1.0 {
		return
	}

	m.logger.WarnContext(ctx, "Slow request",
		args...,
	)
}
