package algorithms_and_data_structures

import (
	"fmt"
	"strings"
)

func TowerOfHanoi(n int, from, to, aux string, result *strings.Builder) string {
	if n == 1 {
		step(result, from, to, n)
		return result.String()
	}

	TowerOfHanoi(n-1, from, aux, to, result)
	step(result, from, to, n)
	TowerOfHanoi(n-1, aux, to, from, result)

	return result.String()
}

func step(b *strings.Builder, from, to string, n int) {
	b.WriteString(fmt.Sprintf("%s, %s, %d\n", from, to, n))
}

// TowerOfHanoi(3, A, B, C)
//
