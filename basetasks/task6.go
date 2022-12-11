package basetasks

import (
	"fmt"
	"time"
)

/*
 Задача 6. Реализовать все возможные способы остановки выполнения горутины
*/

func TaskSix() {
	ch1 := make(chan string)
	run := true

	go func(run *bool) {
		for {
			select {
			// 1 способ: дождаться закрытия канала
			case <-ch1:
				fmt.Println("Goroutine 1 done (<-ch1)!")
				return
			default:
			}
		}
	}(&run)

	time.Sleep(150 * time.Millisecond)
	run = false
	time.Sleep(350 * time.Millisecond)
	// остановка горутины по закрытию канала. Через этот же канал можно получать данные из горутины, но
	// в таком случае нужно синхронизировать момент после закрытия канала и работы горутины с каналом:
	// чтобы горутина не писала в закрытый канал (panic)
	close(ch1)

	ch2 := make(chan int)
	go func() {
		for {
			// 2 способ: используем второе возвращаемое каналом значение
			// num - значение из канала, а more - bool переменная, равная false, если канал закрыт
			num, more := <-ch2
			if !more {
				fmt.Println("Goroutine 2 done!")
				return
			}
			// какая-то полезная работа
			fmt.Printf("Goroutine 2 working says %d\n", num)
		}
	}()

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)

	time.Sleep(350 * time.Millisecond)
}
