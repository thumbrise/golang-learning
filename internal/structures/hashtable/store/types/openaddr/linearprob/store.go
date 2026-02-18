// Package linearprob реализует стратегию линейного пробирования
package linearprob

import (
	"errors"
	"fmt"
	"sync"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store/types/chain"
)

type Store[T any] struct {
	items []store.Item[T]
	mu    sync.RWMutex
}

const (
	sizeMultiplier = 2
	defaultSize    = 5 << 10 // 5 * 2^10 = 5120
)

func NewStore[T any](size int) *Store[T] {
	if size == 0 {
		size = defaultSize
	}

	size *= sizeMultiplier

	items := make([]store.Item[T], size)
	for i := range size {
		items[i] = &store.ZeroItem[T]{}
	}

	return &Store[T]{
		items: items,
	}
}

var ErrNoSpace = errors.New("failed insert key: no free index even after grow")

func failInsert[T any](item store.Item[T], s *Store[T]) {
	err := fmt.Errorf("%w: key=%s hash=%d size=%d", ErrNoSpace, item.GetKey(), item.GetHash(), len(s.items))
	panic(err)
}

func (s *Store[T]) Set(item store.Item[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := findFreeIndex(item, s.items)

	if index == -1 {
		s.grow()

		index = findFreeIndex(item, s.items)
		if index == -1 {
			failInsert(item, s)
		}
	}

	s.items[index] = item.Copy()

	return true
}

func (s *Store[T]) Get(item store.Item[T]) store.Item[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	idx := findItemIndex(item, s.items)

	if idx == -1 {
		return &store.ZeroItem[T]{}
	}

	result := s.items[idx]
	if result.IsZero() {
		return &store.ZeroItem[T]{}
	}

	return result
}

func (s *Store[T]) Delete(item store.Item[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := findItemIndex(item, s.items)

	if idx == -1 {
		return false
	}

	if s.items[idx].IsZero() {
		return false
	}

	s.items[idx] = &chain.TombstoneItem[T]{}

	return true
}

// grow увеличивает размер хеш-таблицы в sizeMultiplier раз
// Расчитывает новые индексы для всех элементов и переносит их в новый массив
func (s *Store[T]) grow() {
	newSize := len(s.items) * sizeMultiplier

	newItems := make([]store.Item[T], newSize)
	for i := range newSize {
		newItems[i] = &store.ZeroItem[T]{}
	}

	for i := range s.items {
		if !s.items[i].IsZero() {
			newIndex := findFreeIndex(s.items[i], newItems)
			if newIndex == -1 {
				failInsert(s.items[i], s)
			}

			newItems[newIndex] = s.items[i]
		}
	}

	s.items = newItems
}

// Size возвращает размер хеш-таблицы
func (s *Store[T]) Size() int {
	// Не использовать эту функцию внутри других функций, которые уже захватывают блокировку
	// для избегания рекурсивных взаимных блокировок
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.items)
}

// FillFactor возвращает коэффициент заполнения хеш-таблицы
func (s *Store[T]) FillFactor() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filled := 0

	for i := range s.items {
		if !s.items[i].IsZero() {
			filled++
		}
	}

	return float64(filled) / float64(len(s.items))
}

// probFreeIndex возвращает индекс свободной ячейки в хеш-таблице учииывая алгоритм линейного пробирования
// Циклически проходит по хеш-таблице start -> s.size -> 0 -> start-1
func findFreeIndex[T any](insertable store.Item[T], items []store.Item[T]) int {
	start := hashToIndex(insertable.GetHash(), items)

	check := func(from, to int, items []store.Item[T]) int {
		for i := from; i < to; i++ {
			// От пользователя ожидается либо вставка нового элемента, либо очевидная перезапись по ключу
			if items[i].IsWritable(insertable) {
				return i
			}
		}

		return -1
	}

	if idx := check(start, len(items), items); idx != -1 {
		return idx
	}

	if idx := check(0, start, items); idx != -1 {
		return idx
	}

	return -1
}

// findItemIndex пытается найти элемент с помощью алгоритма линейного пробирования
func findItemIndex[T any](target store.Item[T], items []store.Item[T]) int {
	start := hashToIndex(target.GetHash(), items)

	check := func(from, to int, items []store.Item[T]) int {
		for i := from; i < to; i++ {
			if items[i].IsZero() {
				return -1
			}

			if items[i].CompareKey(target) {
				return i
			}
		}

		return -1
	}

	if idx := check(start, len(items), items); idx != -1 {
		return idx
	}

	if idx := check(0, start, items); idx != -1 {
		return idx
	}

	return -1
}

func hashToIndex[T any](hash uint64, items []store.Item[T]) int {
	//nolint:gosec // Логика работы хеш-таблицы предполагает такой каст. Номер ведра не может быть отрицательным.
	return int(hash % uint64(len(items)))
}
