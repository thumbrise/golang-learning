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

func TestArrayWithInt(t *testing.T) {
	testsInt := []struct {
		name     string
		size     int
		toSet    map[int]int
		toDelete []int
	}{
		{
			name: "int_5",
			size: 5,
			toSet: map[int]int{
				0: 1,
				1: 2,
				2: 3,
				3: 4,
				4: 5,
			},
			toDelete: []int{0, 2, 4},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			arr := array.NewArray[int](test.size)
			if arr.Len() != test.size {
				t.Fatalf("Array length should be %d, got %d", test.size, arr.Len())
			}

			for k, v := range test.toSet {
				if !arr.Set(k, v) {
					t.Fatalf("Array Set should return true for k=%d v=%d", k, v)
				}
			}

			for k := range test.toSet {
				if arr.Get(k) != test.toSet[k] {
					t.Fatalf("Array value after Set(k, v) should be %d, got %d", test.toSet[k], arr.Get(k))
				}
			}

			for _, k := range test.toDelete {
				if !arr.Set(k, 0) {
					t.Fatalf("Array Set should return true for k=%d v=0", k)
				}
			}

			for k := range test.toDelete {
				if arr.Get(k) != 0 {
					t.Fatalf("Array value after Set(k, 0) should be 0, got %d", arr.Get(k))
				}
			}

			for k := range test.toSet {
				if arr.Get(k) != test.toSet[k] {
					t.Fatalf("Array value after Set(k, v) should be %d, got %d", test.toSet[k], arr.Get(k))
				}
			}

			arr.Clear()

			if !arr.IsCleared() {
				t.Fatal("Array should be cleared")
			}

			for k := range test.toSet {
				if arr.Get(k) != 0 {
					t.Fatalf("Array value after Clear() should be 0, got %d", arr.Get(k))
				}
			}

			for k := range test.toSet {
				if arr.Set(k, test.toSet[k]) {
					t.Fatalf("Array Set should return false after Clear()")
				}
			}
		})
	}
}
