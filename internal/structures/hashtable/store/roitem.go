package store

type ROItem[T any] interface {
	GetKey() string
	GetHash() uint64
	GetValue() T
	CompareKey(other ROItem[T]) bool
	Copy() ROItem[T]
	IsZero() bool
}
