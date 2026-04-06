package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewLogger(cfg app.Config, provider *sdklog.LoggerProvider, res *resource.Resource) *slog.Logger {
	lvl := slog.LevelInfo
	if cfg.Debug {
		lvl = slog.LevelDebug
	}

	handlerOptions := &slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handlerOptions)

	otelHandler := otelslog.NewHandler(cfg.Name,
		otelslog.WithLoggerProvider(provider),
	)

	handler := slogmulti.Fanout(
		jsonHandler,
		otelHandler,
	).WithAttrs(resToAttributes(res))

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}

func resToAttributes(res *resource.Resource) []slog.Attr {
	result := make([]slog.Attr, 0, res.Len())

	for iter := res.Iter(); iter.Next(); {
		k := string(iter.Attribute().Key)
		v := iter.Attribute().Value.AsString()
		result = append(result, slog.Any(k, v))
	}

	return result
}
