package hashers

// FirstRuneReturnHasher возвращает хеш, равный первому символу ключа
// Это плохой хешер, так как большинство ключей будут иметь одинаковый хеш
// Нужен для демонстрации большого фактора коллизии
type FirstRuneReturnHasher struct{}

func NewFirstRuneReturnHasher() *FirstRuneReturnHasher {
	return &FirstRuneReturnHasher{}
}

func (f FirstRuneReturnHasher) Hash(key string) uint64 {
	return uint64(key[0])
}
