package main

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/modules"
	"github.com/thumbrise/demo/golang-demo/internal/config"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	authusecases "github.com/thumbrise/demo/golang-demo/internal/modules/auth/application/usecases"
	authhttp "github.com/thumbrise/demo/golang-demo/internal/modules/auth/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	otp3 "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/jwt"
	authmailers "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	homepagehttp "github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	homepagegenerator "github.com/thumbrise/demo/golang-demo/internal/modules/homepage/infrastucture/generator"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	observabilitymiddlewares "github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/middlewares"
	observabilityrouters "github.com/thumbrise/demo/golang-demo/internal/modules/observability/endpoints/http/routers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/logger"
	observabilityprofiler "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/profiler"
	observabilitytracer "github.com/thumbrise/demo/golang-demo/internal/modules/observability/infrastructure/components/tracer"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	sharederrorsmap "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	sharederrorsmaprouters "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/mail"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis"
	rediscomponents "github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis/infrastructure/components"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
	swaggerhttp "github.com/thumbrise/demo/golang-demo/internal/modules/swagger/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/pkg/env"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ModuleCore предоставляет базовые компоненты приложения
var ModuleCore = fx.Options(
	fx.Provide(
		// internal.Bootloaders,
		bootstrap.NewRunner,
		cmd.NewBootloader,
		cmd.NewKernel,
		http.NewKernel,
	),
)

// ModuleCmd предоставляет команды CLI
var ModuleCmd = fx.Options(
	fx.Provide(
		cmds.NewServe,
		cmds.NewRoute,
		cmds.NewRouteList,
	),
)

// ModuleSharedErrorsMap предоставляет middleware для карты ошибок
var ModuleSharedErrorsMap = fx.Options(
	fx.Provide(
		sharederrorsmap.NewBootloader,
		sharederrorsmaprouters.NewErrorsMapMiddleware,
		sharederrorsmaprouters.NewErrorsMapRouter,
	),
)

// ModuleSharedRedis предоставляет Redis клиент
var ModuleSharedRedis = fx.Options(
	fx.Provide(
		redis.NewBootloader,
		rediscomponents.NewRedisClient,
	),
)

// ModuleSwagger предоставляет Swagger UI
var ModuleSwagger = fx.Options(
	fx.Provide(
		swagger.NewBootloader,
		swaggerhttp.NewSwaggerRouter,
	),
)

// ModuleObservability предоставляет observability компоненты (tracing, profiler, health)
var ModuleObservability = fx.Options(
	fx.Provide(
		observability.NewBootloader,
		observabilitymiddlewares.NewObservabilityMiddleware,
		observabilityrouters.NewHealthRouter,
		observabilityrouters.NewObservabilityRouter,
		observabilityrouters.NewPprofRouter,
		observabilityprofiler.NewProfiler,
	),
	// Для tracer используем fx.Annotate, чтобы указать, что конструктор возвращает *sdktrace.TracerProvider,
	// но мы хотим зарегистрировать его как oteltracer.TracerProvider.
	fx.Provide(
		observabilitytracer.NewTracerProvider,
		observabilitytracer.NewTracer,
		fx.Annotate(
			func() *sdktrace.TracerProvider {
				return &sdktrace.TracerProvider{}
			},
			fx.As(new(oteltracer.TracerProvider)),
		),
	),
)

// ModuleAuth предоставляет всё для авторизации
var ModuleAuth = fx.Options(
	fx.Provide(
		auth.NewBootloader,
		authhttp.NewMiddleware,
		authhttp.NewRouter,
		authmailers.NewOTPMail,
		authusecases.NewAuthCommandSignIn,
		authusecases.NewAuthCommandExchangeOtp,
		authusecases.NewAuthQueryMe,
		authusecases.NewAuthCommandRefresh,
	),
)

// ModuleHomepage предоставляет генератор домашней страницы
var ModuleHomepage = fx.Options(
	fx.Provide(
		homepage.NewBootloader,
		homepagehttp.NewHomePageRouter,
		homepagegenerator.NewGenerator,
	),
)

// ModuleInfrastructure предоставляет общие компоненты инфраструктуры
var ModuleInfrastructure = fx.Options(
	fx.Provide(
		logger.NewLogger,
		jwt.NewJWT,
		// fx.Annotate(
		database.NewDB,
		database.NewBootloader,
		// fx.OnStart(func(ctx context.Context, db *components.DB) error {
		//	return db.Connect(ctx)
		//}),
		//fx.OnStop(func(ctx context.Context, db *components.DB) error {
		//	db.Pool().Close()
		//	slog.Debug("Closing DB POOL")
		//	return nil
		//}),
		//),
	),
	// Привязка интерфейсов через fx.Annotate
	fx.Provide(
		fx.Annotate(
			env.NewLoader,
			fx.As(new(contracts.EnvLoader)),
		),
		fx.Annotate(
			otp.NewGenerator,
			fx.As(new(contracts.OtpGenerator)),
		),
	),
)

// ModuleConfig предоставляет все конфигурации
var ModuleConfig = fx.Options(
	fx.Provide(
		app.NewConfig,
		database.NewDB,
		mail.NewConfig,
		auth.NewConfig,
		http.NewConfig,
		observability.NewConfig,
		config.NewRedis,
	),
)

// ModuleDAL предоставляет репозитории и DAL компоненты
var ModuleDAL = fx.Options(
	fx.Provide(
		dal.NewUserRepository,
		otp3.NewOTPRedisRepository,
		otp3.NewOTPPostgresqlRepository,
	),
)

func main() {
	fx.New(
		modules.Build()...,
	).Run()
}
