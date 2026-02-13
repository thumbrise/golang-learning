package fixtures

import (
	"time"

	"github.com/go-faker/faker/v4"
)

type User struct {
	Email          string
	FavoriteColors []string
	LastAccessTime int64
}

func GenerateTestUsers(count int) []User {
	users := make([]User, count)
	for i := range count {
		users[i] = User{
			Email:          faker.Email(),
			LastAccessTime: time.Now().UnixNano() + int64(i), // монотонно возрастают
		}
		users[i].FavoriteColors = GetRandomColors()
	}

	return users
}
