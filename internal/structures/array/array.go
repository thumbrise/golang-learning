// Package array предоставляет самописную реализацию массива, работающую с сырой памятью и ручным управлением очистки.
// В пакете вовсе не используются стандыртные GO массивы и слайсы. Полностью самостоятельная реализация для учебных целей.
package array

import (
	"log/slog"
	"reflect"
	"unsafe"
)

//#include <stdlib.h>
import "C"

// Array - Самописная реализация массива, работающая с сырой памятью и ручным управлением очистки.
type Array[T any] struct {
	data unsafe.Pointer
	len  int
}

func NewArray[T any](length int) *Array[T] {
	return &Array[T]{
		len:  length,
		data: alloc[T](length),
	}
}

// Get возвращает значение по индексу
//
//nolint:ireturn // Надо
func (a *Array[T]) Get(index int) T {
	return *new(T)
}

// Set устанавливает значение по индексу
func (a *Array[T]) Set(index int, value T) bool {
	// *a.data = value
	return true
}

// Clear освобождает память, выделенную под массив
// Не забывайте вызывать этот метод, чтобы избежать утечек памяти
func (a *Array[T]) Clear() {
	if a.data != nil {
		C.free(a.data)
		a.data = nil
		a.len = 0
	}
}

func (a *Array[T]) Len() int {
	return a.len
}

func (a *Array[T]) SizeBytes() uintptr {
	var t T

	return uintptr(a.len) * unsafe.Sizeof(t)
}

// Data возвращает сырой указатель на данные массива
// ВНИМАНИЕ: Используйте с осторожностью! Убедитесь, что массив не был очищен.
func (a *Array[T]) Data() unsafe.Pointer {
	return a.data
}

func (a *Array[T]) IsCleared() bool {
	return a.data == nil
}

func alloc[T any](length int) unsafe.Pointer {
	var t T
	elemsize := unsafe.Sizeof(t)
	ptr := C.malloc(C.size_t(length) * C.size_t(elemsize))
	if ptr == nil {
		panic("failed to allocate memory pointer")
	}

	slog.Debug("ALLOCATED",
		slog.String("type", reflect.TypeOf(t).String()),
		slog.Int("size", length),
		slog.Any("elemsize", elemsize),
		slog.Any("fullsize", uintptr(length)*elemsize),
		slog.Any("ptr", ptr),
	)

	return ptr
}
