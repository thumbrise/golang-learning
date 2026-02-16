package hashtable

type Hasher interface {
	Hash(key string) int
}
type defaultHasher struct{}

func newDefaultHasher() *defaultHasher {
	return &defaultHasher{}
}

func (h *defaultHasher) Hash(key string) int {
	return int(key[0])
}
