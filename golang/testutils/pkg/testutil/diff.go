package testutil

import (
	"encoding/json"
	"testing"

	"github.com/andreyvit/diff"
)

func ErrorDiff(t *testing.T, msg string, got, want interface{}) {
	t.Helper()

	gotJSON, err := json.MarshalIndent(got, "", "\t")
	if err != nil {
		t.Fatalf("failed to marshal got: %v", err)
	}

	wantJSON, err := json.MarshalIndent(want, "", "\t")
	if err != nil {
		t.Fatalf("failed to marshal want: %v", err)
	}

	t.Errorf("%s\ngot:%s\nwant:%s\ndiff:\n%s\n", msg, string(gotJSON), string(wantJSON), dif(wantJSON, gotJSON))
}

func dif(a, b []byte) string {
	return diff.LineDiff(string(a), string(b))
}
