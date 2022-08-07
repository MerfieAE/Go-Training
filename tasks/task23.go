package tasks

/*
 Задача 23. Удалить i-ый элемент из слайса.
*/

import "fmt"

func TaskTwentyThee(arr []int, i int) []int {
	if len(arr) != 0 && i < len(arr)-1 { // защищаемся от паники
		arr = append(arr[:i], arr[i+1:]...)
	}
	fmt.Println(arr)
	return arr
}
