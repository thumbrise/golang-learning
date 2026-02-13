package indexes_test

import (
	"testing"

	"github.com/thumbrise/indexes/internal/indexes"
	"github.com/thumbrise/indexes/internal/indexes/test/fixtures"
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

func BuildBTree(users []fixtures.User, fieldName string) *indexes.BTree {
	values := map[string]indexes.BTReeFieldValue{}

	for i, user := range users {
		b3Value, ok := values[user.Email]
		if !ok {
			values[user.Email] = indexes.BTReeFieldValue{
				CTIDs: []int{i},
			}
		} else {
			b3Value.CTIDs = append(b3Value.CTIDs, i)
		}
	}

	b3fields := map[string]*indexes.BTreeField{
		fieldName: {
			Values: values,
		},
	}

	btree := indexes.NewBTree(b3fields)

	return btree
}

func prepareSearchables(users []fixtures.User, searchable fixtures.User) {
	searchUser := &users[len(users)-1]

	searchUser.Email = searchable.Email
}

func Benchmark_FindAll(b *testing.B) {
	const usersCount = 100000

	users := fixtures.GenerateTestUsers(usersCount)

	searchable := fixtures.User{
		Email:          "searchable@example.com",
		FavoriteColors: nil, // todo
		LastAccessTime: 0,   // todo
	}

	prepareSearchables(users, searchable)

	btree := BuildBTree(users, "Email")

	b.Run("Email", func(b *testing.B) {
		b.Run("Linear", func(b *testing.B) {
			for range b.N {
				foundedUsers := linearSearchEmail(users, searchable.Email)
				if len(foundedUsers) == 0 {
					b.Errorf("no results, case = %s, v = %s", "Email/Linear", searchable.Email)

					break
				}
			}
		})
		b.Run("BTree", func(b *testing.B) {
			var foundedCount int

			for range b.N {
				foundedUsers := btree.SearchEqual("Email", searchable.Email)
				if foundedCount = len(foundedUsers); foundedCount == 0 {
					b.Errorf("no results, case = %s, v = %s", "Email/BTree", searchable.Email)

					break
				}
			}

			b.StopTimer()
			b.Logf("Search results len %d", foundedCount)
		})
	})

	// 2.1 бенчмарк - поиск по FavoriteColors (без индекса)
	// Поиск по FavoriteColors
	// b.Run("Color/Linear", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		LinearSearchByColor(users, "searchable")
	//	}
	// })
	// 2.2 бенчмарк - поиск по FavoriteColors (GIN)
	// b.Run("Color/GIN", func(b *testing.B) {
	//	idx := BuildGIN(users, "FavoriteColors")
	//	b.ResetTimer()
	//	for i := 0; i < b.N; i++ {
	//		idx.Search("searchable")
	//	}
	// })

	// 3.1 бенчмарк - поиск по LastAccessTime (без индекса)
	// b.Run("Time/Linear", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		LinearSearchByTime(users, searchableLastAccessTime)
	//	}
	// })

	// 3.2 бенчмарк - поиск по LastAccessTime (BRIN)
	// b.Run("Time/BRIN", func(b *testing.B) {
	//	idx := BuildBRIN(users, "LastAccessTime")
	//	b.ResetTimer()
	//	for i := 0; i < b.N; i++ {
	//		idx.SearchRange(searchableLastAccessTime-1000, searchableLastAccessTime+1000)
	//	}
	// })
}
