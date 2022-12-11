package basetasks

import (
	"fmt"
	"time"
)

/*
	Задача 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
*/

func worker(workerName int, in <-chan int) {
	for {
		num := <-in
		fmt.Printf("Goroutine #%d: value: %ds\n", workerName, num)
	}
}

func TaskFour() {
	var N int
	fmt.Println("Количество горутин:")
	fmt.Scanf("%d\n", &N)

	workerInput := make(chan int)

	// создаем горутины
	for i := 0; i < N; i++ {
		go worker(i, workerInput)
	}

	// постоянная запись данных в канал из главного потока
	for {
		workerInput <- time.Now().Second()
		time.Sleep(time.Second)
	}
}
