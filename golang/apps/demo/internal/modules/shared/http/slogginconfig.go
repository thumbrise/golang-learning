package http

import sloggin "github.com/samber/slog-gin"

func NewSlogginConfig() sloggin.Config {
	return sloggin.Config{
		DefaultLevel:       0,
		ClientErrorLevel:   0,
		ServerErrorLevel:   0,
		WithUserAgent:      true,
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  true,
		WithResponseBody:   true,
		WithResponseHeader: true,
		WithSpanID:         true,
		WithTraceID:        true,
		WithClientIP:       true,
		HandleGinDebug:     true,
		Filters:            nil,
	}
}
