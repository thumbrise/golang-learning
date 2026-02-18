package hashtable_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/hashers"
)

// Почему delete работает? Он ведь недоделан??? Надо проверить
func TestHashTableConcurrent(t *testing.T) {
	t.Parallel()

	values := []string{}

	const count = 100000
	for i := range count {
		value := "value" + strconv.Itoa(i)
		values = append(values, value)
	}

	h := hashtable.NewHashTable[string](0, nil, nil)

	const key = "replaceable_key"
	for _, value := range values {
		h.Set(key, value)
	}

	want := "value" + strconv.Itoa(count-1)
	if got := h.Get(key); got != want {
		t.Errorf("Get() = %v, want %v", got, want)
	}

	// delete
	h.Delete(key)

	if got := h.Get(key); got != "" {
		t.Errorf("Get() after delete = %v, want %v", got, "")
	}
}

func TestHashTableDelete(t *testing.T) {
	t.Parallel()

	t.Run("Deleted item really removed", func(t *testing.T) {
		t.Parallel()

		h := hashtable.NewHashTable[string](0, nil, nil)

		const (
			key   = "key"
			value = "value"
			want  = ""
		)

		h.Set(key, value)

		if got := h.Get(key); got != value {
			t.Errorf("Get() = %v, want %v", got, value)
		}

		// delete
		h.Delete(key)

		if got := h.Get(key); got != "" {
			t.Errorf("Get() after delete = %v, want %v", got, want)
		}
	})

	t.Run("Other items still present in table", func(t *testing.T) {
		t.Parallel()

		h := hashtable.NewHashTable[string](0, nil, nil)

		const (
			key1 = "key1"
			key2 = "key2"
			key3 = "key3"
		)

		const (
			value1 = "value1"
			value2 = "value2"
			value3 = "value3"
		)

		h.Set(key1, value1)
		h.Set(key2, value2)
		h.Set(key3, value3)

		// Удаляем только key2
		h.Delete(key2)

		want1 := value1
		want2 := ""
		want3 := value3

		// Ожидаем, что key1 и key3 сохранились, а key2 удален
		if got := h.Get(key1); got != want1 {
			t.Errorf("Get() = %v, want %v", got, want1)
		}

		if got := h.Get(key2); got != want2 {
			t.Errorf("Get() after delete = %v, want %v", got, want2)
		}

		if got := h.Get(key3); got != want3 {
			t.Errorf("Get() = %v, want %v", got, want3)
		}
	})

	t.Run("Other items still present in same bucket", func(t *testing.T) {
		t.Parallel()

		// Создаем хеш-таблицу с фиксированным размером, чтобы все элементы попали в одно ведро
		h := hashtable.NewHashTable[string](1, nil, nil)

		const (
			key1 = "key1"
			key2 = "key2"
			key3 = "key3"
		)

		const (
			value1 = "value1"
			value2 = "value2"
			value3 = "value3"
		)

		h.Set(key1, value1)
		h.Set(key2, value2)
		h.Set(key3, value3)

		// Удаляем только key2
		h.Delete(key2)

		want1 := value1
		want2 := ""
		want3 := value3

		// Ожидаем, что key1 и key3 сохранились, а key2 удален
		if got := h.Get(key1); got != want1 {
			t.Errorf("Get() = %v, want %v", got, want1)
		}

		if got := h.Get(key2); got != want2 {
			t.Errorf("Get() after delete = %v, want %v", got, want2)
		}

		if got := h.Get(key3); got != want3 {
			t.Errorf("Get() = %v, want %v", got, want3)
		}
	})
}

func BenchmarkHashTable_InsertAfterFill(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	fills := []int{0, 10, 100, 1000, 10000, 100000}
	badHasher := hashers.NewFirstRuneReturnHasher()
	goodHasher := hashers.NewMapHashHasher()
	hshrs := []hashtable.Hasher{badHasher, goodHasher}
	strategies := []string{"chain", "open"}

	// Предгенерируем пул ключей (достаточно большой)
	const maxKeys = 200000

	keys := make([]string, maxKeys)
	for i := range keys {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	for _, size := range sizes {
		for _, fill := range fills {
			if fill >= maxKeys {
				continue
			}

			for _, hasher := range hshrs {
				for _, start := range strategies {
					name := fmt.Sprintf("size=%d/fill=%d/hasher=%T/collisionstrat=%sFake",
						size, fill, hasher, start)
					b.Run(name, func(b *testing.B) {
						// Создаём таблицу
						ht := hashtable.NewHashTable[string](size, hasher, nil)
						// Вставляем fill элементов (подготовка)
						for i := range fill {
							ht.Set(keys[i], "value")
						}
						// Новый ключ для вставки
						newKey := keys[fill] // гарантированно не использован

						b.ResetTimer()

						for range b.N {
							ht.Set(newKey, "value") // одна вставка
						}
					})
				}
			}
		}
	}
}
