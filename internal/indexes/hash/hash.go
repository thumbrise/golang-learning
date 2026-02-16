package hash

type Hash struct {
	fields map[string]*Field
}

func NewHash(fields map[string]*Field) *Hash {
	return &Hash{
		fields: fields,
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
