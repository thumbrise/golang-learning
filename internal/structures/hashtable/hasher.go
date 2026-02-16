package hashtable

import (
	"hash/maphash"
)

type Hasher interface {
	Hash(key string) uint64
}
type defaultHasher struct {
	seed maphash.Seed
}

func newDefaultHasher() *defaultHasher {
	return &defaultHasher{
		seed: maphash.MakeSeed(),
	}
}

func (h *defaultHasher) Hash(key string) uint64 {
	return maphash.String(h.seed, key)
}
