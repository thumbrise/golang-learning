package hashtable

import (
	"testing"
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
			h := NewHashTable[string](0, nil)
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
				h := NewHashTable[string](0, nil)
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
	// ht := NewHashTable()
}

func TestHashTableCollision(t *testing.T) {
	t.Parallel()
	// ht := NewHashTable()
}
