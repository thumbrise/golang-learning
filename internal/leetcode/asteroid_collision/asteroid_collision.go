package asteroid_collision

func AsteroidCollision(asteroids []int) []int {
	result := make([]int, 0, len(asteroids))
	for i, j := 0, 0; i < len(asteroids); {
		astNew := asteroids[i]

		if len(result) == 0 {
			result = append(result, astNew)
			i++
			j = 0

			continue
		}

		astOld := result[j]

		// never meets each other, old fly left and new fly right->
		// or
		// same direction
		// push new
		if (astOld < 0 && astNew > 0) ||
			(astOld*astNew > 0) {
			result = append(result, astNew)
			i++
			j++

			continue
		}

		// both explodes
		// pop old and skip new
		if astOld == -astNew {
			result = result[:len(result)-1]
			i++
			j--

			continue
		}

		if abs(astOld) < abs(astNew) {
			// new win
			// pop old and push new
			result = result[:len(result)-1]
			j--

			continue
		}

		// old win
		// skip new
		i++
	}

	return result
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}

	return v
}
