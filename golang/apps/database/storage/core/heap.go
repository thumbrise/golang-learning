package core

type Heap[TRecord Record] struct {
	data map[string]TRecord
}

func NewHeap[TRecord Record](records []TRecord) *Heap[TRecord] {
	dataMap := make(map[string]TRecord)
	for _, rec := range records {
		dataMap[rec.PK()] = rec
	}

	return &Heap[TRecord]{
		data: dataMap,
	}
}

type IterateFunc[TRecord Record] func(record TRecord)

func (h *Heap[TRecord]) Iterate(f IterateFunc[TRecord]) {
	for _, record := range h.data {
		f(record)
	}
}

// Get returns record by ctid
//
//nolint:ireturn //generic
func (h *Heap[TRecord]) Get(ctid string) TRecord {
	return h.data[ctid]
}

func (h *Heap[TRecord]) Len() uint32 {
	//nolint:gosec // Надо
	return uint32(len(h.data))
}
