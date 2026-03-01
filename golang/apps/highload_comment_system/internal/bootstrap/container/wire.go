//go:build wireinject

package container

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
	"github.com/thumbrise/demo/golang-demo/internal"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/config"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/components"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal"
	otp2 "github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	authusecases "github.com/thumbrise/demo/golang-demo/internal/modules/auth/application/usecases"
	authhttp "github.com/thumbrise/demo/golang-demo/internal/modules/auth/endpoints/http"
	authmailers "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	observabilitymiddlewares "github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	observabilityrouters "github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	observabilityprofiler "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	sharederrorsmap "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	sharederrorsmaprouters "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis"
	rediscomponents "github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis/infrastructure/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
	swaggerhttp "github.com/thumbrise/demo/golang-demo/internal/modules/swagger/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/user"
	userusecases "github.com/thumbrise/demo/golang-demo/internal/modules/user/application/usecases"
	userrouters "github.com/thumbrise/demo/golang-demo/internal/modules/user/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/pkg/env"
	"github.com/thumbrise/demo/golang-demo/pkg/otp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

var sAll = wire.NewSet(
	// core
	internal.Bootloaders,
	bootstrap.NewRunner,
	cmd.NewBootloader,
	cmd.NewKernel,
	http.NewKernel,

	// cmd
	cmds.NewServe,
	cmds.NewRoute,
	cmds.NewRouteList,

	// module - sharederrorsmap
	sharederrorsmap.NewBootloader,
	sharederrorsmaprouters.NewErrorsMapMiddleware,
	sharederrorsmaprouters.NewErrorsMapRouter,

	// module - sharedredis
	redis.NewBootloader,
	rediscomponents.NewRedisClient,

	// module - shared - sharedswagger
	swagger.NewBootloader,
	swaggerhttp.NewSwaggerRouter,

	// module - observability
	observability.NewBootloader,
	observabilitymiddlewares.NewObservabilityMiddleware,
	observabilityrouters.NewHealthRouter,
	observabilityrouters.NewObservabilityRouter,
	observabilityrouters.NewPprofRouter,
	wire.NewSet(
		tracer.NewTracerProvider,
		wire.Bind(new(oteltracer.TracerProvider), new(*sdktrace.TracerProvider)),
	),
	tracer.NewTracer,

	// module - auth
	auth.NewBootloader,
	authhttp.NewMiddleware,
	authhttp.NewRouter,
	authmailers.NewOTPMail,
	authusecases.NewAuthCommandSignIn,
	authusecases.NewAuthCommandExchangeOtp,
	authusecases.NewAuthQueryMe,
	authusecases.NewAuthCommandRefresh,

	// module - user
	user.NewBootloader,
	userrouters.NewUsersRouter,
	userusecases.NewUserQueryOne,

	// module - observability
	observabilityprofiler.NewProfiler,

	// infrastructure
	components.NewLogger,
	components.NewJWT,
	components.MustConnect,

	wire.NewSet(
		wire.Bind(new(contracts.EnvLoader), new(*env.Loader)),
		env.NewLoader,
	),
	wire.NewSet(
		wire.Bind(new(contracts.OtpGenerator), new(*otp.Generator)),
		otp.NewGenerator,
	),

	// configs
	config.NewApp,
	config.NewDB,
	config.NewMail,
	config.NewAuth,
	config.NewHttp,
	config.NewObservability,
	config.NewRedis,

	// dal
	dal.NewUserRepository,
	otp2.NewOTPRedisRepository,
	otp2.NewOTPPostgresqlRepository,
)

func InitializeContainer(ctx context.Context) (*Container, error) {
	wire.Build(
		NewContainer,
		sAll,
	)

	return &Container{}, nil
}
