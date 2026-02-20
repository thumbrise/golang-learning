package asteroid_collision

func AsteroidCollision(asteroids []int) []int {
	result := make([]int, 0, len(asteroids))
	result = append(result, asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		astOld := result[len(result)-1]
		astNew := asteroids[i]

		// same direction
		// push new
		if astOld*astNew > 0 {
			result = append(result, astNew)
			continue
		}

		// both explodes
		// pop old and skip new
		if astOld == -astNew {
			result = result[:len(result)-1]
			continue
		}

		if astOld < astNew {
			// new win
			// pop old and push new
			result = append(result[:len(result)-1], astNew)
			continue
		}

		// old win
		// skip new
	}

	return result
}
