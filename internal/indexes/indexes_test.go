package indexes_test

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/thumbrise/golang-learning/internal/indexes/indexes"
	"github.com/thumbrise/golang-learning/internal/indexes/indexes/hash"
	"github.com/thumbrise/golang-learning/internal/indexes/test/dal"
	"github.com/thumbrise/golang-learning/internal/indexes/test/fixtures"
)

//nolint:ireturn //matrix polymorphism
func BuildHash() indexes.Index {
	return hash.NewHash()
}

func prepareSearchables(users []dal.User, searchable dal.User) {
	searchUser := &users[len(users)-1]

	searchUser.Email = searchable.Email
}

func Benchmark_Search(b *testing.B) {
	faker.SetRandomSource(rand.NewSource(12345))

	const usersCount = 100000

	testFields := (&dal.User{}).Fields()
	testIndexes := []func() indexes.Index{
		nil,
		BuildHash,
		// TODO: BTree
		// TODO: GIN
		// TODO: BRIN
	}

	searchable := dal.User{
		Email:          "searchable@example.com",
		Age:            0,   // todo
		FavoriteColors: nil, // todo
		LastAccessTime: 0,   // todo
	}

	// TODO:
	//  - измерять время выполнения
	//  - измерять память
	//  - измерять количество операций чтения/записи
	//  - измерять количество итераций
	//  - измерять количество сравнений
	for _, testField := range testFields {
		for _, testIndex := range testIndexes {
			b.Run(testField, func(b *testing.B) {
				users := fixtures.GenerateTestUsers(usersCount)
				prepareSearchables(users, searchable)
				storage := dal.NewUserStorage(users)
				idxType := "Linear"

				if testIndex != nil {
					idx := testIndex()
					idxType = idx.String()
					storage.CreateIndex(testField, idx)
				}

				b.Run(idxType, func(b *testing.B) {
					b.Run("SearchEqual", func(b *testing.B) {
						for range b.N {
							v, err := searchable.GetString(testField)
							if err != nil {
								b.Fatalf("failed to get string: %v", err)
							}

							results := storage.SearchEqual(testField, v)
							if len(results) == 0 {
								b.Errorf("no results v = %s", v)

								return
							}
						}
					})
				})
				// TODO BTree
			})
		}
	}

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
