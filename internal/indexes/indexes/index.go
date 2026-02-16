package indexes

type Index interface {
	Insert(fieldName string, value string, ctid int)
	Search(fieldName string, value string) []int
	Delete(fieldName string, value string, ctid int)
	Update(fieldName string, oldValue string, newValue string, ctid int)
	SizeBytes() int
	Depth() int
	Stats() map[string]any
	Type() string
}
