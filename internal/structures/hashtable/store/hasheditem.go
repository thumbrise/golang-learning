package store

type HashedItem[T any] struct {
	Key   string
	Hash  uint64
	Value T
}

func (i *HashedItem[T]) IsWritable(other Item[T]) bool {
	return i.CompareKey(other)
}

func (i *HashedItem[T]) GetKey() string {
	return i.Key
}

func (i *HashedItem[T]) GetHash() uint64 {
	return i.Hash
}

//nolint:ireturn // OK
func (i *HashedItem[T]) GetValue() T {
	return i.Value
}

func (i *HashedItem[T]) CompareKey(other Item[T]) bool {
	return i.Hash == other.GetHash() && i.Key == other.GetKey()
}

func (i *HashedItem[T]) Copy() Item[T] {
	return &HashedItem[T]{
		Key:   i.Key,
		Hash:  i.Hash,
		Value: i.Value,
	}
}

func (i *HashedItem[T]) IsZero() bool {
	return false
}
