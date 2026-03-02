package auth

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	authusecases "github.com/thumbrise/demo/golang-demo/internal/modules/auth/application/usecases"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	otpdal "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/jwt"
	authmailers "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	otp "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/otp"
)

var Bindings = wire.NewSet(
	NewModule,
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
	wire.Bind(
		new(contracts.OtpGenerator),
		new(*otp.Generator),
	),

	jwt.NewJWT,
	jwt.NewConfig,
)

type Module struct {
	router *http.Router
}

func NewModule(router *http.Router) Module {
	return Module{
		router: router,
	}
}

func (m *Module) Name() string {
	return "auth"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	m.router.Register()

	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}

func (m *Module) LongRun(ctx context.Context) error {
	return nil
}
