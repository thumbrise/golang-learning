package removing_stars_from_a_string_test

import (
	"testing"

	"github.com/thumbrise/demo/internal/leetcode/removing_stars_from_a_string"
)

func TestRemoveStars(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "leet**cod*e",
			},
			want: "lecoe",
		},
		{
			name: "Example 2",
			args: args{
				s: "erase*****",
			},
			want: "",
		},
		{
			name: "No stars",
			args: args{s: "abcdef"},
			want: "abcdef",
		},
		{
			name: "Single star at end",
			args: args{s: "abcd*"},
			want: "abc",
		},
		{
			name: "Single star in middle",
			args: args{s: "ab*cd"},
			want: "acd",
		},
		{
			name: "Two consecutive stars at end",
			args: args{s: "abcd**"},
			want: "ab",
		},
		{
			name: "Stars interleaved",
			args: args{s: "a*b*c*d"},
			want: "d",
		},
		{
			name: "All characters removed",
			args: args{s: "a*b*c*"},
			want: "",
		},
		{
			name: "Example 1",
			args: args{s: "leet**cod*e"},
			want: "lecoe",
		},
		{
			name: "Example 2",
			args: args{s: "erase*****"},
			want: "",
		},
		{
			name: "Single character",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Two characters no star",
			args: args{s: "ab"},
			want: "ab",
		},
		{
			name: "Complex pattern",
			args: args{s: "abc*d*efg**"},
			want: "abe", // пошагово: abc*d*efg** -> ab*d*efg** -> ab*efg** -> aefg** -> aef* -> ae? пересчёт в стеке даёт "abe"
		},
		{
			name: "Three stars at the end",
			args: args{s: "hello***"},
			want: "he",
		},
		{
			name: "Stars in the middle and end",
			args: args{s: "ab**c*d"},
			want: "d", // ab**c*d: после двух звёзд удаляются a и b, остаётся "c*d", затем удаляется c -> "d"
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removing_stars_from_a_string.RemoveStars(tt.args.s); got != tt.want {
				t.Errorf("RemoveStars() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removing_stars_from_a_string.RemoveStarsTwoPointers(tt.args.s); got != tt.want {
				t.Errorf("RemoveStarsTwoPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveStars(b *testing.B) {
	const testString = "leet**cod*e"

	b.Run("Stack", func(b *testing.B) {
		for range b.N {
			removing_stars_from_a_string.RemoveStars(testString)
		}
	})
	b.Run("TwoPointers", func(b *testing.B) {
		for range b.N {
			removing_stars_from_a_string.RemoveStarsTwoPointers(testString)
		}
	})
}
