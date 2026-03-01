package contracts

import "github.com/gin-gonic/gin"

type JWT interface {
	ClaimsFromGinContext(ctx *gin.Context) *JWTClaims
	Issue(userID int, username string, email string, roles []string) (*JWTPair, error)
	Parse(tokenString string) (*JWTClaims, error)
}
