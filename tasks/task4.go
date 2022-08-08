package tasks

import (
	"fmt"
	"math/rand"
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

func readFromChanel(chanel chan int) {

	for data := range chanel {
		fmt.Printf("now data is %d\n", data)
	}
}

func writeToChanel(chanel chan int) {
	for {
		data := rand.Intn(100)
		chanel <- data
	}
}

func TaskFour() {
	N := 10

	chanel := make(chan int)
	defer close(chanel)
	go writeToChanel(chanel)
	go readFromChanel(chanel)

	time.Sleep(time.Duration(N) * time.Second)
}
