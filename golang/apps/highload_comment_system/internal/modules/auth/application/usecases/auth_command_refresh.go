package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/jwt"
)

type AuthCommandRefresh struct {
	logger *slog.Logger
	jwt    *jwt.JWT
}

func NewAuthCommandRefresh(logger *slog.Logger, jwt *jwt.JWT) *AuthCommandRefresh {
	return &AuthCommandRefresh{logger: logger, jwt: jwt}
}

type AuthCommandRefreshInput struct {
	RefreshToken string `binding:"required" example:"" json:"refreshToken"`
}

type AuthCommandRefreshOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var ErrFailedParseRefreshToken = errors.New("failed to parse refresh token")

func (a *AuthCommandRefresh) Handle(_ context.Context, input AuthCommandRefreshInput) (*AuthCommandRefreshOutput, error) {
	a.logger.Info("AuthCommandRefresh",
		slog.Any("input", input),
	)
	//nolint:godox
	// TODO: check refresh token blacklist
	claims, err := a.jwt.Parse(input.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedParseRefreshToken, err)
	}
	//nolint:godox
	// TODO: mark refresh token blacklisted

	//nolint:godox
	// TODO: get roles from user
	jwtPair, err := a.jwt.Issue(
		claims.Meta.UserID,
		claims.Meta.Username,
		claims.Meta.Email,
		claims.Meta.Roles,
	)
	if err != nil {
		return nil, err
	}

	return &AuthCommandRefreshOutput{
		AccessToken:  jwtPair.AccessToken,
		RefreshToken: jwtPair.RefreshToken,
	}, nil
}
