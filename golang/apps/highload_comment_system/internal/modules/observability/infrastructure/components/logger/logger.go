package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	keyTraceId           = "trace_id"
	keySpanId            = "span_id"
	keyTraceFlagsSampled = "trace_flags.sampled"
)

func NewLogger(cfg app.Config) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}

	// Оборачиваем handler для автоматического добавления trace_id
	baseHandler := slog.NewJSONHandler(os.Stdout, handlerOptions)
	handler := &traceHandler{Handler: baseHandler}

	logger := slog.New(handler).With(
		slog.String(string(semconv.ServiceNameKey), cfg.Name),
		slog.String(string(semconv.ServiceVersionKey), cfg.Version),
		slog.String(string(semconv.DeploymentEnvironmentKey), cfg.Environment),
	)

	slog.SetDefault(logger)

	return logger
}

// traceHandler автоматически добавляет trace_id и span_id
type traceHandler struct {
	slog.Handler
}

func (h *traceHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		spanCtx := span.SpanContext()
		r.AddAttrs(
			slog.String(keyTraceId, spanCtx.TraceID().String()),
			slog.String(keySpanId, spanCtx.SpanID().String()),
			slog.Bool(keyTraceFlagsSampled, spanCtx.IsSampled()),
		)
	}

	return h.Handler.Handle(ctx, r)
}

func (h *traceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &traceHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *traceHandler) WithGroup(name string) slog.Handler {
	return &traceHandler{Handler: h.Handler.WithGroup(name)}
}
