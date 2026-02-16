package dal

import (
	"sync"

	"github.com/thumbrise/golang-learning/internal/indexes/indexes"
)

type UserStorage struct {
	data    map[int]User
	mu      sync.RWMutex
	indexes map[string]indexes.Index
}

func NewUserStorage(data []User) *UserStorage {
	dataMap := make(map[int]User)
	for _, user := range data {
		dataMap[user.ID] = user
	}
	return &UserStorage{
		data: dataMap,
	}
}

func (s *UserStorage) CreateIndex(fieldName string, index indexes.Index) {
	indexer := NewIndexer()
	for _, user := range s.data {
		f := user.DynamicFields()[fieldName]
		indexer.CreateIndex(user.ID, fieldName, f, index)
	}
	s.indexes[index.Type()] = index
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
