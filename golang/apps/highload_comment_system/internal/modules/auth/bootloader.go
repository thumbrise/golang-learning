package auth

import (
	"context"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	authusecases "github.com/thumbrise/demo/golang-demo/internal/modules/auth/application/usecases"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	otpdal "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/jwt"
	authmailers "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	otp "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/otp"
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
		fx.Provide(
			NewBootloader,

			http.NewMiddleware,
			http.NewRouter,

			authmailers.NewOTPMail,

			authusecases.NewAuthCommandSignIn,
			authusecases.NewAuthCommandExchangeOtp,
			authusecases.NewAuthQueryMe,
			authusecases.NewAuthCommandRefresh,

			dal.NewUserRepository,

			otp.NewConfig,
			otpdal.NewOTPRedisRepository,
			otpdal.NewOTPPostgresqlRepository,
			otp.NewGenerator,
			fx.Annotate(
				otp.NewGenerator,
				fx.As(new(contracts.OtpGenerator)),
			),

			jwt.NewJWT,
			jwt.NewConfig,
		),
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
