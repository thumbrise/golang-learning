package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/dal"
	domainerrors "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/domain/errors"
)

type UserQueryOne struct {
	logger         *slog.Logger
	userRepository *dal.UserRepository
}

func NewUserQueryOne(logger *slog.Logger, userRepository *dal.UserRepository) *UserQueryOne {
	return &UserQueryOne{logger: logger, userRepository: userRepository}
}

type UserQueryOneInput struct {
	Id int `binding:"required"`
}

var ErrFailedToFindUser = errors.New("failed to find user")

func (u *UserQueryOne) Handle(ctx context.Context, input UserQueryOneInput) (*UserQueryOneOutput, error) {
	u.logger.Info("UserQueryOne",
		slog.Any("input", input),
	)

	user, err := u.userRepository.FindByID(ctx, input.Id)
	if err != nil {
		if dal.IsNotFound(err) {
			return nil, domainerrors.NewNotFoundError("user not found")
		}

		return nil, fmt.Errorf("%w: %w", ErrFailedToFindUser, err)
	}

	return &UserQueryOneOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

type UserQueryOneOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	BornAt    time.Time `json:"bornAt"`
	CreatedAt time.Time `json:"createdAt"`
}
