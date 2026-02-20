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
		switch {
		case len(space) == 0 || space[len(space)-1] <= 0 || asteroids[r] >= 0:
			// Не с чем проверять || Не столкнутся никогда - просто добавляем правого
			space = append(space, asteroids[r])
			r++
		case space[len(space)-1] < -asteroids[r]:
			// Левый меньше по модулю - убиваем левого
			space = space[:len(space)-1]
		case space[len(space)-1] > -asteroids[r]:
			// Правый меньше по модулю - убиваем правого
			r++
		default:
			// Оба уничтожены =(
			space = space[:len(space)-1]
			r++
		}
	}

	return space
}
