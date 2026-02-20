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

		// new win
		// pop old and push new
		if abs(astOld) < abs(astNew) {
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

func AsteroidCollisionImproved(asteroids []int) []int {
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
		oldSurvive, newSurvive := predictSurvive(astOld, astNew)

		switch {
		case oldSurvive && newSurvive:
			result = append(result, astNew)
			i++
			j++
		case !oldSurvive && !newSurvive:
			result = result[:len(result)-1]
			j--
			i++
		case !oldSurvive:
			result = result[:len(result)-1]
			j--
		default:
			i++
		}
	}

	return result
}
func predictSurvive(l, r int) (bool, bool) {
	sameDirection := l < 0 && r > 0
	neverMeets := l*r > 0
	lmod := abs(l)
	rmod := abs(r)
	explodesBoth := lmod == rmod
	leftLose := lmod < rmod

	switch {
	case sameDirection || neverMeets:
		return true, true
	case explodesBoth:
		return false, false
	case leftLose:
		return false, true
	}
	return true, false
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}

	return v
}
