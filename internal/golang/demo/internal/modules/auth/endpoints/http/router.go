package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/components"
	bootstraphttp "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
	usecases2 "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth/application/usecases"
)

type Router struct {
	kernel                 *bootstraphttp.Kernel
	authMiddleware         *Middleware
	jwt                    *components.JWT
	authCommandSignIn      *usecases2.AuthCommandSignIn
	authCommandExchangeOtp *usecases2.AuthCommandExchangeOtp
	authQueryMe            *usecases2.AuthQueryMe
	authCommandRefresh     *usecases2.AuthCommandRefresh
}

func NewRouter(
	kernel *bootstraphttp.Kernel,
	authCommandSignUp *usecases2.AuthCommandSignIn,
	authCommandExchangeOtp *usecases2.AuthCommandExchangeOtp,
	authQueryMe *usecases2.AuthQueryMe,
	authCommandRefresh *usecases2.AuthCommandRefresh,
	authMiddleware *Middleware,
	jwt *components.JWT,
) *Router {
	return &Router{
		kernel:                 kernel,
		authCommandSignIn:      authCommandSignUp,
		authCommandExchangeOtp: authCommandExchangeOtp,
		authQueryMe:            authQueryMe,
		authMiddleware:         authMiddleware,
		authCommandRefresh:     authCommandRefresh,
		jwt:                    jwt,
	}
}

func (h *Router) Register() {
	grp := h.kernel.Gin().Group("/api/auth")

	h.signIn(grp)
	h.exchangeOtp(grp)
	h.refresh(grp)

	protected := grp.Use(h.authMiddleware.Handler())
	h.me(protected)
}

//	@Summary	Sign in
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		input	body		usecases.AuthCommandSignInInput	true	"input"
//	@Success	200		{object}	usecases.AuthCommandSignInOutput
//	@Failure	400		{object}	usecases.AuthCommandSignInOutput
//	@Router		/api/auth/sign-in [post]
//
// -
func (h *Router) signIn(grp gin.IRoutes) {
	grp.POST("/sign-in", func(ctx *gin.Context) {
		var input usecases2.AuthCommandSignInInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			_ = ctx.Error(err)

			return
		}

		output, err := h.authCommandSignIn.Handle(ctx, input)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}

//	@Summary	Exchange OTP
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		input	body		usecases.AuthCommandExchangeOtpInput	true	"input"
//	@Success	200		{object}	usecases.AuthCommandExchangeOtpOutput
//	@Failure	400		{object}	usecases.AuthCommandExchangeOtpOutput
//	@Router		/api/auth/exchange-otp [post]
//
// -
func (h *Router) exchangeOtp(grp gin.IRoutes) {
	grp.POST("/exchange-otp", func(ctx *gin.Context) {
		var input usecases2.AuthCommandExchangeOtpInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			_ = ctx.Error(err)

			return
		}

		output, err := h.authCommandExchangeOtp.Handle(ctx, input)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}

//	@Summary	Me
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	usecases.AuthQueryMeOutput
//	@Failure	400	{object}	usecases.AuthQueryMeOutput
//	@Router		/api/auth/me [get]
//	@Security	ApiKeyAuth
//
// -
func (h *Router) me(grp gin.IRoutes) {
	grp.GET("/me", func(ctx *gin.Context) {
		claims := h.jwt.ClaimsFromGinContext(ctx)

		output, err := h.authQueryMe.Handle(ctx, usecases2.AuthQueryMeInput{Claims: claims})
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}

//	@Summary	Refresh
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		input	body		usecases.AuthCommandRefreshInput	true	"input"
//	@Success	200		{object}	usecases.AuthCommandRefreshOutput
//	@Failure	400		{object}	usecases.AuthCommandRefreshOutput
//	@Router		/api/auth/refresh [post]
//
// -
func (h *Router) refresh(grp gin.IRoutes) {
	grp.POST("/refresh", func(ctx *gin.Context) {
		var input usecases2.AuthCommandRefreshInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			_ = ctx.Error(err)

			return
		}

		output, err := h.authCommandRefresh.Handle(ctx, input)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}
