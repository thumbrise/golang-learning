package store

type Item[T any] interface {
	GetKey() string
	GetHash() uint64
	GetValue() T
	CompareKey(other Item[T]) bool
	Copy() Item[T]
	IsZero() bool
	IsWritable(other Item[T]) bool
}
