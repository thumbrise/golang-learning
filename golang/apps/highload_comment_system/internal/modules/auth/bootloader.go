package auth

import (
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

var Module = fx.Module("auth",
	fx.Provide(
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
	fx.Invoke(func(router *http.Router) {
		router.Register()
	}),
)
