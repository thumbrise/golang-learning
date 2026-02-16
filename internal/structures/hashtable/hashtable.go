package hashtable

import "sync"

const defaultSize = 10

type HashTable[T any] struct {
	size   int
	data   [][]T
	hasher Hasher
	mu     sync.RWMutex
}

// NewHashTable создает новую хеш-таблицу
//
// size - размер хеш-таблицы
// hasher - структура реализующая интерфейс Hasher
//
// defaults:
//   - size: defaultSize
//   - hasher: nil (будет использоваться стандартная хеш-функция)
//

func NewHashTable[T any](size int, hasher Hasher) *HashTable[T] {
	if size == 0 {
		size = defaultSize
	}

	if hasher == nil {
		hasher = newDefaultHasher()
	}

	return &HashTable[T]{
		data:   make([][]T, size),
		size:   size,
		hasher: hasher,
	}
}

func (h *HashTable[T]) Set(key string, value T) {
	h.mu.Lock()
	defer h.mu.Unlock()
	hsh := h.hash(key)
	h.data[hsh] = append(h.data[hsh], value)
}

func (h *HashTable[T]) Get(key string) T {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var zero T

	hsh := h.hash(key)
	if h.data[hsh] == nil {
		return zero
	}

	return h.data[hsh][0]
}

func (h *HashTable[T]) hash(key string) int {
	return h.hasher.Hash(key) % h.size
}
