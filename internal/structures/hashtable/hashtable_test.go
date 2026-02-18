package hashtable_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/hashers"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store/types/chain"
	"github.com/thumbrise/golang-learning/internal/structures/hashtable/store/types/openaddr/linearprob"
)

type StoreFactory[T any] struct {
	Name    string
	Factory func(size int) store.Store[T]
}

func getStores[T any]() []StoreFactory[T] {
	return []StoreFactory[T]{
		{
			Name: "chain",
			Factory: func(size int) store.Store[T] {
				return chain.NewStore[T](size)
			},
		},
		{
			Name: "open addr linear prob",
			Factory: func(size int) store.Store[T] {
				return linearprob.NewStore[T](size)
			},
		},
	}
}

func TestHashTableConcurrentUniqueKeys(t *testing.T) {
	t.Parallel()

	const count = 1000

	for _, st := range getStores[string]() {
		t.Run("store="+st.Name, func(t *testing.T) {
			t.Parallel()

			ht := hashtable.NewHashTable[string](0, nil, st.Factory)

			var wg sync.WaitGroup
			for i := range count {
				wg.Add(1)

				go func(n int) {
					defer wg.Done()

					key := fmt.Sprintf("key-%d", n)
					ht.Set(key, fmt.Sprintf("value-%d", n))
				}(i)
			}

			wg.Wait()
			// Проверяем, что все ключи на месте
			for i := range count {
				key := fmt.Sprintf("key-%d", i)

				want := fmt.Sprintf("value-%d", i)
				if got := ht.Get(key); got != want {
					t.Errorf("for key %s: got %s, want %s", key, got, want)
				}
			}
		})
	}
}

func TestHashTableSet(t *testing.T) {
	t.Parallel()

	for _, st := range getStores[string]() {
		t.Run("store="+st.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("Set", func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, st.Factory)

				const (
					key   = "key"
					value = "value"
				)

				h.Set(key, value)

				if got := h.Get(key); got != value {
					t.Errorf("Get() = %#v, want %#v", got, value)
				}
			})

			t.Run("Overwrite value", func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, st.Factory)

				const (
					key   = "key"
					value = "value"
				)

				h.Set(key, value)

				if got := h.Get(key); got != value {
					t.Errorf("Get() = %#v, want %#v", got, value)
				}

				newValue := "new value"
				h.Set(key, newValue)

				if got := h.Get(key); got != newValue {
					t.Errorf("Get() = %#v, want %#v", got, newValue)
				}
			})
			t.Run("Overwrite key", func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, st.Factory)

				const (
					key   = "key"
					value = "value"
				)

				h.Set(key, value)

				if got := h.Get(key); got != value {
					t.Errorf("Get() = %#v, want %#v", got, value)
				}

				newKey := "new key"
				h.Set(newKey, value)

				if got := h.Get(key); got != value {
					t.Errorf("Get() = %#v, want %#v", got, value)
				}

				if got := h.Get(newKey); got != value {
					t.Errorf("Get() = %#v, want %#v", got, value)
				}
			})

			t.Run("Same keys overwrite instead of stacking", func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, st.Factory)

				const (
					key = "key"
				)

				values := [50]string{}

				for i := range 50 {
					values[i] = fmt.Sprintf("value%d", i)
					h.Set(key, values[i])
				}

				// TODO:
				//     hashtable_test.go:165: Get() = "value0", want "value49"
				//--- FAIL: TestHashTableSet/store=open_addr_linear_prob/Same_keys_overwrite_instead_of_stacking (0.00s)
				if got := h.Get(key); got != values[49] {
					t.Errorf("Get() = %#v, want %#v", got, values[49])
				}

				h.Delete(key)

				if got := h.Get(key); got != "" {
					t.Errorf("Get() = %#v, want %#v", got, "")
				}

				veryNewValue := "very new value"
				h.Set(key, veryNewValue)

				if got := h.Get(key); got != veryNewValue {
					t.Errorf("Get() = %#v, want %#v", got, veryNewValue)
				}
			})
		})
	}
}

func TestHashTableGet(t *testing.T) {
	t.Parallel()

	for _, st := range getStores[string]() {
		t.Run("store="+st.Name, func(t *testing.T) {
			t.Parallel()

			h := hashtable.NewHashTable[string](0, nil, st.Factory)

			const (
				key   = "key"
				value = "value"
			)

			h.Set(key, value)

			if got := h.Get(key); got != value {
				t.Errorf("Get() = %#v, want %#v", got, value)
			}

			t.Run("Get non-existent key", func(t *testing.T) {
				t.Parallel()

				h := hashtable.NewHashTable[string](0, nil, st.Factory)

				const (
					key  = "non-existent-key"
					want = ""
				)

				if got := h.Get(key); got != want {
					t.Errorf("Get() = %#v, want %#v", got, want)
				}
			})
		})
	}
}

func TestHashTableDelete(t *testing.T) {
	t.Parallel()

	for _, st := range getStores[string]() {
		t.Run("store="+st.Name, func(t *testing.T) {
			t.Parallel()

			h := hashtable.NewHashTable[string](0, nil, st.Factory)

			const (
				key   = "key"
				value = "value"
				want  = ""
			)

			h.Set(key, value)

			if got := h.Get(key); got != value {
				t.Errorf("Get() = %#v, want %#v", got, value)
			}

			// delete
			h.Delete(key)

			if got := h.Get(key); got != "" {
				t.Errorf("Get() after delete = %#v, want %#v", got, want)
			}

			presenceTests := []struct {
				name string
				size int
			}{
				{name: "Other items still present in table", size: 0},
				{name: "Other items still present in same bucket", size: 1},
			}
			for _, test := range presenceTests {
				t.Run(test.name, func(t *testing.T) {
					t.Parallel()

					h := hashtable.NewHashTable[string](test.size, nil, st.Factory)

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

					// TODO: === CONT  TestHashTableDelete/store=open_addr_linear_prob/Other_items_still_present_in_same_bucket
					//    hashtable_test.go:293: Get() = "", want "value1"
					//    --- FAIL: TestHashTableDelete/store=open_addr_linear_prob/Other_items_still_present_in_same_bucket (0.00s)
					//    Тут где-то есть гонка
					//    Ошибка раз через раз

					// Ожидаем, что key1 и key3 сохранились, а key2 удален
					if got := h.Get(key1); got != want1 {
						t.Errorf("Get() = %#v, want %#v", got, want1)
					}

					if got := h.Get(key2); got != want2 {
						t.Errorf("Get() after delete = %#v, want %#v", got, want2)
					}

					if got := h.Get(key3); got != want3 {
						t.Errorf("Get() = %#v, want %#v", got, want3)
					}
				})
			}
		})
	}
}

func BenchmarkHashTable_InsertAfterFill(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	fills := []int{10, 100, 1000, 10000}
	badHasher := hashers.NewFirstRuneReturnHasher()
	goodHasher := hashers.NewMapHashHasher()
	hshrs := []hashtable.Hasher{badHasher, goodHasher}
	strategies := []string{"chain", "open"}

	// Предгенерируем пул ключей (достаточно большой)
	const maxKeys = 20000

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
