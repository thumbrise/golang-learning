package dal

import (
	"github.com/thumbrise/golang-learning/internal/indexes/indexes"
)

type UserStorage struct {
	data    map[int]User
	indexes map[string]indexes.Index
}

func NewUserStorage(data []User) *UserStorage {
	dataMap := make(map[int]User)
	for _, user := range data {
		dataMap[user.ID] = user
	}

	return &UserStorage{
		data:    dataMap,
		indexes: make(map[string]indexes.Index),
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
	result := make([]User, 0)

	// TODO: Нужен планировщик
	if len(s.indexes) > 0 {
		for _, index := range s.indexes {
			ctids := index.Search(field, value)
			// TODO: Использовать ctids для поиска пользователей
			for _, ctid := range ctids {
				result = append(result, s.data[ctid])
			}
		}

	} else {
		for _, user := range s.data {
			if user.Get(field) == value {
				result = append(result, user)
			}
		}
	}

	return result
}

func (s *UserStorage) SearchRange(field string, from string, to string) []User {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *UserStorage) SearchPrefix(field string, prefix string) []User {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *UserStorage) SearchSuffix(field string, suffix string) []User {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *UserStorage) SearchContains(field string, substring string) []User {
	// TODO: Использовать индекс если он есть
	return nil
}

func (s *UserStorage) SearchIn(field string, values []string) []User {
	// TODO: Использовать индекс если он есть
	return nil
}
