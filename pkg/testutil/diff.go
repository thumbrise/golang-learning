package testutil

import (
	"encoding/json"
	"testing"

	"github.com/thumbrise/golang-learning/pkg/diff"
)

func ErrorDiff(t *testing.T, msg string, got, want interface{}) {
	aJSON, _ := json.MarshalIndent(got, "", "\t")
	bJSON, _ := json.MarshalIndent(want, "", "\t")
	t.Errorf("%s\ngot:%s\nwant:%s\ndiff:\n%s\n", msg, string(aJSON), string(bJSON), diff.Any(aJSON, bJSON))
}
