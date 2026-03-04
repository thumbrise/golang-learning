package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewGinEngine(logger *slog.Logger, slogginConfig sloggin.Config) *gin.Engine {
	engine := gin.New()

	engine.Use(sloggin.NewWithConfig(logger, slogginConfig))
	engine.Use(gin.Recovery())

	return engine
}
