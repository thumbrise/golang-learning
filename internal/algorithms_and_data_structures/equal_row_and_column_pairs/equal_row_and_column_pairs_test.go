package equal_row_and_column_pairs_test

import (
	"testing"

	"github.com/thumbrise/golang-learning/internal/algorithms_and_data_structures/equal_row_and_column_pairs"
)

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{3, 2, 1},
				{1, 7, 6},
				{2, 7, 7},
			},
			want: 1,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
			want: 3,
		},
		{
			name: "Single element grid",
			grid: [][]int{{42}},
			want: 1,
		},
		{
			name: "No matches",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 0,
		},
		{
			name: "All rows equal to all columns",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 4,
		},
		{
			name: "Diagonal matches only",
			grid: [][]int{
				{1, 2, 3},
				{4, 1, 6},
				{7, 8, 1},
			},
			want: 0,
		},
		{
			name: "Multiple matches with duplicates",
			grid: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 0,
		},
		{
			name: "Complex case with multiple equal rows and columns",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "Maximum size small test",
			grid: [][]int{
				{1, 2},
				{2, 1},
			},
			want: 2,
		},
		{
			name: "Ambiguous concatenation - numbers without separator",
			grid: [][]int{
				{1, 23},
				{12, 3},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal_row_and_column_pairs.EqualPairs(tt.grid); got != tt.want {
				t.Errorf("EqualPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
