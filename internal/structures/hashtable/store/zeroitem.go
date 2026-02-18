package store

type Zero[T any] struct{}

func (i *Zero[T]) GetKey() string {
	return ""
}

func (i *Zero[T]) GetHash() uint64 {
	return 0
}

//nolint:ireturn // OK
func (i *Zero[T]) GetValue() T {
	var zero T

	return zero
}

func (i *Zero[T]) CompareKey(other ROItem[T]) bool {
	return false
}

func (i *Zero[T]) Copy() ROItem[T] {
	return i
}

func (i *Zero[T]) IsZero() bool {
	return true
}
