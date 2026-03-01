package core

type Record interface {
	PK() string
	ToMap() map[string]interface{}
	Fields() []string
	Get(field string) any
	GetString(field string) (string, error)
	GetInt(field string) (int, error)
	GetInt64(field string) (int64, error)
	GetSlice(field string) ([]string, error)
}
