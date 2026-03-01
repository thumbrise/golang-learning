package usecases

import (
	"context"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/jwt"
)

type AuthQueryMe struct {
	logger *slog.Logger
}

func NewAuthQueryMe(logger *slog.Logger) *AuthQueryMe {
	return &AuthQueryMe{logger: logger}
}

type AuthQueryMeInput struct {
	Claims *jwt.JWTClaims
}

type AuthQueryMeOutput struct {
	// Claims *components.JWTClaims
	Message string
}

func (a *AuthQueryMe) Handle(ctx context.Context, input AuthQueryMeInput) (*AuthQueryMeOutput, error) {
	a.logger.Info("AuthQueryMe",
		slog.Any("input", input),
	)

	return &AuthQueryMeOutput{
		// Claims: input.Claims,
		Message: "Success",
	}, nil
}
