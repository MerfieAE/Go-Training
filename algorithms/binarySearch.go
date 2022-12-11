package algorithms

import "fmt"

func BinarySearch(strs []string, str string) int {
	low := 0
	high := len(strs) - 1

	for low <= high {
		mid := (low + high) / 2
		if strs[mid] == str {
			fmt.Println(mid) // check in main
			return mid
		}
		if strs[mid] > str {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println(-1)
	return -1
}
