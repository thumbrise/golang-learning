// Package array предоставляет самописную реализацию массива, работающую с сырой памятью и ручным управлением очистки.
// В пакете вовсе не используются стандыртные GO массивы и слайсы. Полностью самостоятельная реализация для учебных целей.
package array

import (
	"fmt"
	"unsafe"
)

// Array - Самописная реализация массива, работающая с сырой памятью и ручным управлением очистки.
type Array[T any] struct {
	data      unsafe.Pointer
	len       int
	allocator *Allocator[T]
}

func NewArray[T any](length int) *Array[T] {
	allocator := NewAllocator[T]()

	return &Array[T]{
		len:       length,
		data:      allocator.Alloc(length),
		allocator: allocator,
	}
}

// Get возвращает значение по индексу
//
//nolint:ireturn // Надо
func (a *Array[T]) Get(index int) T {
	a.checkBounds(index)

	//nolint:govet,gosec // Надо
	return *(*T)(unsafe.Pointer(a.addr(index)))
}

func (a *Array[T]) checkBounds(index int) {
	if a == nil {
		panic("array is nil")
	}

	if a.data == nil {
		panic("array pointer is nil")
	}

	if index < 0 || index >= a.len {
		msg := fmt.Sprintf("index %d out of bounds (%d-%d):", index, 0, a.len-1)
		panic(msg)
	}
}

// Set устанавливает значение по индексу
func (a *Array[T]) Set(index int, value T) {
	a.checkBounds(index)

	valueAddr := a.addr(index)
	//nolint:govet,gosec // Надо
	*(*T)(unsafe.Pointer(valueAddr)) = value
}

func (a *Array[T]) addr(index int) uintptr {
	return uintptr(a.data) + uintptr(index)*unsafe.Sizeof(*new(T))
}

// Clear освобождает память, выделенную под массив
// Не забывайте вызывать этот метод, чтобы избежать утечек памяти
func (a *Array[T]) Clear() {
	if a.data != nil {
		a.allocator.free(a.data)
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
