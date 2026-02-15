package determine_if_two_strings_are_close

import (
	"maps"
	"slices"
)

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	if word1 == word2 {
		return true
	}

	entries1 := map[string]int{}
	for _, char := range word1 {
		entries1[string(char)]++
	}

	entries2 := map[string]int{}
	for _, char := range word2 {
		_, ok := entries1[string(char)]
		if !ok {
			return false
		}

		entries2[string(char)]++
	}

	seq1 := slices.Sorted(maps.Values(entries1))
	seq2 := slices.Sorted(maps.Values(entries2))

	for i := range seq1 {
		if seq1[i] != seq2[i] {
			return false
		}
	}

	return true
}
