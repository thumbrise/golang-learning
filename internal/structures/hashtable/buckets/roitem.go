package buckets

type ROItem[T any] interface {
	GetKey() string
	GetHash() uint64
	GetValue() T
}
