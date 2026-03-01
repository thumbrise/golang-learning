package search

type Indexer struct{}

func NewIndexer() *Indexer {
	return &Indexer{}
}

func (i *Indexer) CreateIndex(ctid string, fieldName string, fieldValue any, index Index) {
	vstr := i.parseString(fieldValue)
	if vstr != "" {
		i.index(ctid, fieldName, vstr, index)

		return
	}

	vslice := i.parseSliceOfStrings(fieldValue)
	if vslice != nil {
		for _, v := range vslice {
			i.index(ctid, fieldName, v, index)
		}

		return
	}
}

func (i *Indexer) index(ctid string, fieldName string, fieldValue string, index Index) {
	index.Insert(ctid, fieldName, fieldValue)
}

func (i *Indexer) parseString(value any) string {
	if v, ok := value.(string); ok {
		return v
	}

	return ""
}

func (i *Indexer) parseSliceOfStrings(value any) []string {
	if v, ok := value.([]string); ok {
		return v
	}

	return nil
}
