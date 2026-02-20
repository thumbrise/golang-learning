package asteroid_collision

func AsteroidCollision(asteroids []int) []int {
	result := make([]int, 0, len(asteroids))
	result = append(result, asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		a := result[len(result)-1]
		b := asteroids[i]

		if a*b > 0 {
			// same direction
			// push new
		}

		if a > b {
			// old win
			// skip new
		}

		if a == -b {
			// both explodes
			// pop old and skip new
		}

		// new win
		// pop old and push new
	}

	return result
}
