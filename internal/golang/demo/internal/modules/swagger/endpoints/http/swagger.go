package http

//	@title						Taskman API
//	@version					1.0
//	@description				Taskman API
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token. <-- description (optional)
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/thumbrise-task-manager/task-manager-backend/docs"
	http2 "gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
)

type SwaggerRouter struct {
	kernel *http2.Kernel
}

func NewSwaggerRouter(kernel *http2.Kernel) *SwaggerRouter {
	return &SwaggerRouter{kernel: kernel}
}

func (h *SwaggerRouter) Register() {
	h.kernel.Gin().Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/swagger" || c.Request.URL.Path == "/swagger/" {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			c.Abort()

			return
		}
	})
	h.kernel.Gin().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
