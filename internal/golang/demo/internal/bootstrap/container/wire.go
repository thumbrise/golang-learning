//go:build wireinject

package container

import (
	"context"

	"github.com/google/wire"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/cmd"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/cmd/cmds"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/bootstrap"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/config"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/contracts"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/components"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/dal"
	otp2 "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/dal/otp"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth"
	authusecases "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth/application/usecases"
	authhttp "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth/endpoints/http"
	authmailers "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth/infrastructure/mailers"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability"
	observabilitymiddlewares "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability/endpoints/http/middlewares"
	observabilityrouters "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability/endpoints/http/routers"
	observabilityprofiler "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability/infrastructure/components/profiler"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability/infrastructure/components/tracer"
	sharederrorsmap "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/errorsmap"
	sharederrorsmaprouters "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/errorsmap/endpoints/http"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/redis"
	rediscomponents "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/redis/infrastructure/components"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/swagger"
	swaggerhttp "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/swagger/endpoints/http"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/user"
	userusecases "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/user/application/usecases"
	userrouters "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/user/endpoints/http/routers"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/pkg/env"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/pkg/otp"
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
