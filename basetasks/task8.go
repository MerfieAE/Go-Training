package basetasks

/*
 Задача 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в
 1 или 0.
*/

import "fmt"

func TaskEight() {
	var number int64 = 100
	var numberOfBite int64 = 5
	switchToOne := false
	switch switchToOne {
	case true:
		// | поразрядная дизъюнкция, Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
		number |= 1 << numberOfBite
	default:
		// ^ поразрядное исключающее ИЛИ. Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
		number &^= 1 << numberOfBite
	}

	fmt.Println("result is - ", number)
}
