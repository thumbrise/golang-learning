package asteroid_collision_test

import (
	"reflect"
	"testing"

	"github.com/thumbrise/demo/internal/leetcode/asteroid_collision"
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
		{
			name: "All positive",
			args: args{asteroids: []int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "All negative",
			args: args{asteroids: []int{-1, -2, -3, -4}},
			want: []int{-1, -2, -3, -4},
		},
		{
			name: "Positive then negative, larger positive",
			args: args{asteroids: []int{10, -5}},
			want: []int{10},
		},
		{
			name: "Positive then negative, larger negative",
			args: args{asteroids: []int{5, -10}},
			want: []int{-10},
		},
		{
			name: "Positive then negative, equal",
			args: args{asteroids: []int{5, -5}},
			want: []int{},
		},
		{
			name: "Negative then positive, no collision",
			args: args{asteroids: []int{-5, 10}},
			want: []int{-5, 10},
		},
		{
			name: "Multiple collisions in sequence",
			args: args{asteroids: []int{8, 3, -6}},
			want: []int{8}, // 8,3 -> -6 уничтожает 3, затем 8 и -6: 8 > 6, -6 уничтожается
		},
		{
			name: "Negative destroys multiple positives",
			args: args{asteroids: []int{5, 4, 3, -10}},
			want: []int{-10},
		},
		{
			name: "Positive destroys multiple negatives",
			args: args{asteroids: []int{10, -5, -5}},
			want: []int{10},
		},
		{
			name: "Mixed with survivors left and right",
			args: args{asteroids: []int{1, -2, 3, -4}},
			want: []int{-2, -4}, // 1 и -2 -> -2; затем 3 (справа от -2) не сталкиваются; -4 уничтожает 3, затем оба отрицательных остаются
		},
		{
			name: "Alternating with survivors",
			args: args{asteroids: []int{1, -2, 3, -4, 5, -6}},
			want: []int{-2, -4, -6},
		},
		{
			name: "Chain reaction with equal sizes",
			args: args{asteroids: []int{8, -8, 8, -8}},
			want: []int{},
		},
		{
			name: "Large values",
			args: args{asteroids: []int{1000, -500, 200, -1000}},
			want: []int{}, // 1000 и -500: 1000 > 500, -500 уничтожается; затем 200; -1000 уничтожает 200 и сталкивается с 1000, оба равны — уничтожаются
		},
		{
			name: "Negative destroys positive then no further",
			args: args{asteroids: []int{500, -1000}},
			want: []int{-1000},
		},
		{
			name: "Positives at the end",
			args: args{asteroids: []int{-5, 10, 20}},
			want: []int{-5, 10, 20},
		},
		{
			name: "Minimum length both positive",
			args: args{asteroids: []int{1, 2}},
			want: []int{1, 2},
		},
		{
			name: "Minimum length positive then negative",
			args: args{asteroids: []int{1, -2}},
			want: []int{-2},
		},
		{
			name: "Minimum length negative then positive",
			args: args{asteroids: []int{-1, 2}},
			want: []int{-1, 2},
		},
		{
			name: "Multiple collisions with same size",
			args: args{asteroids: []int{10, -10, 10, -10}},
			want: []int{},
		},
		{
			name: "All positives after collisions",
			args: args{asteroids: []int{20, 15, -10, -5}},
			want: []int{20, 15}, // -10 и -5 уничтожаются по очереди, 20 и 15 остаются
		},
		{
			name: "Negative destroys positives then meets negative",
			args: args{asteroids: []int{5, 4, 3, -6, -7}},
			want: []int{-6, -7}, // -6 уничтожает 5,4,3, затем добавляется -7, оба отрицательные
		},
		{
			name: "Negative and positive alternating but no collision",
			args: args{asteroids: []int{-2, -1, 1, 2}},
			want: []int{-2, -1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("AsteroidCollisionClassic", func(t *testing.T) {
				if got := asteroid_collision.AsteroidCollisionClassic(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AsteroidCollisionClassic() = %v, want %v", got, tt.want)
				}
			})
			t.Run("AsteroidCollisionImproved", func(t *testing.T) {
				if got := asteroid_collision.AsteroidCollisionImproved(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AsteroidCollisionImproved() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}

func BenchmarkAsteroidCollision(b *testing.B) {
	asteroids := []int{5, 4, 3, -6, -7}

	b.Run("Classic", func(b *testing.B) {
		for range b.N {
			asteroid_collision.AsteroidCollisionClassic(asteroids)
		}
	})
	b.Run("Improved", func(b *testing.B) {
		for range b.N {
			asteroid_collision.AsteroidCollisionImproved(asteroids)
		}
	})
}
