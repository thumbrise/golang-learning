package indexes_test

import (
	"testing"

	"github.com/thumbrise/golang-learning/internal/indexes/indexes/hash"
	"github.com/thumbrise/golang-learning/internal/indexes/test/fixtures"
)

func linearSearchEmail(users []fixtures.User, email string) []*fixtures.User {
	result := make([]*fixtures.User, 0)

	for i, user := range users {
		if user.Email == email {
			result = append(result, &users[i])
		}
	}

	return result
}

type Searcher interface {
	SearchEqual(fieldName string, value string) []int
	SearchRange(fieldName string, from string, to string) []int
	SearchPrefix(fieldName string, prefix string) []int
	SearchSuffix(fieldName string, suffix string) []int
	SearchContains(fieldName string, substring string) []int
	SearchIn(fieldName string, values []string) []int
}

func BuildHash(users []fixtures.User, fieldName string) *hash.Hash {
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

func prepareSearchables(users []fixtures.User, searchable fixtures.User) {
	searchUser := &users[len(users)-1]

	searchUser.Email = searchable.Email
}

func Benchmark_Search(b *testing.B) {
	const usersCount = 100000

	users := fixtures.GenerateTestUsers(usersCount)

	searchable := fixtures.User{
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

	// TODO FavoriteColors
	//  Linear
	//  Hash
	//  BTree
	// 	GIN

	// TODO LastAccessTime
	//  Linear
	//  Hash
	//  BTree
	//  BRIN
}
