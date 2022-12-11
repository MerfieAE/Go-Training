package algorithms

func RecSum(arr []int, n int) int {
	if n <= 0 {
		return 0
	}
	return RecSum(arr, n-1) + arr[n-1]
}
