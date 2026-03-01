package removing_stars_from_a_string

// RemoveStarts
//
// Time Complexity: O(n)
// Space Complexity: O(n)
//
// Логика алгоритма:
// 1. Используем слайс как стек, чтобы накапливать символы.
// 2. Проходим по строке, если символ не '*', добавляем его в стек.
// 3. Если символ '*', удаляем последний элемент из стека.
//
// BenchmarkRemoveStars/Stack-14           128717398                8.708 ns/op           0 B/op          0 allocs/op
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

// RemoveStarsTwoPointers
//
// Time Complexity: O(n)
// Space Complexity: O(n)
//
// Логика алгоритма:
// 1. Используем два указателя: один для чтения, другой для записи.
// 2. Проходим по строке, если символ не '*', записываем его в позицию j, увеличиваем j.
// 3. Если символ '*', уменьшаем j.
//
// Как это работает:
// - j указывает на позицию, куда нужно записать следующий символ.
// - Когда встречаем '*', уменьшаем j, effectively "удаляя" последний добавленный символ.
//
// В целом, это подход из C++ где строки являются изменяемыми массивами байтов.
// Там можно делать замену символов in place.
//
// BenchmarkRemoveStars/TwoPointers-14     150823132                8.017 ns/op           0 B/op          0 allocs/op
func RemoveStarsTwoPointers(s string) string {
	b := []byte(s)
	j := 0

	for i := 0; i < len(b); i++ {
		if b[i] == '*' {
			j--
		} else {
			b[j] = b[i]
			j++
		}
	}

	return string(b[:j])
}
