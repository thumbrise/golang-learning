package hashers

import "hash/maphash"

// MapHashHasher использует стандартный хешер из пакета maphash
// Это хороший хешер, так как он распределяет ключи равномерно
// Best Practice для использования в production
type MapHashHasher struct {
	seed maphash.Seed
}

func NewMapHashHasher() *MapHashHasher {
	return &MapHashHasher{
		seed: maphash.MakeSeed(),
	}
}

func (h *MapHashHasher) Hash(key string) uint64 {
	return maphash.String(h.seed, key)
}
