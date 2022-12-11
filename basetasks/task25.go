package basetasks

import (
	"fmt"
	"time"
)

/*
 Задача 25.	Реализовать собственную функцию sleep.
*/

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}

func TaskTwentyFive() {
	fmt.Println("Start...")
	Sleep(3)
	fmt.Println("Finish!")
}
