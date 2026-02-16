package storage

import (
	"github.com/thumbrise/golang-learning/internal/search/indexes"
)

type Storage[TRecord Record] struct {
	heap    *Heap[TRecord]
	indexes map[string]indexes.Index
}

func NewStorage[TRecord Record](data []TRecord) *Storage[TRecord] {
	heap := NewHeap(data)

	return &Storage[TRecord]{
		heap:    heap,
		indexes: make(map[string]indexes.Index),
	}
}

func (s *Storage[TRecord]) CreateIndex(field string, index indexes.Index) {
	indexer := NewIndexer()

	s.heap.Iterate(func(rec TRecord) {
		indexer.CreateIndex(rec.PK(), field, rec.Get(field), index)
	})

	s.indexes[index.Type()] = index
}

func (s *Storage[TRecord]) SearchEqual(field string, value string) []TRecord {
	result := make([]TRecord, 0)

	// TODO: Нужен планировщик
	if len(s.indexes) > 0 {
		for _, index := range s.indexes {
			ctids := index.Search(field, value)
			for _, ctid := range ctids {
				result = append(result, s.heap.Get(ctid))
			}
		}
	} else {
		s.heap.Iterate(func(rec TRecord) {
			if rec.Get(field) == value {
				result = append(result, rec)
			}
		})
	}

	return result
}

func (s *Storage[TRecord]) SearchRange(field string, from string, to string) []TRecord {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *Storage[TRecord]) SearchPrefix(field string, prefix string) []TRecord {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *Storage[TRecord]) SearchSuffix(field string, suffix string) []TRecord {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *Storage[TRecord]) SearchContains(field string, substring string) []TRecord {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *Storage[TRecord]) SearchIn(field string, values []string) []TRecord {
	// TODO: Использовать индекс если он есть
	return nil
}
