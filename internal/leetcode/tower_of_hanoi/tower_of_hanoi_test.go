package tower_of_hanoi_test

import (
	"reflect"
	"testing"

	"github.com/thumbrise/demo/internal/leetcode/tower_of_hanoi"
	"github.com/thumbrise/demo/pkg/testutil"
)

type towerSet struct {
	Origin      *tower_of_hanoi.Tower
	Auxiliary   *tower_of_hanoi.Tower
	Destination *tower_of_hanoi.Tower
}

func (ts towerSet) toSlice() []*tower_of_hanoi.Tower {
	return []*tower_of_hanoi.Tower{ts.Origin, ts.Auxiliary, ts.Destination}
}

func TestTowerOfHanoi(t *testing.T) {
	t.Parallel()

	const n = 3

	given := towerSet{
		Origin:      tower_of_hanoi.NewTower("A", n),
		Auxiliary:   tower_of_hanoi.NewTower("B", 0),
		Destination: tower_of_hanoi.NewTower("C", 0),
	}

	want := towerSet{
		Origin:      tower_of_hanoi.NewTower("A", 0),
		Auxiliary:   tower_of_hanoi.NewTower("B", 0),
		Destination: tower_of_hanoi.NewTower("C", n),
	}

	tower_of_hanoi.TowerOfHanoi(
		n,
		given.Origin,      // A
		given.Destination, // C
		given.Auxiliary,   // B
	)

	got := given

	t.Logf("Origin disks: %v", given.Origin.Disks.Values())
	t.Logf("Auxiliary disks: %v", given.Auxiliary.Disks.Values())
	t.Logf("Destination disks: %v", given.Destination.Disks.Values())

	if !reflect.DeepEqual(got.toSlice(), want.toSlice()) {
		testutil.ErrorDiff(t, "TowerOfHanoi()", got.toSlice(), want.toSlice())
	}
}
