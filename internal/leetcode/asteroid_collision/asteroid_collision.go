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
	space := make([]int, 0, len(asteroids))
	for r, l := 0, 0; r < len(asteroids); {
		rightAsteroid := asteroids[r]

		if len(space) == 0 {
			space = append(space, rightAsteroid)
			r++
			l = 0

			continue
		}

		leftAsteroid := space[l]
		leftSurvive, rightSurvive := predictSurvive(leftAsteroid, rightAsteroid)

		// refactor conditions
		if !leftSurvive {
			space = space[:len(space)-1]
			l--
		} else {
			//r++
		}

		switch {
		case leftSurvive && rightSurvive:
			space = append(space, rightAsteroid)
			r++
			l++
		case !leftSurvive && !rightSurvive:
			r++
		case !rightSurvive:
			r++
		}

	}

	return space
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
