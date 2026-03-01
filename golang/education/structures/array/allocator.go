package array

import (
	"log/slog"
	"reflect"
	"unsafe"
)

//#include <stdlib.h>
import "C"

type Allocator[T any] struct{}

func NewAllocator[T any]() *Allocator[T] {
	return &Allocator[T]{}
}

func (a *Allocator[T]) Alloc(length int) unsafe.Pointer {
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

func (a *Allocator[T]) free(ptr unsafe.Pointer) {
	C.free(ptr)
}
