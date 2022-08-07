package tasks

import (
	"fmt"
	"strings"
)

/*
 Задача 20. Разработать программу, которая переворачивает слова в строке.
 Пример: «snow dog sun — sun dog snow».
*/

func TaskTwenty(s string) string {
	arr := strings.Fields(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(strings.Join(arr, " "))
	return strings.Join(arr, " ")
}
