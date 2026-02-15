package tower_of_hanoi_test

import (
	"reflect"
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/thumbrise/golang-learning/internal/algorithms_and_data_structures/tower_of_hanoi"
	"github.com/thumbrise/golang-learning/pkg/testutil"
)

func TestTowerOfHanoi(t *testing.T) {
	t.Parallel()

	type towersSet struct {
		origin      *tower_of_hanoi.Tower
		destination *tower_of_hanoi.Tower
		auxiliary   *tower_of_hanoi.Tower
	}
	tests := []struct {
		name string
		args towersSet
		want towersSet
	}{
		{
			name: "Classic",
			args: towersSet{
				origin: &tower_of_hanoi.Tower{
					Name: "A",
					Disks: (func() *arraystack.Stack {
						d := arraystack.New()
						_ = d.FromJSON([]byte("[1,2,3]"))
						return d
					})(),
				},
				destination: &tower_of_hanoi.Tower{
					Name:  "B",
					Disks: arraystack.New(),
				},
				auxiliary: &tower_of_hanoi.Tower{
					Name:  "C",
					Disks: arraystack.New(),
				},
			},
			want: towersSet{
				origin: &tower_of_hanoi.Tower{
					Name:  "A",
					Disks: arraystack.New(),
				},
				destination: &tower_of_hanoi.Tower{
					Name:  "B",
					Disks: arraystack.New(),
				},
				auxiliary: &tower_of_hanoi.Tower{
					Name: "C",
					Disks: (func() *arraystack.Stack {
						d := arraystack.New()
						_ = d.FromJSON([]byte("[1,2,3]"))
						return d
					})(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tower_of_hanoi.TowerOfHanoi(tt.args.origin, tt.args.destination, tt.args.auxiliary)
			got := []*tower_of_hanoi.Tower{
				tt.args.origin,
				tt.args.destination,
				tt.args.auxiliary,
			}
			want := []*tower_of_hanoi.Tower{
				tt.want.origin,
				tt.want.destination,
				tt.want.auxiliary,
			}
			for i := range got {
				if !reflect.DeepEqual(got[i], want[i]) {
					testutil.ErrorDiff(t, "TowerOfHanoi()", got, want)
					break
				}
			}
		})
	}
}
