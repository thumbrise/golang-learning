package otp

import (
	"context"
	"fmt"

	database2 "github.com/thumbrise/demo/golang-demo/internal/modules/plugins/database"
)

type OTPPostresqlRepository struct {
	db *database2.DB
}

func NewOTPPostgresqlRepository(db *database2.DB) *OTPPostresqlRepository {
	return &OTPPostresqlRepository{db: db}
}

func (r *OTPPostresqlRepository) Create(ctx context.Context, otp *OTP) error {
	sql := "INSERT INTO otps (user_id, code, expires_at, created_at) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int
	if err := r.db.Pool().QueryRow(ctx, sql, otp.UserID, otp.Code, otp.ExpiresAt, otp.CreatedAt).Scan(&id); err != nil {
		return fmt.Errorf("%w: %w", database2.ErrFailedQuery, err)
	}

	otp.ID = id

	return nil
}

func (r *OTPPostresqlRepository) ExistsValid(ctx context.Context, userID int, code string) (bool, error) {
	sql := "SELECT EXISTS (SELECT 1 FROM otps WHERE user_id = $1 AND code = $2 AND expires_at > NOW())"

	var exists bool
	if err := r.db.Pool().QueryRow(ctx, sql, userID, code).Scan(&exists); err != nil {
		return false, fmt.Errorf("%w: %w", database2.ErrFailedQuery, err)
	}

	return exists, nil
}
