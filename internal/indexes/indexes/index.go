package indexes

type Index interface {
	Insert(ctid int, fieldName string, value string)
	Search(fieldName string, value string) []int
	Delete(ctid int, fieldName string, value string)
	Update(ctid int, fieldName string, oldValue string, newValue string)
	SizeBytes() int
	Depth() int
	Stats() map[string]any
	Type() string
}
