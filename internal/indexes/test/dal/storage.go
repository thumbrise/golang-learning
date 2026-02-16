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

func (s *UserStorage) CreateIndex(field string, index indexes.Index) {
	indexer := NewIndexer()
	for _, user := range s.data {
		indexer.CreateIndex(user.ID, field, user.Get(field), index)
	}
	s.indexes[index.Type()] = index
}

func (s *UserStorage) SearchEqual(field string, value string) []User {
	return s.linearSearch(field, value)
}

func (s *UserStorage) SearchRange(field string, from string, to string) []User {
	return s.linearSearch(field, from)
}

func (s *UserStorage) SearchPrefix(field string, prefix string) []User {
	return s.linearSearch(field, prefix)
}

func (s *UserStorage) SearchSuffix(field string, suffix string) []User {
	return s.linearSearch(field, suffix)
}

func (s *UserStorage) SearchContains(field string, substring string) []User {
	return s.linearSearch(field, substring)
}

func (s *UserStorage) SearchIn(field string, values []string) []User {
	return s.linearSearch(field, values[0]) // Simplified for now
}

func (s *UserStorage) linearSearch(field string, value string) []User {
	result := make([]User, 0)

	for _, user := range s.data {
		if user.Get(field) == value {
			result = append(result, user)
		}
	}

	return result
}
