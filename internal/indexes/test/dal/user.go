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

func (u *User) GetString(field string) string {
	return u.Get(field).(string)
}

func (u *User) GetInt(field string) int {
	return u.Get(field).(int)
}

func (u *User) GetInt64(field string) int64 {
	return u.Get(field).(int64)
}

func (u *User) GetSlice(field string) []string {
	return u.Get(field).([]string)
}
