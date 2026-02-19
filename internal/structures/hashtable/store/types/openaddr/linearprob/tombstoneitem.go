package linearprob

import "github.com/thumbrise/golang-learning/internal/structures/hashtable/store"

type TombstoneItem[T any] struct{}

func (i *TombstoneItem[T]) IsWritable(other store.Item[T]) bool {
	return true
}

func (i *TombstoneItem[T]) GetKey() string {
	return ""
}

func (i *TombstoneItem[T]) GetHash() uint64 {
	return 0
}

//nolint:ireturn // OK
func (i *TombstoneItem[T]) GetValue() T {
	var zero T

	return zero
}

func (i *TombstoneItem[T]) CompareKey(other store.Item[T]) bool {
	return false
}

func (i *TombstoneItem[T]) Copy() store.Item[T] {
	return i
}

func (i *TombstoneItem[T]) IsZero() bool {
	return false
}
