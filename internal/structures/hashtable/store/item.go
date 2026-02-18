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

func (i *Item[T]) Compare(other ROItem[T]) bool {
	return i.Hash == other.GetHash() && i.Key == other.GetKey()
}

func (i *Item[T]) Copy() ROItem[T] {
	return &Item[T]{
		Key:   i.Key,
		Hash:  i.Hash,
		Value: i.Value,
	}
}
func (i *Item[T]) IsZero() bool {
	return false
}
