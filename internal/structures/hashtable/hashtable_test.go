package hashtable_test

import (
	"strconv"
	"testing"

	"github.com/thumbrise/golang-learning/internal/structures/hashtable"
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

			h := hashtable.NewHashTable[string](0, nil)

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

				h := hashtable.NewHashTable[string](0, nil)
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

	h := hashtable.NewHashTable[string](0, nil)

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
