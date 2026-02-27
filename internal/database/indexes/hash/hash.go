package hash

import (
	"fmt"
	"sync"

	"github.com/thumbrise/demo/internal/database/storage/stats"
)

type (
	CTIDs  = []string
	Values = map[string]CTIDs
	Keys   = map[string]Values
)

type Hash struct {
	mu    sync.RWMutex
	table Keys
}

func NewHash() *Hash {
	return &Hash{
		table: make(Keys),
	}
}

func (h *Hash) Type() string {
	return h.String()
}

func (h *Hash) Insert(ctid string, fieldName string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.table[fieldName] == nil {
		h.table[fieldName] = make(Values)
	}

	if h.table[fieldName][value] == nil {
		h.table[fieldName][value] = make(CTIDs, 0, 8)
	}

	h.table[fieldName][value] = append(h.table[fieldName][value], ctid)
}

func (h *Hash) Search(fieldName string, value string) []string {
	result := make([]string, 0)

	vals, ok := h.table[fieldName]
	if !ok {
		return result
	}

	v, ok := vals[value]
	if !ok {
		return result
	}

	return v
}

func (h *Hash) Delete(fieldName string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.table[fieldName] == nil {
		return
	}

	if h.table[fieldName][value] == nil {
		return
	}

	delete(h.table[fieldName], value)

	h.cleanup(fieldName, value)
}

// Full list scan on delete: Instead of marking deletions lazily or using a map for O(1) removal, Delete performs a full scan to filter out the target ctid. This implies performance degrades linearly with the number of duplicated values per key.
// TODO: Consider using a more efficient deletion strategy if performance becomes an issue.
func (h *Hash) DeleteCTID(ctid string, fieldName string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.table[fieldName] == nil {
		return
	}

	if h.table[fieldName][value] == nil {
		return
	}
	// Delete all occurrences of ctid, even if it appears multiple times
	original := h.table[fieldName][value]

	filtered := make([]string, 0, len(original))
	for _, ctidStored := range original {
		if ctidStored != ctid {
			filtered = append(filtered, ctidStored)
		}
	}

	h.table[fieldName][value] = filtered

	h.cleanup(fieldName, value)
}

func (h *Hash) Update(ctid string, fieldName string, oldValue string, newValue string) {
	h.DeleteCTID(ctid, fieldName, oldValue)
	h.Insert(ctid, fieldName, newValue)
}

func (h *Hash) SizeBytes() int {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) Depth() int {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) Stats() *stats.IndexStats {
	//nolint:gosec // Надо
	rows := uint32(len(h.table))

	return &stats.IndexStats{
		Cost: 1,
		Rows: rows,
	}
}

func (h *Hash) String() string {
	return fmt.Sprintf("%T", h)
}

func (h *Hash) cleanup(fieldName string, value string) {
	if h.table[fieldName] == nil {
		return
	}

	if h.table[fieldName][value] == nil {
		return
	}
	// Clean up empty slices
	if len(h.table[fieldName][value]) == 0 {
		delete(h.table[fieldName], value)
	}
	// Clean up empty maps
	if len(h.table[fieldName]) == 0 {
		delete(h.table, fieldName)
	}
}
