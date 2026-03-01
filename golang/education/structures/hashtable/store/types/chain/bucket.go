package chain

import (
	"sync"

	"github.com/thumbrise/demo/golang/internal/structures/hashtable/store"
)

// Bucket - Адресное пространство, реализующее стратегию цепочек
type Bucket[T any] struct {
	mu    sync.RWMutex
	items []*store.HashedItem[T]
}

func NewBucket[T any]() *Bucket[T] {
	return &Bucket[T]{
		items: make([]*store.HashedItem[T], 0),
	}
}

func (h *Bucket[T]) Set(item store.Item[T]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	itemToSet := &store.HashedItem[T]{
		Hash:  item.GetHash(),
		Key:   item.GetKey(),
		Value: item.GetValue(),
	}

	for i, existingItem := range h.items {
		if existingItem.CompareKey(item) {
			h.items[i] = itemToSet

			return true
		}
	}

	h.items = append(h.items, itemToSet)

	return true
}

func (h *Bucket[T]) Get(item store.Item[T]) store.Item[T] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Полагаемся на метод Set. Он должен недопустить дубликатов ключей
	for _, existingItem := range h.items {
		if existingItem.CompareKey(item) {
			return existingItem
		}
	}

	return &store.ZeroItem[T]{}
}

func (h *Bucket[T]) Delete(item store.Item[T]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	for i, existingItem := range h.items {
		if existingItem.CompareKey(item) {
			h.items = append(h.items[:i], h.items[i+1:]...)

			return true
		}
	}

	return false
}
