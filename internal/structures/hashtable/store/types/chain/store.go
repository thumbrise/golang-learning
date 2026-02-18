// Package chain реализует стратегию цепочек
package chain

import (
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
)

// defaultSize = 5 * 2^10 = 5120 buckets = 40KB mean size
// Its ok for default?
const defaultSize = 5 << 10

type Store[T any] struct {
	size    int
	buckets []*Bucket[T]
}

func NewStore[T any](size int) *Store[T] {
	if size == 0 {
		size = defaultSize
	}
	buckets := make([]*Bucket[T], size)
	for i := range size {
		buckets[i] = NewBucket[T]()
	}

	return &Store[T]{
		buckets: buckets,
		size:    size,
	}
}

// Set добавляет item в адресное пространство
func (s *Store[T]) Set(item store.ROItem[T]) bool {
	bucket := s.bucket(item.GetHash())

	addr := s.buckets[bucket]
	if addr == nil {
		return false
	}

	return addr.Set(item)
}

// Get возвращает item по bucket и key
//
// Возвращает nil, если item не найден
func (s *Store[T]) Get(item store.ROItem[T]) store.ROItem[T] {
	hash := item.GetHash()
	bucket := s.bucket(hash)

	addr := s.buckets[bucket]
	if addr == nil {
		return &store.Zero[T]{}
	}

	return addr.Get(item)
}

// Delete удаляет item по bucket и key
func (s *Store[T]) Delete(item store.ROItem[T]) bool {
	bucket := s.bucket(item.GetHash())

	addr := s.buckets[bucket]
	if addr == nil {
		return false
	}

	return addr.Delete(item)
}

// bucket возвращает нужное ведро и оригинальный хеш
// store - индекс ведра
// bucket - оригинальный хеш
//

func (s *Store[T]) bucket(hash uint64) int {
	// G115: integer overflow conversion int -> uint64
	// И что делать?
	// Можно сделать проверку на переполнение, но это будет замедлять работу.
	// Значит можно представить, что это нюанс?
	// Но в итоге то это int, а значит может быть отрицательное ведро?
	// Это нормально, так как uint64 % int всегда даст положительный результат
	//nolint:gosec // Логика работы хеш-таблицы предполагает такой каст. Номер ведра не может быть отрицательным.
	return int(hash % uint64(s.size))
}
