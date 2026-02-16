package dal

import (
	"errors"
	"fmt"
)

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

var ErrCantCast = errors.New("can't cast")

func (u *User) GetString(field string) (string, error) {
	v, ok := u.Get(field).(string)
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrCantCast, field)
	}

	return v, nil
}

func (u *User) GetInt(field string) (int, error) {
	v, ok := u.Get(field).(int)
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrCantCast, field)
	}

	return v, nil
}

func (u *User) GetInt64(field string) (int64, error) {
	v, ok := u.Get(field).(int64)
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrCantCast, field)
	}

	return v, nil
}

func (u *User) GetSlice(field string) ([]string, error) {
	v, ok := u.Get(field).([]string)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrCantCast, field)
	}

	return v, nil
}
