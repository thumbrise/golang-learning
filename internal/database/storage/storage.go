package storage

import (
	"github.com/thumbrise/golang-learning/internal/database/storage/core"
	"github.com/thumbrise/golang-learning/internal/database/storage/planner"
	"github.com/thumbrise/golang-learning/internal/database/storage/search"
)

type Storage[TRecord core.Record] struct {
	heap    *core.Heap[TRecord]
	indexes map[string]search.Index
	planner *planner.Planner[TRecord]
}

func NewStorage[TRecord core.Record](data []TRecord) *Storage[TRecord] {
	heap := core.NewHeap(data)

	idxs := make(map[string]search.Index)
	seqSearcher := search.NewSeqSearcher(heap)
	idxs[seqSearcher.Type()] = seqSearcher

	return &Storage[TRecord]{
		heap:    heap,
		indexes: idxs,
		planner: planner.NewPlanner(heap, idxs),
	}
}

func (s *Storage[TRecord]) CreateIndex(field string, index search.Index) {
	indexer := search.NewIndexer()

	s.heap.Iterate(func(rec TRecord) {
		indexer.CreateIndex(rec.PK(), field, rec.Get(field), index)
	})

	s.indexes[index.Type()] = index
}

func (s *Storage[TRecord]) SearchEqual(field string, value string) []TRecord {
	conds := []search.Condition{{Field: field, Value: value, Op: search.OpEqual}}
	idx := s.planner.SuggestIndex(conds)

	ctids := idx.Search(field, value)

	result := s.records(ctids)

	return result
}

func (s *Storage[TRecord]) records(ctids []string) []TRecord {
	result := make([]TRecord, 0)
	for _, ctid := range ctids {
		result = append(result, s.heap.Get(ctid))
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
