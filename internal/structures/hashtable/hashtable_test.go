package hashtable_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/hashers"
)

func TestHashTableBasic(t *testing.T) {
	t.Parallel()

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic_1", args{"key1", "value1"}, "value1"},
		{"basic_2", args{"key2", "value2"}, "value2"},
		{"basic_3", args{"key3", "value3"}, "value3"},
	}

	for _, tt := range tests {
		t.Run("concurrent", func(t *testing.T) {
			t.Parallel()

			h := hashtable.NewHashTable[string](0, nil, nil)

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				h.Set(tt.args.key, tt.args.value)

				if got := h.Get(tt.args.key); got != tt.want {
					t.Errorf("Get() = %v, want %v", got, tt.want)
				}
			})
		})
		t.Run("sequential", func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, nil)
				h.Set(tt.args.key, tt.args.value)

				if got := h.Get(tt.args.key); got != tt.want {
					t.Errorf("Get() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}

func TestHashTableReplace(t *testing.T) {
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
}

func TestHashTableCollision(t *testing.T) {
	t.Parallel()
	// ht := NewHashTable()
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
