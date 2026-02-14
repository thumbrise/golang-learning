package algorithms_and_data_structures_test

import (
	"strings"
	"testing"

	"github.com/thumbrise/golang-learning/internal/algorithms_and_data_structures"
)

func TestTowerOfHanoi(t *testing.T) {
	var result = &strings.Builder{}
	t.Run("ok", func(t *testing.T) {
		got := algorithms_and_data_structures.TowerOfHanoi(3, "A", "B", "C", result)
		t.Logf("\n%s", got)
	})
}
