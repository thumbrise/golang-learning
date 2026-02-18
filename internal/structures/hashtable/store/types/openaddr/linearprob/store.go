// Package linearprob реализует стратегию линейного пробирования
package linearprob

import (
	"errors"
	"fmt"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
)

type Store[T any] struct {
	size  int
	items []store.ROItem[T]
}

const sizeMultiplier = 2
const defaultSize = 5 << 10 // 5 * 2^10 = 5120

func NewStore[T any](size int) *Store[T] {
	if size == 0 {
		size = defaultSize
	}
	size *= sizeMultiplier

	items := make([]store.ROItem[T], size)
	for i := range size {
		items[i] = &store.Zero[T]{}
	}

	return &Store[T]{
		items: items,
		size:  size,
	}
}

var ErrNoFreeIndex = errors.New("no free index")

// Set добавляет item в адресное пространство
func (s *Store[T]) Set(item store.ROItem[T]) bool {
	index := s.probFreeIndex(item.GetHash())

	if index == -1 {
		err := fmt.Errorf("%w: key=%s hash=%d size=%d", ErrNoFreeIndex, item.GetKey(), item.GetHash(), s.size)
		panic(err)
	}

	s.items[index] = item.Copy()
	return true
}

func (s *Store[T]) Get(item store.ROItem[T]) store.ROItem[T] {
	idx := s.probItemIndex(item)

	if idx == -1 {
		return &store.Zero[T]{}
	}

	addr := s.items[idx]
	if s.isZero(addr) {
		return &store.Zero[T]{}
	}

	return addr
}

// Delete удаляет item по bucket и key
func (s *Store[T]) Delete(item store.ROItem[T]) bool {
	idx := s.probItemIndex(item)

	if idx == -1 {
		return false
	}

	if s.isZero(s.items[idx]) {
		return false
	}

	s.items[idx] = &store.Zero[T]{}

	return true
}

// probFreeIndex возвращает индекс свободной ячейки в хеш-таблице учииывая алгоритм линейного пробирования
// Циклически проходит по хеш-таблице start -> s.size -> 0 -> start-1
func (s *Store[T]) probFreeIndex(hash uint64) int {
	start := s.hashToIndex(hash)

	for i := start; i < s.size; i++ {
		if s.isZero(s.items[i]) {
			return i
		}
	}

	for i := 0; i < start; i++ {
		if s.isZero(s.items[i]) {
			return i
		}
	}

	return -1
}

// probItemIndex пытается найти элемент с помощью алгоритма линейного пробирования
func (s *Store[T]) probItemIndex(item store.ROItem[T]) int {
	start := s.hashToIndex(item.GetHash())

	for i := start; i < s.size; i++ {
		if s.isZero(s.items[i]) {
			return -1
		}

		if s.items[i].Compare(item) {
			return i
		}
	}

	for i := 0; i < start; i++ {
		if s.isZero(s.items[i]) {
			return -1
		}

		if s.items[i].Compare(item) {
			return i
		}
	}

	return -1
}

func (s *Store[T]) hashToIndex(hash uint64) int {
	//nolint:gosec // Логика работы хеш-таблицы предполагает такой каст. Номер ведра не может быть отрицательным.
	return int(hash % uint64(s.size))
}

func (s *Store[T]) isZero(item store.ROItem[T]) bool {
	return item == nil || item.IsZero()
}
