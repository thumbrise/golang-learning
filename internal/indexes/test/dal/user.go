package dal

type User struct {
	ID             int
	Email          string
	Age            int
	FavoriteColors []string
	LastAccessTime int64
}
