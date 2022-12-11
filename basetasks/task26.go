package basetasks

import "strings"

/*
 Задача 26. Разработать программу, которая проверяет, что все символы в строке
 уникальные (true — если уникальные, false etc). Функция проверки должна быть
 регистронезависимой.
*/

func TaskTwentySix(s string) bool {
	// map обеспечит уникальность
	m := make(map[rune]bool)

	s = strings.ToLower(s)
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
