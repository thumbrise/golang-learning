package asteroid_collision

func AsteroidCollisionClassic(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, ast := range asteroids {
		// пока есть столкновение (последний положительный, текущий отрицательный)
		for len(stack) > 0 && stack[len(stack)-1] > 0 && ast < 0 {
			last := stack[len(stack)-1]
			if last < -ast { // левый взрывается
				stack = stack[:len(stack)-1]

				continue // продолжаем проверять с новым последним
			} else if last == -ast { // оба взрываются
				stack = stack[:len(stack)-1]
				ast = 0

				break
			} else { // правый взрывается
				ast = 0

				break
			}
		}

		if ast != 0 {
			stack = append(stack, ast)
		}
	}

	return stack
}

func AsteroidCollisionImproved(asteroids []int) []int {
	space := make([]int, 0, len(asteroids))

	r := 0
	for r < len(asteroids) {
		rightAsteroid := asteroids[r]

		if len(space) == 0 {
			space = append(space, rightAsteroid)
			r++

			continue
		}

		leftAsteroid := space[len(space)-1]

		switch {
		case leftAsteroid <= 0 || rightAsteroid >= 0:
			space = append(space, rightAsteroid)
			r++
		case leftAsteroid < -rightAsteroid:
			space = space[:len(space)-1]
		case leftAsteroid > -rightAsteroid:
			r++
		default:
			space = space[:len(space)-1]
			r++
		}
	}

	return space
}
