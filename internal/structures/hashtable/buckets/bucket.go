package buckets

// Bucket представляет адресное пространство хранения для хэш-таблицы.
//
// Это контейнер, который сам знает как управлять своими элементами.
// Таким образом мы можем использовать разные стратегии хранения.
// И даже в рамках одной и той же таблицы использовать разные реализации пространств.
//
// !!! Каждый реализатор должен сам обрабатывать race conditions, иметь свои собственные mutex-ы. Sharded locking.
type Bucket[T any] interface {
	Set(item *Item[T]) bool
	Get(hash uint64, key string) ROItem[T]
	Delete(hash uint64, key string) bool
}
