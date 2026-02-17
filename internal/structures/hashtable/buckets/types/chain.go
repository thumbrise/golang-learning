package types

import (
	"sync"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable/buckets"
)

// Chain - Адресное пространство, реализующее стратегию цепочек
type Chain[T any] struct {
	mu    sync.RWMutex
	items []*buckets.Item[T]
}

func NewChain[T any]() *Chain[T] {
	return &Chain[T]{
		items: make([]*buckets.Item[T], 0),
	}
}

// Set добавляет item в адресное пространство
func (h *Chain[T]) Set(item *buckets.Item[T]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	for i, existing := range h.items {
		if existing.Hash == item.Hash && existing.Key == item.Key {
			h.items[i] = item

			return true
		}
	}

	h.items = append(h.items, item)

	return true
}

// Get возвращает item по hash и key
//
// Возвращает nil, если item не найден
//

func (h *Chain[T]) Get(hash uint64, key string) buckets.ROItem[T] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var zero *buckets.Item[T]

	// Полагаемся на метод Set. Он должен недопустить дубликатов ключей
	for _, item := range h.items {
		if item.Hash == hash && item.Key == key {
			return item
		}
	}

	return zero
}

// Delete удаляет item по hash и key
func (h *Chain[T]) Delete(hash uint64, key string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	for i, item := range h.items {
		if item.Hash == hash && item.Key == key {
			h.items = append(h.items[:i], h.items[i+1:]...)

			return true
		}
	}

	return false
}
