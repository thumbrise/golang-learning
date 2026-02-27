package array_test

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/thumbrise/demo/internal/structures/array"
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
	testArrayGeneric(t, 5, map[int]testItem[int]{
		0: {42, true},
		1: {100, false},
		2: {7, true},
		3: {0, false},
		4: {12345, true},
	}, 0)
}

func TestArrayWithString(t *testing.T) {
	testArrayGeneric(t, 5, map[int]testItem[string]{
		0: {"apple", true},
		1: {"banana", false},
		2: {"cherry", true},
		3: {"date", false},
		4: {"elderberry", true},
	}, "")
}

func TestArrayWithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	zeroPerson := Person{}
	testArrayGeneric(t, 3, map[int]testItem[Person]{
		0: {Person{"Alice", 30}, true},
		1: {Person{"Bob", 25}, false},
		2: {Person{"Charlie", 35}, true},
	}, zeroPerson)
}

type testItem[T any] struct {
	value    T
	toDelete bool
}

func testArrayGeneric[T any](t *testing.T, size int, items map[int]testItem[T], zero T) {
	t.Helper()

	arr := array.NewArray[T](size)
	defer arr.Clear()

	// Проверка длины
	if arr.Len() != size {
		t.Fatalf("Len: expected %d, got %d", size, arr.Len())
	}

	// Установка значений
	for idx, item := range items {
		arr.Set(idx, item.value)
	}

	// Проверка Get после Set
	for idx, item := range items {
		got := arr.Get(idx)
		if !reflect.DeepEqual(got, item.value) {
			t.Fatalf("Get(%d): expected %+v, got %+v", idx, item.value, got)
		}
	}

	// «Удаление» (зануление) помеченных элементов
	for idx, item := range items {
		if item.toDelete {
			arr.Set(idx, zero)
		}
	}

	// Проверка, что удалённые стали нулевыми, а остальные не изменились
	for idx, item := range items {
		got := arr.Get(idx)
		if item.toDelete {
			if !reflect.DeepEqual(got, zero) {
				t.Fatalf("After delete, index %d: expected zero (%+v), got %+v", idx, zero, got)
			}
		} else {
			if !reflect.DeepEqual(got, item.value) {
				t.Fatalf("After delete, index %d: expected %+v, got %+v", idx, item.value, got)
			}
		}
	}

	// Тесты граничных индексов
	boundsTests := []struct {
		name      string
		bound     int
		wantPanic bool
	}{
		{"negative", -1, true},
		{"equal to size", size, true},
		{"zero", 0, false},
		{"last", size - 1, false},
		{"size+1", size + 1, true},
		{"half", size / 2, false},
		{"size*2", size * 2, true},
		{"-size", -size, true},
		{"far negative", size - 10000000, true},
		{"far positive", size * size, true},
		{"max int", int(^uint(0) >> 1), true},
	}

	for _, tc := range boundsTests {
		t.Run("bounds/"+tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && tc.wantPanic {
					t.Errorf("Expected panic for index %d, but got none", tc.bound)
				}
			}()

			arr.Set(tc.bound, zero)
		})
	}

	// Тест Clear
	t.Run("Clear", func(t *testing.T) {
		arr.Clear()

		if !arr.IsCleared() {
			t.Error("IsCleared = false after Clear")
		}

		// Попытки доступа после очистки должны паниковать
		for idx := range items {
			func(i int) {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic after Clear for index %d", i)
					}
				}()

				arr.Get(i)
			}(idx)

			func(i int) {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic after Clear for index %d", i)
					}
				}()

				arr.Set(i, zero)
			}(idx)
		}
	})
}
