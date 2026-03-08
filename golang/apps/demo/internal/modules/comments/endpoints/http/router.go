package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/usecases"
	bootstraphttp "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/components"
)

type Router struct {
	kernel                 *bootstraphttp.Kernel
	commentsCommandPublish *usecases.CommentsCommandPublish
}

func NewRouter(commentsCommandPublish *usecases.CommentsCommandPublish, kernel *bootstraphttp.Kernel) *Router {
	return &Router{commentsCommandPublish: commentsCommandPublish, kernel: kernel}
}

func (h *Router) Register() {
	grp := h.kernel.Gin().Group("/api/comments")

	h.publish(grp)
}

//	@Summary	Publish
//	@Tags		Comments
//	@Accept		json
//	@Produce	json
//	@Param		input	body		usecases.CommentsCommandPublishInput	true	"input"
//	@Success	200		{object}	usecases.CommentsCommandPublishOutput
//	@Failure	400		{object}	usecases.CommentsCommandPublishOutput
//	@Router		/api/comments/publish [post]
//
// -
func (h *Router) publish(grp gin.IRoutes) {
	grp.POST("/publish", func(ctx *gin.Context) {
		var input usecases.CommentsCommandPublishInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			_ = ctx.Error(err)

			return
		}

		output, err := h.commentsCommandPublish.Handle(ctx, input)
		if err != nil {
			_ = ctx.Error(err)

			return
		}

		ctx.JSON(http.StatusOK, output)
	})
}
