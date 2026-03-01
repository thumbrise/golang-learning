package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	otpdal "github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/mailers"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/otp"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
)

type AuthCommandSignIn struct {
	logger                  *slog.Logger
	otpMailer               *mailers.OTPMailer
	otpConfig               otp.Config
	otpGenerator            contracts.OtpGenerator
	userRepository          *dal.UserRepository
	otpRedisRepository      *otpdal.OTPRedisRepository
	otpPostgresqlRepository *otpdal.OTPPostresqlRepository
}

func NewAuthCommandSignIn(logger *slog.Logger,
	otpMailer *mailers.OTPMailer,
	config otp.Config,
	otpGenerator contracts.OtpGenerator,
	userRepository *dal.UserRepository,
	otpRepository *otpdal.OTPRedisRepository,
	otpPostgresqlRepository *otpdal.OTPPostresqlRepository,
) *AuthCommandSignIn {
	return &AuthCommandSignIn{
		logger:                  logger,
		otpMailer:               otpMailer,
		otpConfig:               config,
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

	expiredAt := time.Now().Add(time.Minute * time.Duration(a.otpConfig.TTLMinutes))

	err = a.otpMailer.Send(input.Email, otpEntity.Code, expiredAt)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedSendOTPMail, err)
	}

	return output, nil
}

func (a *AuthCommandSignIn) ensureExists(ctx context.Context, email string) (*dal.User, error) {
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if database.IsNotFound(err) {
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

func (a *AuthCommandSignIn) generateCode(ctx context.Context, user *dal.User) (*otpdal.OTP, error) {
	expiredAt := time.Now().Add(time.Minute * time.Duration(a.otpConfig.TTLMinutes))
	otpEntity := &otpdal.OTP{
		UserID:    user.ID,
		ExpiresAt: expiredAt,
	}

	var otpCode string
	if a.otpConfig.ForcedValue != "" {
		otpCode = a.otpConfig.ForcedValue
		otpEntity.Code = otpCode
	} else {
		var err error

		otpCode, err = a.otpGenerator.Generate(a.otpConfig.Length)
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
