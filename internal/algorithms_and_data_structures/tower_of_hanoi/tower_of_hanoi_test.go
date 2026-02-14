package tower_of_hanoi_test

import (
	"strings"
	"testing"

	"github.com/thumbrise/golang-learning/internal/algorithms_and_data_structures/tower_of_hanoi"
)

func TestTowerOfHanoi(t *testing.T) {
	var result = &strings.Builder{}
	t.Run("ok", func(t *testing.T) {
		got := tower_of_hanoi.TowerOfHanoi(3, "A", "B", "C", result)
		t.Logf("\n%s", got)
	})
}
