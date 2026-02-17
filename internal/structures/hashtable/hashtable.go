package hashtable

import (
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/buckets"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/buckets/types"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/hashers"
)

const defaultSize = 10

type HashTable[T any] struct {
	size     int
	addrspcs []buckets.Bucket[T]
	hasher   Hasher
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
		hasher = hashers.NewMapHashHasher()
	}

	addrspcs := make([]buckets.Bucket[T], size)
	for i := range size {
		addrspcs[i] = types.NewChain[T]()
	}

	return &HashTable[T]{
		addrspcs: addrspcs,
		size:     size,
		hasher:   hasher,
	}
}

func (h *HashTable[T]) Set(key string, value T) {
	bucket, hash := h.hash(key)
	item := &buckets.Item[T]{
		Key:   key,
		Hash:  hash,
		Value: value,
	}
	h.addrspcs[bucket].Set(item)
}

// Get возвращает значение по ключу
// если ключ не найден, возвращает нулевое значение типа T
//
//nolint:ireturn // полиморфизм
func (h *HashTable[T]) Get(key string) T {
	var zero T

	bucket, hash := h.hash(key)

	addr := h.addrspcs[bucket]
	if addr == nil {
		return zero
	}

	item := addr.Get(hash, key)
	if item == nil {
		return zero
	}

	return item.GetValue()
}

func (h *HashTable[T]) Delete(key string) {
	bucket, hash := h.hash(key)

	addr := h.addrspcs[bucket]
	if addr == nil {
		return
	}

	addr.Delete(hash, key)
}

// hash возвращает нужное ведро и оригинальный хеш
// buckets - индекс ведра
// hash - оригинальный хеш
//
//nolint:nonamedreturns // имеет смысл
func (h *HashTable[T]) hash(key string) (bucket int, hash uint64) {
	hash = h.hasher.Hash(key)
	// G115: integer overflow conversion int -> uint64
	// И что делать?
	// Можно сделать проверку на переполнение, но это будет замедлять работу
	// Значит можно представить, что это нюанс?
	// Но в итоге то это int, а значит может быть отрицательное ведро?
	// Это нормально, так как uint64 % int всегда даст положительный результат
	//nolint:gosec // Логика работы хеш-таблицы предполагает такой каст. Номер ведра не может быть отрицательным.
	bucket = int(hash % uint64(h.size))

	return
}
