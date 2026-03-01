//go:build test

package otp_test

import (
	"testing"

	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/otp"
)

func TestGenerator_Generate(t *testing.T) {
	const length = 6

	t.Run("generate generates string of length", func(t *testing.T) {
		g := otp.NewGenerator()
		code, _ := g.Generate(length)

		got := len(code)
		if got != length {
			t.Errorf("Generate() = %v, want %v", got, length)
		}
	})

	t.Run("generate generates different strings", func(t *testing.T) {
		g := otp.NewGenerator()

		results := map[string]struct{}{}

		times := 10
		for range times {
			code, _ := g.Generate(length)
			results[code] = struct{}{}
		}

		if len(results) != times {
			t.Errorf("Generate() = %v, want 10 different strings", len(results))
		}
	})
}
