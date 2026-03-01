package components

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/config"
)

const JWTContextKeyUser = "jwt"

type JWT struct {
	config config.Auth
}

func NewJWT(config config.Auth) *JWT {
	return &JWT{config: config}
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Meta JWTClaimsMeta `json:"meta"`
}
type JWTClaimsMeta struct {
	UserID   int      `json:"userId"`
	Email    string   `json:"email,omitempty"`
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}
type JWTPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (j *JWT) Issue(userID int, username string, email string, roles []string) (*JWTPair, error) {
	meta := JWTClaimsMeta{
		UserID:   userID,
		Username: username,
		Email:    email,
		Roles:    roles,
	}

	accessToken, err := j.createSignedClaims(meta, j.config.JWTAccessTTLMinutes)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.createSignedClaims(meta, j.config.JWTRefreshTTLMinutes)
	if err != nil {
		return nil, err
	}

	return &JWTPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *JWT) createSignedClaims(meta JWTClaimsMeta, ttlMinutes int) (string, error) {
	c := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttlMinutes) * time.Minute)),
			Subject:   strconv.Itoa(meta.UserID),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.JWTIssuer,
		},
		Meta: meta,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return t.SignedString([]byte(j.config.JWTSecret))
}

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

func (j *JWT) Parse(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %s", ErrUnexpectedSigningMethod, token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.config.JWTSecret), nil
	})

	switch {
	case t != nil && t.Valid:
		return claims, nil
	// case errors.Is(err, jwt.ErrTokenMalformed):
	//	return nil, NewAuthJwtError(err.Error())
	// case errors.Is(err, jwt.ErrTokenSignatureInvalid):
	//	return nil, NewAuthJwtError(err.Error())
	// case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
	//	return nil, NewAuthJwtError(err.Error())
	default:
		return nil, fmt.Errorf("couldn't handle this token: %w", err)
	}
}

var ErrUnexpectedJWTClaimsType = errors.New("unexpected JWTClaims type")

func (j *JWT) ClaimsFromGinContext(ctx *gin.Context) *JWTClaims {
	claimsRaw := ctx.MustGet(JWTContextKeyUser)

	result, ok := claimsRaw.(*JWTClaims)
	if !ok {
		panic(fmt.Errorf("%w: %T", ErrUnexpectedJWTClaimsType, claimsRaw))
	}

	return result
}
