package removing_stars_from_a_string

import "strings"

func RemoveStars(s string) string {
	const star = '*'

	rns := make([]rune, 0, len(s))

	for _, ch := range s {
		if ch == star {
			rns = append(rns[:len(rns)-1])
			continue
		}

		rns = append(rns, ch)
	}

	b := strings.Builder{}
	b.Grow(len(s) + 1)
	for _, r := range rns {
		b.WriteRune(r)
	}

	return b.String()
}
