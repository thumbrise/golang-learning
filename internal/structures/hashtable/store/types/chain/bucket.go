package chain

import (
	"sync"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
)

// Bucket - Адресное пространство, реализующее стратегию цепочек
type Bucket[T any] struct {
	mu    sync.RWMutex
	items []*store.Item[T]
}

func NewBucket[T any]() *Bucket[T] {
	return &Bucket[T]{
		items: make([]*store.Item[T], 0),
	}
}

// Set добавляет item в адресное пространство
func (h *Bucket[T]) Set(item store.ROItem[T]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	itemToSet := &store.Item[T]{
		Hash:  item.GetHash(),
		Key:   item.GetKey(),
		Value: item.GetValue(),
	}

	for i, existing := range h.items {
		if existing.Hash == item.GetHash() && existing.Key == item.GetKey() {
			h.items[i] = itemToSet

			return true
		}
	}

	h.items = append(h.items, itemToSet)

	return true
}

// Get возвращает item по bucket и key
//
// Возвращает nil, если item не найден
//

func (h *Bucket[T]) Get(item store.ROItem[T]) store.ROItem[T] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Полагаемся на метод Set. Он должен недопустить дубликатов ключей
	for _, item := range h.items {
		if item.Hash == item.GetHash() && item.Key == item.GetKey() {
			return item
		}
	}

	return &store.Zero[T]{}
}

// Delete удаляет item по bucket и key
func (h *Bucket[T]) Delete(item store.ROItem[T]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	// TODO: Item затеняется. Но это не мешает?
	//  Так как мы не используем item в цикле
	//  А зачем мы тогда приняли item в параметрах?
	//  Возможно, это сделано для единообразия с другими методами?
	//  Нет, мы должны сравнивать existingItem == passedItem
	//  А щас удаляются АБСОЛЮТНО все данные из ведра, критический ужасный баг.
	//  Нужно составить тест на этот случай.
	for i, item := range h.items {
		if item.Hash == item.GetHash() && item.Key == item.GetKey() {
			h.items = append(h.items[:i], h.items[i+1:]...)

			return true
		}
	}

	return false
}
