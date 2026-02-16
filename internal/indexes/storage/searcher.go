package storage

type Searcher interface {
	SearchEqual(fieldName string, value string) []int
	SearchRange(fieldName string, from string, to string) []int
	SearchPrefix(fieldName string, prefix string) []int
	SearchSuffix(fieldName string, suffix string) []int
	SearchContains(fieldName string, substring string) []int
	SearchIn(fieldName string, values []string) []int
}
