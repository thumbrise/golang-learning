package logger

import (
	"context"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	keyTraceId           = "trace_id"
	keySpanId            = "span_id"
	keyTraceFlagsSampled = "trace_flags.sampled"
)

func NewLogger(cfg app.Config, attrs *CommonAttributesProvider) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handlerOptions)

	// attrs := otelslog.WithAttributes(
	//	attrs.GetAttributes()...,
	//)
	otelHandler := otelslog.NewHandler(cfg.Name)

	// logger := slog.New(handler).With(
	//	slog.String(string(semconv.ServiceNameKey), cfg.Name),
	//	slog.String(string(semconv.ServiceVersionKey), cfg.Version),
	//	slog.String(string(semconv.DeploymentEnvironmentKey), cfg.Environment),
	//)
	handler := slogmulti.Fanout(
		jsonHandler,
		otelHandler,
	)
	// handler = slogmulti.Pipe(addOtelAttributes()).Handler(handler)

	logger := slog.New(handler).With(
		slog.String(string(semconv.ServiceNameKey), cfg.Name),
		slog.String(string(semconv.ServiceVersionKey), cfg.Version),
		slog.String(string(semconv.DeploymentEnvironmentKey), cfg.Environment),
	)
	slog.SetDefault(logger)

	return logger
}

func addOtelAttributes() slogmulti.Middleware {
	return slogmulti.NewHandleInlineMiddleware(func(ctx context.Context, record slog.Record, next func(context.Context, slog.Record) error) error {
		span := trace.SpanFromContext(ctx)
		if span.SpanContext().IsValid() {
			spanCtx := span.SpanContext()
			record.AddAttrs(
				slog.String(keyTraceId, spanCtx.TraceID().String()),
				slog.String(keySpanId, spanCtx.SpanID().String()),
				slog.Bool(keyTraceFlagsSampled, spanCtx.IsSampled()),
			)
		} else {
			msg := "invalid span context"
			if record.Message != msg {
				slog.WarnContext(ctx, msg)
			}
		}

		return next(ctx, record)
	})
}
