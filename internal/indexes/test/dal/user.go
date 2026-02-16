package dal

type User struct {
	ID             int
	Email          string
	Age            int
	FavoriteColors []string
	LastAccessTime int64
}

func (u *User) DynamicFields() map[string]interface{} {
	return map[string]interface{}{
		"id":               u.ID,
		"email":            u.Email,
		"age":              u.Age,
		"favorite_colors":  u.FavoriteColors,
		"last_access_time": u.LastAccessTime,
	}
}
