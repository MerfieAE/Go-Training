package tasks

import (
	"fmt"
	"sync"
)

/*
 Задача 18.	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
 По завершению программа должна выводить итоговое значение счетчика.
*/

type incr struct {
	i int
}

func (i *incr) increment() {
	i.i++

}

func TaskEighteen() {
	inc := incr{0}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			inc.increment()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(inc)
}
