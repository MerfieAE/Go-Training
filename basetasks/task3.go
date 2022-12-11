package basetasks

/*
3. Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

func squareSum(arr int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := 1; i <= arr; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	fmt.Println(sum)
}

func TaskThree() {
	var wg sync.WaitGroup

	wg.Add(1)
	go squareSum(10, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
