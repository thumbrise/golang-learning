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
			t.Run("Bounds", func(t *testing.T) {
				testsBounds := []struct {
					name      string
					bound     int
					wantPanic bool
				}{
					{name: "negative", bound: -1, wantPanic: true},
					{name: "positive", bound: test.size, wantPanic: true},
					{name: "zero", bound: 0, wantPanic: false},
					{name: "last", bound: test.size - 1, wantPanic: false},
					{name: "additional1", bound: test.size + 1, wantPanic: true},
					{name: "additional2", bound: test.size / 2, wantPanic: false},
					{name: "additional3", bound: test.size * 2, wantPanic: true},
					{name: "additional4", bound: -test.size, wantPanic: true},
					{name: "additional5", bound: test.size - 10000000, wantPanic: true},
					{name: "additional6", bound: test.size * test.size, wantPanic: true},
					{name: "additional7", bound: int(^uint(0) >> 1), wantPanic: true},
				}

				for _, testBound := range testsBounds {
					t.Run(testBound.name, func(t *testing.T) {
						func() {
							defer func() {
								if r := recover(); r == nil && testBound.wantPanic {
									t.Errorf("Expected panic when try access index %d of array with size %d", testBound.bound, test.size)
								}
							}()

							arr.Set(testBound.bound, 42)
						}()
					})
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
func TestArrayWithString(t *testing.T) {
	type Item struct {
		value    string
		toDelete bool
	}

	tests := []struct {
		name string
		size int
		data map[int]Item
	}{
		{
			name: "string_5",
			size: 5,
			data: map[int]Item{
				0: {value: "apple", toDelete: true},
				1: {value: "banana", toDelete: false},
				2: {value: "cherry", toDelete: true},
				3: {value: "date", toDelete: false},
				4: {value: "elderberry", toDelete: true},
			},
		},
		// Добавь другие случаи при необходимости
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := array.NewArray[string](test.size)

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
						t.Fatalf("Array value after Set(i, item) should be %q, got %q", item.value, arr.Get(i))
					}
				}
			})

			t.Run("Delete", func(t *testing.T) {
				// Удаляем (зануляем) элементы с toDelete = true
				for i, item := range test.data {
					if !item.toDelete {
						continue
					}
					arr.Set(i, "") // пустая строка — аналог нуля для int
				}

				// Проверяем, что удалённые стали пустыми
				for i, item := range test.data {
					if !item.toDelete {
						continue
					}
					if arr.Get(i) != "" {
						t.Fatalf("Array deleted value should be empty string, got %q", arr.Get(i))
					}
				}

				// Проверяем, что неудалённые остались неизменными
				for i, item := range test.data {
					if item.toDelete {
						continue
					}
					if got := arr.Get(i); got != item.value {
						t.Fatalf("Array original value for index %d should be %q, got %q", i, item.value, got)
					}
				}
			})

			t.Run("Bounds", func(t *testing.T) {
				testsBounds := []struct {
					name      string
					bound     int
					wantPanic bool
				}{
					{name: "negative", bound: -1, wantPanic: true},
					{name: "positive", bound: test.size, wantPanic: true},
					{name: "zero", bound: 0, wantPanic: false},
					{name: "last", bound: test.size - 1, wantPanic: false},
					{name: "additional1", bound: test.size + 1, wantPanic: true},
					{name: "additional2", bound: test.size / 2, wantPanic: false},
					{name: "additional3", bound: test.size * 2, wantPanic: true},
					{name: "additional4", bound: -test.size, wantPanic: true},
					{name: "additional5", bound: test.size - 10000000, wantPanic: true},
					{name: "additional6", bound: test.size * test.size, wantPanic: true},
					{name: "additional7", bound: int(^uint(0) >> 1), wantPanic: true},
				}

				for _, testBound := range testsBounds {
					t.Run(testBound.name, func(t *testing.T) {
						func() {
							defer func() {
								if r := recover(); r == nil && testBound.wantPanic {
									t.Errorf("Expected panic when try access index %d of array with size %d", testBound.bound, test.size)
								}
							}()
							arr.Set(testBound.bound, "test")
						}()
					})
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

func TestArrayWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	type Item struct {
		value    Person
		toDelete bool
	}

	tests := []struct {
		name string
		size int
		data map[int]Item
	}{
		{
			name: "person_3",
			size: 3,
			data: map[int]Item{
				0: {value: Person{"Alice", 30}, toDelete: true},
				1: {value: Person{"Bob", 25}, toDelete: false},
				2: {value: Person{"Charlie", 35}, toDelete: true},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := array.NewArray[Person](test.size)

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
					got := arr.Get(i)
					if got != item.value {
						t.Fatalf("Array value after Set(i, item) should be %+v, got %+v", item.value, got)
					}
				}
			})

			t.Run("Delete", func(t *testing.T) {
				zero := Person{}
				for i, item := range test.data {
					if !item.toDelete {
						continue
					}
					arr.Set(i, zero) // «обнуляем» структуру
				}

				for i, item := range test.data {
					if !item.toDelete {
						continue
					}
					if got := arr.Get(i); got != zero {
						t.Fatalf("Array deleted value should be %+v, got %+v", zero, got)
					}
				}

				for i, item := range test.data {
					if item.toDelete {
						continue
					}
					if got := arr.Get(i); got != item.value {
						t.Fatalf("Array original value for index %d should be %+v, got %+v", i, item.value, got)
					}
				}
			})

			t.Run("Bounds", func(t *testing.T) {
				testsBounds := []struct {
					name      string
					bound     int
					wantPanic bool
				}{
					{name: "negative", bound: -1, wantPanic: true},
					{name: "positive", bound: test.size, wantPanic: true},
					{name: "zero", bound: 0, wantPanic: false},
					{name: "last", bound: test.size - 1, wantPanic: false},
					{name: "additional1", bound: test.size + 1, wantPanic: true},
					{name: "additional2", bound: test.size / 2, wantPanic: false},
					{name: "additional3", bound: test.size * 2, wantPanic: true},
					{name: "additional4", bound: -test.size, wantPanic: true},
					{name: "additional5", bound: test.size - 10000000, wantPanic: true},
					{name: "additional6", bound: test.size * test.size, wantPanic: true},
					{name: "additional7", bound: int(^uint(0) >> 1), wantPanic: true},
				}

				for _, testBound := range testsBounds {
					t.Run(testBound.name, func(t *testing.T) {
						func() {
							defer func() {
								if r := recover(); r == nil && testBound.wantPanic {
									t.Errorf("Expected panic when try access index %d of array with size %d", testBound.bound, test.size)
								}
							}()
							arr.Set(testBound.bound, Person{"X", 0})
						}()
					})
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
