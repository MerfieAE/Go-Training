package tasks

/*
5. Разработать программу, которая будет последовательно отправлять значения в канал,
 а с другой стороны канала — читать. По истечению N секунд программа должна завершаться
*/

import (
	"fmt"
	"log"
	"time"
)

var N time.Duration = 7 // Допустим N = 7 секунд

func connectToResource(c chan string) {
	//Тут time.Sleep() для наглядности
	time.Sleep(5 * time.Second) // чтобы main получила значение из этой горутины, тут время должно быть меньше, чем в main
	c <- "Info: connection is created"
}

func TaskFive() {
	fmt.Println("Main start ...")

	ch := make(chan string)

	go connectToResource(ch)

	select {
	case res := <-ch:
		log.Println(res)
	case <-time.After(N * time.Second): // А тут время должно быть больше, чем в connectToResource, иначе ошибка.
		log.Println("Err: connection aborted: timeout")
	}

	fmt.Println("Main stop ...")
}
