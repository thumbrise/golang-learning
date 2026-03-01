package components

import (
	"context"
	"log/slog"
	"os"

	"github.com/thumbrise/demo/golang-demo/internal/config"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger(cfg config.App) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}

	// Оборачиваем handler для автоматического добавления trace_id
	baseHandler := slog.NewJSONHandler(os.Stdout, handlerOptions)
	handler := &traceHandler{Handler: baseHandler}

	logger := slog.New(handler).With(
		slog.String("service.name", cfg.Name),
		slog.String("service.version", cfg.Version),
		slog.String("environment", cfg.Environment),
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
			slog.String("trace_id", spanCtx.TraceID().String()),
			slog.String("span_id", spanCtx.SpanID().String()),
			slog.Bool("trace_flags.sampled", spanCtx.IsSampled()),
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
