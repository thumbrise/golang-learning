package store

type ZeroItem[T any] struct{}

func (i *ZeroItem[T]) IsWritable(other Item[T]) bool {
	return true
}

func (i *ZeroItem[T]) GetKey() string {
	return ""
}

func (i *ZeroItem[T]) GetHash() uint64 {
	return 0
}

// GetValue возвращает значение типа T
//
//nolint:ireturn // Надо
func (i *ZeroItem[T]) GetValue() T {
	var zero T

	return zero
}

func (i *ZeroItem[T]) CompareKey(other Item[T]) bool {
	return false
}

func (i *ZeroItem[T]) Copy() Item[T] {
	return i
}

func (i *ZeroItem[T]) IsZero() bool {
	return true
}
