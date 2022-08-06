package tasks

import (
	"fmt"
	"sync"
)

func squareNum(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for idx := 1; idx <= len(arr)-1; idx++ {
		fmt.Printf("Square number of %d is %d\n", arr[idx], arr[idx]*arr[idx])
	}
}

func TaskTwo() {
	var wg sync.WaitGroup

	wg.Add(1)
	go squareNum([]int{2, 4, 6, 8, 10}, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
