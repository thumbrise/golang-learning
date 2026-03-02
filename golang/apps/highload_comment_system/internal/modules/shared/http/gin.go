package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewGinEngine(logger *slog.Logger) *gin.Engine {
	engine := gin.New()
	engine.Use(sloggin.New(logger))
	engine.Use(gin.Recovery())

	return engine
}
