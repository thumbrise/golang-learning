package hash

import "fmt"

type CTIDs = []int
type Values = map[string]CTIDs
type Keys = map[string]Values

type Hash struct {
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

func (h *Hash) Insert(ctid int, fieldName string, value string) {
	if h.table[fieldName] == nil {
		h.table[fieldName] = make(Values)
	}

	if h.table[fieldName][value] == nil {
		h.table[fieldName][value] = make(CTIDs, 0, 8)
	}

	h.table[fieldName][value] = append(h.table[fieldName][value], ctid)
}

func (h *Hash) Search(fieldName string, value string) []int {
	result := make([]int, 0)

	b3field, ok := h.table[fieldName]
	if !ok {
		return result
	}

	b3value, ok := b3field[value]
	if !ok {
		return result
	}

	return b3value
}

func (h *Hash) Delete(ctid int, fieldName string, value string) {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) Update(ctid int, fieldName string, oldValue string, newValue string) {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) SizeBytes() int {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) Depth() int {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) Stats() map[string]any {
	// TODO implement me
	panic("implement me")
}

func (h *Hash) String() string {
	return fmt.Sprintf("%T", h)
}
