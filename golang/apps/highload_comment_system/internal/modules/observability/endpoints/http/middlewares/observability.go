// internal/interfaces/http/middlewares/observability.go
package middlewares

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/config"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ObservabilityMiddleware struct {
	cfgApp           app.Config
	cfgObservability config.Observability
	pyroscopeClient  *profiler.Profiler
	logger           *slog.Logger
}

func NewObservabilityMiddleware(cfgApp app.Config, cfgObservability config.Observability, pyroscopeClient *profiler.Profiler, logger *slog.Logger) *ObservabilityMiddleware {
	return &ObservabilityMiddleware{cfgApp: cfgApp, cfgObservability: cfgObservability, pyroscopeClient: pyroscopeClient, logger: logger}
}

// Метрики Prometheus
var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	}, []string{"method", "path", "status", "handler"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "path", "handler"})

	httpRequestSize = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_size_bytes",
		Help:    "HTTP request size in bytes",
		Buckets: prometheus.ExponentialBuckets(100, 10, 8),
	}, []string{"method", "path"})

	httpResponseSize = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_response_size_bytes",
		Help:    "HTTP response size in bytes",
		Buckets: prometheus.ExponentialBuckets(100, 10, 8),
	}, []string{"method", "path"})
)

func (m *ObservabilityMiddleware) Handler() gin.HandlerFunc {
	tracer := otel.Tracer("gin-http-server")

	return func(c *gin.Context) {
		start := time.Now()

		// 1. Получаем путь обработчика (если доступен)
		handler := "unknown"
		if c.HandlerName() != "" {
			handler = c.HandlerName()
		}

		// 2. Измеряем размер запроса
		reqSize := float64(c.Request.ContentLength)
		if reqSize < 0 {
			reqSize = 0
		}

		// 3. Начинаем трейс (спан) для этого запроса
		ctx, span := tracer.Start(
			c.Request.Context(),
			fmt.Sprintf("%s %s", c.Request.Method, c.FullPath()),
			trace.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.path", c.FullPath()),
				attribute.String("http.handler", handler),
				attribute.String("http.user_agent", c.Request.UserAgent()),
				attribute.String("http.scheme", c.Request.URL.Scheme),
				attribute.String("http.host", c.Request.Host),
				attribute.Float64("http.request.size", reqSize),
			),
		)

		// 4. Добавляем теги Pyroscope в контекст (опционально)
		if m.pyroscopeClient != nil {
			// Можно добавить динамические теги для профилирования
			labels := map[string]string{
				"http_method": c.Request.Method,
				"http_path":   c.FullPath(),
				"handler":     handler,
			}

			// Создаем контекст с тегами
			ctx = context.WithValue(ctx, "pyroscope_labels", labels)
		}

		// 5. Прокидываем обновленный контекст
		c.Request = c.Request.WithContext(ctx)

		// 6. Обрабатываем запрос
		c.Next()

		// 7. Измеряем время выполнения
		duration := time.Since(start).Seconds()

		// 8. Получаем статус ответа
		status := c.Writer.Status()
		respSize := float64(c.Writer.Size())

		// 9. Записываем метрики Prometheus
		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			strconv.Itoa(status),
			handler,
		).Inc()

		httpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			handler,
		).Observe(duration)

		httpRequestSize.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(reqSize)

		httpResponseSize.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(respSize)

		// 10. Обновляем спан (трейс)
		span.SetAttributes(
			attribute.Int("http.status_code", status),
			attribute.Float64("http.duration_seconds", duration),
			attribute.Float64("http.response.size", respSize),
			attribute.Int("http.response.headers", len(c.Writer.Header())),
		)

		// 11. Помечаем ошибки в трейсе
		if status >= 400 {
			span.SetAttributes(
				attribute.Bool("error", true),
				attribute.String("error.type", "http_error"),
			)

			if status >= 500 {
				span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", status))
			}
		}

		// 12. Завершаем спан
		span.End()

		// 13. Логируем медленные запросы
		if duration > 1.0 { // Более 1 секунды
			m.logger.Warn("Slow request",
				"method", c.Request.Method,
				"path", c.FullPath(),
				"duration", duration,
				"status", status,
				"handler", handler,
			)
		}
	}
}
