package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/components"
)

type Middleware struct {
	jwt *components.JWT
}

func NewMiddleware(jwt *components.JWT) *Middleware {
	return &Middleware{jwt: jwt}
}

func (m *Middleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims, err := m.jwt.Parse(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()

			return
		}

		c.Set(components.JWTContextKeyUser, claims)

		c.Next()
	}
}
