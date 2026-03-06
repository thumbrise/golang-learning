package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewLogger(cfg app.Config, provider *sdklog.LoggerProvider) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handlerOptions)

	otelHandler := otelslog.NewHandler(cfg.Name,
		otelslog.WithLoggerProvider(provider),
	)

	handler := slogmulti.Fanout(
		jsonHandler,
		otelHandler,
	)

	logger := slog.New(handler).With(
		slog.String(string(semconv.ServiceNameKey), cfg.Name),
		slog.String(string(semconv.ServiceVersionKey), cfg.Version),
		slog.String(string(semconv.DeploymentEnvironmentKey), cfg.Environment),
	)
	slog.SetDefault(logger)

	return logger
}
