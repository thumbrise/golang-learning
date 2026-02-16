package indexes_test

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/thumbrise/golang-learning/internal/indexes/indexes/hash"
	"github.com/thumbrise/golang-learning/internal/indexes/test/dal"
	"github.com/thumbrise/golang-learning/internal/indexes/test/fixtures"
)

func linearSearchEmail(users []dal.User, email string) []*dal.User {
	result := make([]*dal.User, 0)

	for i, user := range users {
		if user.Email == email {
			result = append(result, &users[i])
		}
	}

	return result
}

func BuildHash(users []dal.User, fieldName string) *hash.Hash {
	values := map[string]hash.FieldValue{}

	for i, user := range users {
		b3Value, ok := values[user.Email]
		if !ok {
			values[user.Email] = hash.FieldValue{
				CTIDs: []int{i},
			}
		} else {
			b3Value.CTIDs = append(b3Value.CTIDs, i)
		}
	}

	b3fields := map[string]*hash.Field{
		fieldName: {
			Values: values,
		},
	}

	btree := hash.NewHash(b3fields)

	return btree
}

func prepareSearchables(users []dal.User, searchable dal.User) {
	searchUser := &users[len(users)-1]

	searchUser.Email = searchable.Email
}

func Benchmark_Search(b *testing.B) {
	faker.SetRandomSource(rand.NewSource(12345))

	const usersCount = 100000
	// TODO: Матрица состоящая из полей, индексов, аргументов поиска, тому подобное
	//  Для каждого варианта:
	//  - измерять время выполнения
	//  - измерять память
	//  - измерять количество операций чтения/записи
	//  - измерять количество итераций
	//  - измерять количество сравнений
	users := fixtures.GenerateTestUsers(usersCount)

	searchable := dal.User{
		Email:          "searchable@example.com",
		Age:            0,   // todo
		FavoriteColors: nil, // todo
		LastAccessTime: 0,   // todo
	}

	prepareSearchables(users, searchable)

	hsh := BuildHash(users, "Email")

	b.Run("Email", func(b *testing.B) {
		b.Run("Linear", func(b *testing.B) {
			for range b.N {
				foundedUsers := linearSearchEmail(users, searchable.Email)
				if len(foundedUsers) == 0 {
					b.Errorf("no results, case = %s, v = %s", "Email/Linear", searchable.Email)

					return
				}
			}
		})
		b.Run("Hash", func(b *testing.B) {
			for range b.N {
				foundedUsers := hsh.SearchEqual("Email", searchable.Email)
				if len(foundedUsers) == 0 {
					b.Errorf("no results, case = %s, v = %s", "Email/Hash", searchable.Email)

					return
				}
			}
		})
		// TODO BTree
	})

	// TODO Age
	//  Linear
	//  Hash
	//  BTree
	//  GIN
	//  BRIN

	// TODO FavoriteColors
	//  Linear
	//  Hash
	//  BTree
	//  GIN

	// TODO LastAccessTime
	//  Linear
	//  Hash
	//  BTree
	//  BRIN
}
