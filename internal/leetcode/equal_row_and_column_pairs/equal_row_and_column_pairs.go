package equal_row_and_column_pairs

import (
	"fmt"
)

// EqualPairs
//
// BenchmarkEqualPairs/String_Version-14         	     136	   8741902 ns/op
func EqualPairs(grid [][]int) int {
	rows := map[string]int{}
	cols := map[string]int{}

	for y := range grid {
		row := ""
		col := ""

		for x := range grid[y] {
			row += fmt.Sprintf("%d,", grid[y][x])
			col += fmt.Sprintf("%d,", grid[x][y])
		}

		rows[row]++
		cols[col]++
	}

	result := 0

	for k, v := range rows {
		if cols[k] > 0 {
			result += v * cols[k]
		}
	}

	return result
}

const maxN = 200

// EqualPairs2
//
// BenchmarkEqualPairs/Slice_Version-14          	   13845	     86531 ns/op
func EqualPairs2(grid [][]int) int {
	n := len(grid)
	rows := map[[maxN]int]int{}

	rowkey := [maxN]int{}
	for y := range grid {
		copy(rowkey[:], grid[y])
		rows[rowkey]++
	}

	result := 0

	for y := range n {
		colkey := [maxN]int{}
		for x := range n {
			colkey[x] = grid[x][y]
		}

		result += rows[colkey]
	}

	return result
}
