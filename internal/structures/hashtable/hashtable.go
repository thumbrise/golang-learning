package hashtable

import "sync"

const defaultSize = 10

type BucketItem[T any] struct {
	key   string
	hash  uint64
	value T
}
type HashTable[T any] struct {
	size    int
	buckets [][]*BucketItem[T]
	hasher  Hasher
	mu      sync.RWMutex
}

// NewHashTable создает новую хеш-таблицу
//
// size - размер хеш-таблицы
// hasher - структура реализующая интерфейс Hasher
//
// defaults:
//   - size: defaultSize
//   - hasher: nil (будет использоваться стандартная хеш-функция)
func NewHashTable[T any](size int, hasher Hasher) *HashTable[T] {
	if size == 0 {
		size = defaultSize
	}

	if hasher == nil {
		hasher = newDefaultHasher()
	}

	return &HashTable[T]{
		buckets: make([][]*BucketItem[T], size),
		size:    size,
		hasher:  hasher,
	}
}

func (h *HashTable[T]) Set(key string, value T) {
	h.mu.Lock()
	defer h.mu.Unlock()

	bucket, hash := h.hash(key)
	item := &BucketItem[T]{
		key:   key,
		hash:  hash,
		value: value,
	}

	for i, existing := range h.buckets[bucket] {
		if existing.hash == hash && existing.key == key {
			h.buckets[bucket][i] = item

			return
		}
	}

	h.buckets[bucket] = append(h.buckets[bucket], item)
}

// Get возвращает значение по ключу
// если ключ не найден, возвращает нулевое значение типа T
//
//nolint:ireturn // полиморфизм
func (h *HashTable[T]) Get(key string) T {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var zero T

	bucket, hash := h.hash(key)
	if h.buckets[bucket] == nil {
		return zero
	}

	// Полагаемся на метод Set. Он должен недопустить дубликатов ключей
	for _, item := range h.buckets[bucket] {
		if item.hash == hash && item.key == key {
			return item.value
		}
	}

	return zero
}

func (h *HashTable[T]) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	bucket, hash := h.hash(key)
	if h.buckets[bucket] == nil {
		return
	}

	replacement := make([]*BucketItem[T], 0, len(h.buckets[bucket])-1)
	for i, item := range h.buckets[bucket] {
		if item.hash != hash || item.key != key {
			replacement = append(replacement, h.buckets[bucket][i])
		}
	}

	h.buckets[bucket] = replacement
}

// hash возвращает нужное ведро и оригинальный хеш
// bucket - индекс ведра
// hash - оригинальный хеш
//
//nolint:nonamedreturns // имеет смысл
func (h *HashTable[T]) hash(key string) (bucket int, hash uint64) {
	hash = h.hasher.Hash(key)
	bucket = int(hash % uint64(h.size))

	return
}
