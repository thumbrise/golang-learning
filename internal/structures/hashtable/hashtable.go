package hashtable

import (
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/hashers"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store/types/chain"
)

// defaultSize = 5 * 2^10 = 5120 buckets = 40KB mean size
// Its ok for default?
const defaultSize = 5 << 10

type HashTable[T any] struct {
	store  store.Store[T]
	hasher Hasher
}

type StoreFactory[T any] func(size int) store.Store[T]

func defaultStoreFactory[T any](size int) store.Store[T] {
	return chain.NewStore[T](size)
}

// NewHashTable создает новую хеш-таблицу
//
// size - размер хеш-таблицы
// hasher - структура реализующая интерфейс Hasher
//
// defaults:
//   - size: defaultSize
//   - hasher: nil (будет использоваться стандартная хеш-функция)
func NewHashTable[T any](size int, hasher Hasher, storeFactory StoreFactory[T]) *HashTable[T] {
	if size == 0 {
		size = defaultSize
	}

	if hasher == nil {
		hasher = hashers.NewMapHashHasher()
	}

	if storeFactory == nil {
		storeFactory = defaultStoreFactory[T]
	}

	stor := storeFactory(size)
	if stor == nil {
		panic("store factory is nil")
	}

	return &HashTable[T]{
		store:  stor,
		hasher: hasher,
	}
}

func (h *HashTable[T]) Set(key string, value T) bool {
	hash := h.hash(key)

	item := &store.Item[T]{
		Key:   key,
		Hash:  hash,
		Value: value,
	}

	return h.store.Set(item)
}

// Get возвращает значение по ключу
// если ключ не найден, возвращает нулевое значение типа T
//
//nolint:ireturn // полиморфизм
func (h *HashTable[T]) Get(key string) T {
	hash := h.hash(key)

	item := h.store.Get(&store.Item[T]{
		Key:  key,
		Hash: hash,
	})

	return item.GetValue()
}

func (h *HashTable[T]) Delete(key string) bool {
	hash := h.hash(key)

	return h.store.Delete(&store.Item[T]{
		Key:  key,
		Hash: hash,
	})
}

// hash возвращает хеш
func (h *HashTable[T]) hash(key string) uint64 {
	return h.hasher.Hash(key)
}
