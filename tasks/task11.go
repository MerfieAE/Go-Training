package tasks

/*
 Задача 11. Реализовать пересечение двух неупорядоченных множеств
*/

import "fmt"

func TaskEleven() {
	set1 := []int{1, 7, 3, 5, 3}
	set2 := []int{3, 6, 1, 9, 3, 7, 2, 4}
	result := []int{}

	for _, val1 := range set1 {
		for _, val2 := range set2 {
			if val1 == val2 {
				result = append(result, val1)
			}
		}
	}
	fmt.Println(result)
}
