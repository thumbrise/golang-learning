package otp

import "time"

type OTP struct {
	ID        int
	UserID    int
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}
