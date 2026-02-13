package indexes

import (
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/thumbrise/indexes/internal/indexes/test"
)

type User struct {
	Name           string
	Email          string
	FavoriteColors []string
}

func Test(t *testing.T) {
	t.Parallel()

	var users [100]User
	for i := range users {
		users[i] = User{
			Name:  faker.Name(),
			Email: faker.Email(),
		}

		users[i].FavoriteColors = test.GetRandomColors()
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
