package determine_if_two_strings_are_close_test

import (
	"testing"

	"github.com/thumbrise/golang-learning/internal/leetcode/determine_if_two_strings_are_close"
)

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		{
			name:  "Example 1 - simple permutation",
			word1: "abc",
			word2: "bca",
			want:  true,
		},
		{
			name:  "Example 2 - different lengths",
			word1: "a",
			word2: "aa",
			want:  false,
		},
		{
			name:  "Example 3 - complex transformation",
			word1: "cabbba",
			word2: "abbccc",
			want:  true,
		},
		{
			name:  "Same strings",
			word1: "hello",
			word2: "hello",
			want:  true,
		},
		{
			name:  "Different character sets",
			word1: "abc",
			word2: "def",
			want:  false,
		},
		{
			name:  "Same frequency pattern but different chars",
			word1: "aabbcc",
			word2: "xxyyzz",
			want:  false,
		},
		{
			name:  "Different frequency patterns",
			word1: "aaabbb",
			word2: "ababab",
			want:  true,
		},
		{
			name:  "Empty strings",
			word1: "",
			word2: "",
			want:  true,
		},
		{
			name:  "Different unique chars count",
			word1: "abc",
			word2: "abcd",
			want:  false,
		},
		{
			name:  "Leetcode test - different frequencies",
			word1: "uau",
			word2: "ssx",
			want:  false,
		},
		{
			name:  "Leetcode test - same frequencies different sets",
			word1: "aaabbbbccddeeeee",
			word2: "aaaaabbbbcccdde",
			want:  false,
		},
		{
			name:  "All same characters",
			word1: "aaaa",
			word2: "aaaa",
			want:  true,
		},
		{
			name:  "One char transformation",
			word1: "ab",
			word2: "ba",
			want:  true,
		},
		{
			name:  "Failed case 1 - cabbba vs aabbss",
			word1: "cabbba",
			word2: "aabbss",
			want:  false,
		},
		{
			name:  "Failed case 2 - abbbzcf vs babzzcz",
			word1: "abbbzcf",
			word2: "babzzcz",
			want:  false,
		},
		{
			name:  "Additional test - different frequencies same chars",
			word1: "aabbbc",
			word2: "abbbcc",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determine_if_two_strings_are_close.CloseStrings(tt.word1, tt.word2); got != tt.want {
				t.Errorf("CloseStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
