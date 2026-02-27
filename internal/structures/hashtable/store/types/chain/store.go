// Package chain реализует стратегию цепочек
package chain

import (
	"github.com/thumbrise/demo/internal/structures/hashtable/store"
)

const defaultSize = 5 << 10 // 5120

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

func (s *Store[T]) Set(item store.Item[T]) bool {
	bucket := s.bucket(item.GetHash())

	addr := s.buckets[bucket]
	if addr == nil {
		return false
	}

	return addr.Set(item)
}

func (s *Store[T]) Get(item store.Item[T]) store.Item[T] {
	hash := item.GetHash()
	bucket := s.bucket(hash)

	addr := s.buckets[bucket]
	if addr == nil {
		return &store.ZeroItem[T]{}
	}

	return addr.Get(item)
}

func (s *Store[T]) Delete(item store.Item[T]) bool {
	bucket := s.bucket(item.GetHash())

	addr := s.buckets[bucket]
	if addr == nil {
		return false
	}

	return addr.Delete(item)
}

func (s *Store[T]) bucket(hash uint64) int {
	//nolint:gosec // Логика работы хеш-таблицы предполагает такой каст. Неотрицательный хеш и остаток от деления будут положительными. Int в хеш таблице используется для упрощения работы
	return int(hash % uint64(s.size))
}
