package indexes

type BTree struct {
	fields map[string]*BTreeField
}

func NewBTree(fields map[string]*BTreeField) *BTree {
	return &BTree{
		fields: fields,
	}
}

func (b *BTree) SearchEqual(field string, value string) []int {
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

type BTreeField struct {
	Values map[string]BTReeFieldValue
}
type BTReeFieldValue struct {
	CTIDs []int
}
