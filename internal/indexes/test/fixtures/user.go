package fixtures

import (
	"time"

	"github.com/go-faker/faker/v4"
)

type User struct {
	ID             int
	Email          string
	Age            int
	FavoriteColors []string
	LastAccessTime int64
}

func GenerateTestUsers(count int) []User {
	users := make([]User, count)
	for i := range count {
		users[i] = User{
			ID:             i,
			Email:          faker.Email(),
			LastAccessTime: time.Now().UnixNano() + int64(i), // монотонно возрастают
		}
		users[i].FavoriteColors = GetRandomColors()
	}

	return users
}

// CREATE INDEX users(LAST_ACCESS_IP) USING BRIN;

// AGE = 30 - 98
// CHILDREN_COUNT = 2

// |[AGE] -> * [CHILDREN_COUNT]
// [root] -> [
//			[
//				2 -> [3, 4, 5, 6],
//				7 -> [8, 9, 10, 11],
//				12 -> [13, 14, 15, 16],
//			]
//	]
//
// [CHILDREN_COUNT] -> * [AGE]
