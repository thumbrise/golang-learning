package removing_stars_from_a_string_test

import (
	"testing"

	"github.com/thumbrise/golang-learning/internal/leetcode/removing_stars_from_a_string"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removing_stars_from_a_string.RemoveStars(tt.args.s); got != tt.want {
				t.Errorf("RemoveStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
