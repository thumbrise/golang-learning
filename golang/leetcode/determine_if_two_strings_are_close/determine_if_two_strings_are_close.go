package determine_if_two_strings_are_close

import (
	"maps"
	"slices"
)

// CloseStrings
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Логика алгоритма:
//
// Чтобы понять что строки близки мы:
//  1. Проверяем длину строк - просто проверяем что длины строк равны
//  2. Проверяем, что строки состоят из одинаковых символов.
//     2.1. Делаем две мапы счетчика.
//     2.2. Проверяем, что в обеих мапах есть одни и те же ключи. Иначе сразу return false.
//     2.3. Для оптимизации сразу считаем частоты символов в строках. Делаем инкременты в мапах с помощью m[string(char)]++
//  3. Проверяем, что частоты символов в строках совпадают.
//     3.1. Сортируем значения мап счетчиков и сравниваем их. Получаем две сортированные последовательности с помощью slices.Sorted().
//     3.2. В цикле буквально смотрим совпадают ли значения последовательностей. seq1[i] == seq2[i]. Иначе false.
//     3.3 Таким образом мы понимаем, что количества Replacable символов в строках совпадают.
func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	if word1 == word2 {
		return true
	}

	entries1 := map[string]int{}
	for _, char := range word1 {
		entries1[string(char)]++
	}

	entries2 := map[string]int{}
	for _, char := range word2 {
		_, ok := entries1[string(char)]
		if !ok {
			return false
		}

		entries2[string(char)]++
	}

	seq1 := slices.Sorted(maps.Values(entries1))
	seq2 := slices.Sorted(maps.Values(entries2))

	for i := range seq1 {
		if seq1[i] != seq2[i] {
			return false
		}
	}

	return true
}
