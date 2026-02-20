package asteroid_collision_test

import (
	"reflect"
	"testing"

	"github.com/thumbrise/golang-learning/internal/leetcode/asteroid_collision"
)

func TestAsteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				asteroids: []int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "test 2",
			args: args{
				asteroids: []int{8, -8},
			},
			want: []int{},
		},
		{
			name: "test 3",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			want: []int{10},
		},
		{
			name: "test 4",
			args: args{
				asteroids: []int{3, 5, -6, 2, -1, 4},
			},
			want: []int{-6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroid_collision.AsteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
