package buckets

type Item[T any] struct {
	Key   string
	Hash  uint64
	Value T
}
