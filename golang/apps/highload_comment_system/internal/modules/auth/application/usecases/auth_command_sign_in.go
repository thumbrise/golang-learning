package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/config"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal"
	otp2 "github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	"github.com/thumbrise/demo/golang-demo/pkg/otp"
)

type AuthCommandSignIn struct {
	logger                  *slog.Logger
	otpMailer               *mailers.OTPMailer
	config                  config.Auth
	otpGenerator            *otp.Generator
	userRepository          *dal.UserRepository
	otpRedisRepository      *otp2.OTPRedisRepository
	otpPostgresqlRepository *otp2.OTPPostresqlRepository
}

func NewAuthCommandSignIn(logger *slog.Logger,
	otpMailer *mailers.OTPMailer,
	config config.Auth,
	otpGenerator *otp.Generator,
	userRepository *dal.UserRepository,
	otpRepository *otp2.OTPRedisRepository,
	otpPostgresqlRepository *otp2.OTPPostresqlRepository,
) *AuthCommandSignIn {
	return &AuthCommandSignIn{
		logger:                  logger,
		otpMailer:               otpMailer,
		config:                  config,
		otpGenerator:            otpGenerator,
		userRepository:          userRepository,
		otpRedisRepository:      otpRepository,
		otpPostgresqlRepository: otpPostgresqlRepository,
	}
}

type AuthCommandSignInInput struct {
	Email string `binding:"required,email" example:"murat@murat.murat" json:"email"`
}

type AuthCommandSignInOutput struct {
	Message string `json:"message"`
}

var (
	ErrFailedGenerateOTPToken = errors.New("failed to generate OTP token")
	ErrFailedEnsureExists     = errors.New("failed to ensure exists")
	ErrFailedEmailExists      = errors.New("failed to check email exists")
	ErrFailedCreate           = errors.New("failed to create")
	ErrFailedSendOTPMail      = errors.New("failed to send OTP mail")
	ErrFailedCount            = errors.New("failed to count")
)

func (a *AuthCommandSignIn) Handle(ctx context.Context, input AuthCommandSignInInput) (*AuthCommandSignInOutput, error) {
	a.logger.Info("AuthCommandSignIn",
		slog.Any("input", input),
	)

	output := &AuthCommandSignInOutput{
		Message: "Check you email",
	}

	user, err := a.ensureExists(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedEnsureExists, err)
	}

	otpEntity, err := a.generateCode(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedGenerateOTPToken, err)
	}

	expiredAt := time.Now().Add(time.Minute * time.Duration(a.config.OTPTTLMinutes))

	err = a.otpMailer.Send(input.Email, otpEntity.Code, expiredAt)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedSendOTPMail, err)
	}

	return output, nil
}

func (a *AuthCommandSignIn) ensureExists(ctx context.Context, email string) (*dal.User, error) {
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if dal.IsNotFound(err) {
			return a.createUser(ctx, email)
		}

		return nil, fmt.Errorf("%w: %w", ErrFailedEmailExists, err)
	}

	if user == nil {
		return a.createUser(ctx, email)
	}

	return user, nil
}

func (a *AuthCommandSignIn) createUser(ctx context.Context, email string) (*dal.User, error) {
	count, err := a.userRepository.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedCount, err)
	}

	name := "guest" + strconv.Itoa(count+1)

	user := &dal.User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err = a.userRepository.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedCreate, err)
	}

	return user, nil
}

func (a *AuthCommandSignIn) generateCode(ctx context.Context, user *dal.User) (*otp2.OTP, error) {
	expiredAt := time.Now().Add(time.Minute * time.Duration(a.config.OTPTTLMinutes))
	otpEntity := &otp2.OTP{
		UserID:    user.ID,
		ExpiresAt: expiredAt,
	}

	var otpCode string
	if a.config.OTPForcedValue != "" {
		otpCode = a.config.OTPForcedValue
		otpEntity.Code = otpCode
	} else {
		var err error

		otpCode, err = a.otpGenerator.Generate(a.config.OTPLength)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrFailedGenerateOTPToken, err)
		}

		otpEntity.Code = otpCode
	}

	err := a.otpRedisRepository.Create(ctx, otpEntity)
	// err := a.otpPostgresqlRepository.Create(ctx, otpEntity)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedCreate, err)
	}

	return otpEntity, nil
}
