package search

type Index interface {
	Insert(ctid string, fieldName string, value string)
	Search(fieldName string, value string) []string
	Delete(fieldName string, value string)
	DeleteCTID(ctid string, fieldName string, value string)
	Update(ctid string, fieldName string, oldValue string, newValue string)
	SizeBytes() int
	Depth() int
	Stats() map[string]any
	Type() string
	String() string
}
