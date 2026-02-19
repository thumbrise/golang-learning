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
	type Item struct {
		value    int
		toDelete bool
	}

	testsInt := []struct {
		name string
		size int
		data map[int]Item
	}{
		{
			name: "int_5",
			size: 5,
			data: map[int]Item{
				0: {value: 1, toDelete: true},
				1: {value: 2, toDelete: false},
				2: {value: 3, toDelete: true},
				3: {value: 4, toDelete: false},
				4: {value: 5, toDelete: true},
			},
		},
	}
	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			arr := array.NewArray[int](test.size)

			t.Run("Len", func(t *testing.T) {
				if arr.Len() != test.size {
					t.Fatalf("Array length should be %d, got %d", test.size, arr.Len())
				}
			})

			t.Run("Set", func(t *testing.T) {
				for i, item := range test.data {
					arr.Set(i, item.value)
				}
			})

			t.Run("Get", func(t *testing.T) {
				for i, item := range test.data {
					if arr.Get(i) != item.value {
						t.Fatalf("Array value after Set(i, item) should be %d, got %d", item.value, arr.Get(i))
					}
				}
			})

			t.Run("Delete", func(t *testing.T) {
				for i, item := range test.data {
					if !item.toDelete {
						continue
					}

					arr.Set(i, 0)
				}

				for i, item := range test.data {
					if !item.toDelete {
						continue
					}

					if arr.Get(i) != 0 {
						t.Fatalf("Array deleted value after Set(item, 0) should be 0, got %d", arr.Get(item.value))
					}
				}

				for i, item := range test.data {
					if item.toDelete {
						continue
					}

					// check original not deleted values
					got := arr.Get(i)

					want := item.value
					if got != want {
						t.Fatalf("Array original value after Set() for index %d should be %d, got %d", i, want, got)
					}
				}
			})

			t.Run("Clear", func(t *testing.T) {
				arr.Clear()
				t.Run("IsCleared", func(t *testing.T) {
					if !arr.IsCleared() {
						t.Fatal("Array should be cleared")
					}
				})
				t.Run("Get", func(t *testing.T) {
					for i, item := range test.data {
						func() {
							defer func() {
								if r := recover(); r == nil {
									t.Errorf("Expected panic after Clear() for item %#v", item)
								}
							}()

							arr.Get(i)
						}()
					}
				})

				t.Run("Set", func(t *testing.T) {
					for i, item := range test.data {
						func() {
							defer func() {
								if r := recover(); r == nil {
									t.Errorf("Expected panic after Clear() for item %#v", item)
								}
							}()

							arr.Set(i, item.value)
						}()
					}
				})
			})
		})
	}
}
