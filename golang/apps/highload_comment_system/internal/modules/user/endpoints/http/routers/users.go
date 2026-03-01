package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bootstraphttp "github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/user/application/usecases"
)

type UsersRouter struct {
	userQueryOne *usecases.UserQueryOne
	kernel       *bootstraphttp.Kernel
}

func NewUsersRouter(userQueryOne *usecases.UserQueryOne, kernel *bootstraphttp.Kernel) *UsersRouter {
	return &UsersRouter{userQueryOne: userQueryOne, kernel: kernel}
}

func (u *UsersRouter) Register() {
	grp := u.kernel.Gin().Group("/api/users")

	grp.GET("/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		input := usecases.UserQueryOneInput{Id: id}

		output, err := u.userQueryOne.Handle(ctx, input)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}
