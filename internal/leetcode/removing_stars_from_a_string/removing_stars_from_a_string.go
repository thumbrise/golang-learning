package removing_stars_from_a_string

func RemoveStars(s string) string {
	r := make([]byte, 0, len(s))

	for _, b := range []byte(s) {
		if b == '*' {
			r = r[:len(r)-1]

			continue
		}

		r = append(r, b)
	}

	return string(r)
}
