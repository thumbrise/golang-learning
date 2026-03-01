package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Kernel struct {
	srv    *http.Server
	config Config
	logger *slog.Logger
	engine *gin.Engine
}

func NewKernel(
	config Config,
	logger *slog.Logger,
	engine *gin.Engine,
) *Kernel {
	router := &Kernel{
		config: config,
		logger: logger,
		engine: engine,
	}

	router.srv = &http.Server{
		Addr:              ":" + config.Port,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return router
}

func (k *Kernel) Server() *http.Server {
	return k.srv
}

func (k *Kernel) Gin() *gin.Engine {
	return k.engine
}

func (k *Kernel) Start(ctx context.Context) error {
	msg := "started server on http://localhost:" + k.config.Port
	k.logger.Info(msg)

	return k.srv.ListenAndServe()
}

func (k *Kernel) Shutdown(ctx context.Context) error {
	return k.srv.Shutdown(ctx)
}

func (k *Kernel) Name() string {
	return "http"
}
