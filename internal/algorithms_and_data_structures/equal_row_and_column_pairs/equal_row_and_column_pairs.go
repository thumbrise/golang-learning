package equal_row_and_column_pairs

import (
	"fmt"
)

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

func EqualPairs2(grid [][]int) int {
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
