package fixtures

import (
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/thumbrise/golang-learning/internal/indexes/test/dal"
)

func GenerateTestUsers(count int) []dal.User {
	users := make([]dal.User, count)
	for i := range count {
		users[i] = dal.User{
			ID:             i,
			Email:          faker.Email(),
			LastAccessTime: time.Now().UnixNano() + int64(i), // монотонно возрастают
		}
		users[i].FavoriteColors = GetRandomColors()
	}

	return users
}
