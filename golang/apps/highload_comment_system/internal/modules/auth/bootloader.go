package auth

import (
	"context"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal/otp"
	"go.uber.org/fx"
)

type Bootloader struct {
	logger *slog.Logger
	router *http.Router
}

func NewBootloader(
	logger *slog.Logger,
	router *http.Router,
) *Bootloader {
	return &Bootloader{
		logger: logger,
		router: router,
	}
}
func (b *Bootloader) Name() string {
	return "auth"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(http.NewRouter),
		fx.Provide(dal.NewUserRepository),
		fx.Provide(otp.NewOTPRedisRepository),
		fx.Provide(otp.NewOTPPostgresqlRepository),
	}
}

func (b *Bootloader) BeforeStart() error {
	b.router.Register()
	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}
