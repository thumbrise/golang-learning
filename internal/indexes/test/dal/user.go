package dal

type User struct {
	ID             int
	Email          string
	Age            int
	FavoriteColors []string
	LastAccessTime int64
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               u.ID,
		"email":            u.Email,
		"age":              u.Age,
		"favorite_colors":  u.FavoriteColors,
		"last_access_time": u.LastAccessTime,
	}
}

func (u *User) Fields() []string {
	return []string{
		"id",
		"email",
		"age",
		"favorite_colors",
		"last_access_time",
	}
}

func (u *User) Get(field string) any {
	return u.ToMap()[field]
}
