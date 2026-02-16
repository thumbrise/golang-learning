package fixtures

import "github.com/thumbrise/golang-learning/internal/indexes/indexes"

type UserStorage struct {
	data    []User
	indexes []indexes.Index
}

func NewUserStorage(data []User, indexes []indexes.Index) *UserStorage {
	return &UserStorage{
		data:    data,
		indexes: indexes,
	}
}

func (s *UserStorage) CreateIndex(fieldName string) map[string][]int {
	return nil
}

func (s *UserStorage) SearchEqual(fieldName string, value string) []User {
	return nil
}

func (s *UserStorage) SearchRange(fieldName string, from string, to string) []User {
	return nil
}

func (s *UserStorage) SearchPrefix(fieldName string, prefix string) []User {
	return nil
}

func (s *UserStorage) SearchSuffix(fieldName string, suffix string) []User {
	return nil
}

func (s *UserStorage) SearchContains(fieldName string, substring string) []User {
	return nil
}

func (s *UserStorage) SearchIn(fieldName string, values []string) []User {
	return nil
}
