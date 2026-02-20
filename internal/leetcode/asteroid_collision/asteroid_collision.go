package asteroid_collision

// AsteroidCollision (Classic)
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Логика алгоритма:
//
// Нужно вычислить конечное состояние астероидов после всех возможных столкновений.
// 1. Используем стек для отслеживания астероидов, которые не взорвались.
// 2. Проходим по массиву астероидов.
// 3. Каждый раз берем пару left из стека и right из аргумента asteroids.
// 4. Проверяем возможные сценарии того, кто из них выживет.
// 5. Если левый взорвется, удаляем его из стека. Если правый при этом выжил, то добавляем его в стек.
// 6. Если левый выжил, то в случае смерти правого делаем continue, если правый тоже выжил, то добавляем его в стек рядом с левым
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

// AsteroidCollisionImproved (Improved)
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Логика алгоритма:
//
// Нужно вычислить конечное состояние астероидов после всех возможных столкновений.
// 1. Используем стек для отслеживания астероидов, которые не взорвались.
// 2. Проходим по массиву астероидов.
// 3. Каждый раз берем пару left из стека и right из аргумента asteroids.
// 4. Проверяем возможные сценарии того, кто из них выживет.
// 5. Если левый взорвется, удаляем его из стека. Если правый при этом выжил, то добавляем его в стек.
// 6. Если левый выжил, то в случае смерти правого делаем continue, если правый тоже выжил, то добавляем его в стек рядом с левым
func AsteroidCollisionImproved(asteroids []int) []int {
	space := make([]int, 0, len(asteroids))
	// space = append(space, asteroids[0])
	for l, r := 0, 0; r < len(asteroids); {
		rightAsteroid := asteroids[r]

		if len(space) == 0 {
			space = append(space, rightAsteroid)
			r++
			l = 0

			continue
		}

		leftAsteroid := space[l]
		leftSurvive, rightSurvive := predictSurvive(leftAsteroid, rightAsteroid)

		if !leftSurvive {
			space = space[:len(space)-1]
			l--
		}

		if leftSurvive && rightSurvive {
			space = append(space, rightAsteroid)
			r++
			l++
		} else if !rightSurvive {
			r++
		}
	}

	return space
}

func predictSurvive(l, r int) (bool, bool) {
	// never meets or same direction
	if l < 0 && r > 0 || l*r > 0 {
		return true, true
	}

	lmod := abs(l)
	rmod := abs(r)

	// left explode
	if lmod < rmod {
		return false, true
	}

	// right explode
	if lmod > rmod {
		return true, false
	}

	// both explode
	return false, false
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}

	return v
}
