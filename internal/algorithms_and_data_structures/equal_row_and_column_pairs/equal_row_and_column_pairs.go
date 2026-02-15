package equal_row_and_column_pairs

import (
	"encoding/json"
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

	matches := make(map[string]int)

	for k, v := range rows {
		if cols[k] > 0 {
			matches[k] = v * cols[k]
		}
	}

	result := 0
	for _, v := range matches {
		result += v
	}

	prnt("rows", rows)
	prnt("cols", cols)
	prnt("matches", matches)

	return result
}

func prnt(label string, entries map[string]int) {
	jsn, err := json.MarshalIndent(entries, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s:\n%s\n", label, jsn)
}
