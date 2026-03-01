package otp

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

var ErrFailedRandInt = errors.New("failed rand.Int")

// Generate генерирует OTP заданной длины
func (g *Generator) Generate(length int) (string, error) {
	const digits = "0123456789"

	otp := make([]byte, length)

	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrFailedRandInt, err)
		}

		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}
