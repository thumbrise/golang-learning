package store

type Item[T any] struct {
	Key   string
	Hash  uint64
	Value T
}

func (i *Item[T]) GetKey() string {
	return i.Key
}

func (i *Item[T]) GetHash() uint64 {
	return i.Hash
}

//nolint:ireturn // OK
func (i *Item[T]) GetValue() T {
	return i.Value
}
