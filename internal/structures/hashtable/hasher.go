package hashtable

type Hasher interface {
	Hash(key string) uint64
}
