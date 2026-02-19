package array_test

import (
	"log/slog"
	"testing"

	"github.com/thumbrise/golang-learning/internal/structures/array"
)

func TestArray(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	const expectedLength = 5

	arr := array.NewArray[string](expectedLength)

	if arr.IsCleared() {
		t.Fatal("Array should not be cleared")
	}

	if arr.Len() != expectedLength {
		t.Fatalf("Array length should be %d, got %d", expectedLength, arr.Len())
	}

	if arr.SizeBytes() == 0 {
		t.Fatal("Array size should not be 0")
	}

	if arr.Data() == nil {
		t.Fatal("Array data should not be nil")
	}

	arr.Clear()

	if !arr.IsCleared() {
		t.Fatal("Array should be cleared")
	}

	if arr.Len() != 0 {
		t.Fatal("Array length should be 0")
	}

	if arr.SizeBytes() != 0 {
		t.Fatal("Array size should be 0")
	}

	if arr.Data() != nil {
		t.Fatal("Array data should be nil")
	}
}
