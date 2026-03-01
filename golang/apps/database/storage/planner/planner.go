package planner

import (
	"github.com/thumbrise/demo/golang/internal/database/storage/core"
	"github.com/thumbrise/demo/golang/internal/database/storage/search"
)

type Planner[TR core.Record] struct {
	indexes map[string]search.Index
	heap    *core.Heap[TR]
}

func NewPlanner[TRecord core.Record](heap *core.Heap[TRecord], indexes map[string]search.Index) *Planner[TRecord] {
	return &Planner[TRecord]{
		heap:    heap,
		indexes: indexes,
	}
}

// SuggestIndex выбирает лучший индекс на основе стоимости
//
//nolint:ireturn //matrix polymorphism
func (p *Planner[TR]) SuggestIndex(conditions []search.Condition) search.Index {
	var (
		cheapest search.Index
		minCost  uint32 = 0
	)

	for _, index := range p.indexes {
		analysis := p.analyze(index, conditions)
		if cheapest == nil || analysis.Cost < minCost {
			cheapest = index
			minCost = analysis.Cost
		}
	}

	return cheapest
}

// analyze анализирует индекс и возвращает его стоимость
func (p *Planner[TR]) analyze(index search.Index, conditions []search.Condition) Analysis {
	st := index.Stats()

	return NewAnalysis(
		st.Cost,
		st.Rows,
	)
}
