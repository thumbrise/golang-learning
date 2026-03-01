package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/components"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal/otp"
	domainerrors "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/domain/errors"
)

type AuthCommandExchangeOtp struct {
	logger         *slog.Logger
	jwt            *components.JWT
	otpRepository  *otp.OTPRedisRepository
	userRepository *dal.UserRepository
}

func NewAuthCommandExchangeOtp(logger *slog.Logger, jwt *components.JWT, otpRepository *otp.OTPRedisRepository, userRepository *dal.UserRepository) *AuthCommandExchangeOtp {
	return &AuthCommandExchangeOtp{logger: logger, jwt: jwt, otpRepository: otpRepository, userRepository: userRepository}
}

type AuthCommandExchangeOtpInput struct {
	Email string `binding:"required,email" example:"murat@murat.murat" json:"email"`
	Otp   string `binding:"required"       example:""                  json:"otp"`
}

type AuthCommandExchangeOtpOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var (
	ErrFailedFindUserByEmail = errors.New("failed to find user by email")
	ErrFailedCheckOtpExists  = errors.New("failed to check otp exists")
)

func (a *AuthCommandExchangeOtp) Handle(ctx context.Context, input AuthCommandExchangeOtpInput) (*AuthCommandExchangeOtpOutput, error) {
	a.logger.Info("AuthCommandExchangeOtp",
		slog.Any("input", input),
	)

	user, err := a.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		if dal.IsNotFound(err) {
			a.logger.Debug("AuthCommandExchangeOtp.Handle: failed find user by email", slog.Any("email", input.Email))

			return nil, domainerrors.NewUnauthenticatedError("unauthorized")
		}

		return nil, fmt.Errorf("%w: %w", ErrFailedFindUserByEmail, err)
	}

	if user == nil {
		return nil, domainerrors.NewUnauthenticatedError("unauthorized")
	}

	otpExists, err := a.otpRepository.ExistsValid(ctx, user.ID, input.Otp)
	if err != nil {
		if dal.IsNotFound(err) {
			a.logger.Debug("AuthCommandExchangeOtp.Handle: failed check otp exists", slog.Any("email", input.Email))

			return nil, domainerrors.NewUnauthenticatedError("unauthorized")
		}

		return nil, fmt.Errorf("%w: %w", ErrFailedCheckOtpExists, err)
	}

	if !otpExists {
		return nil, domainerrors.NewUnauthenticatedError("unauthorized")
	}

	id := user.ID
	username := user.Name
	email := user.Email

	//nolint:godox
	// TODO: get roles from user
	roles := []string{"customer", "moderator", "admin"}

	jwtPair, err := a.jwt.Issue(id, username, email, roles)
	if err != nil {
		return nil, err
	}

	return &AuthCommandExchangeOtpOutput{
		AccessToken:  jwtPair.AccessToken,
		RefreshToken: jwtPair.RefreshToken,
	}, nil
}
