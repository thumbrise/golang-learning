package hash

import "fmt"

type Hash struct {
	fields map[string]*Field
}

func (b *Hash) Type() string {
	return b.String()
}

func (b *Hash) Insert(ctid int, fieldName string, value string) {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) Search(fieldName string, value string) []int {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) Delete(ctid int, fieldName string, value string) {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) Update(ctid int, fieldName string, oldValue string, newValue string) {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) SizeBytes() int {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) Depth() int {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) Stats() map[string]any {
	// TODO implement me
	panic("implement me")
}

func (b *Hash) String() string {
	return fmt.Sprintf("%T", b)
}

func NewHash() *Hash {
	return &Hash{
		fields: make(map[string]*Field),
	}
}

func (b *Hash) SearchEqual(field string, value string) []int {
	result := make([]int, 0)

	b3field, ok := b.fields[field]
	if !ok {
		return result
	}

	b3value, ok := b3field.Values[value]
	if !ok {
		return result
	}

	return b3value.CTIDs
}

type Field struct {
	Values map[string]FieldValue
}
type FieldValue struct {
	CTIDs []int
}
