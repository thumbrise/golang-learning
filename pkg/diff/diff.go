package diff

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Any(a, b []byte) string {
	d := diffmatchpatch.New()
	df := d.DiffMain(string(a), string(b), true)

	return d.DiffPrettyText(df)
}
