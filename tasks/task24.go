package tasks

import (
	"fmt"
	"math"
)

/*
 Задача 24. Разработать программу нахождения расстояния между двумя точками, которые
 представлены в виде структуры Point с инкапсулированными параметрами x,y и
 конструктором.
*/

type Point struct {
	x, y float64 // значения инкапсулированы в пакете
}

// "конструктор" возвращает указатель на структуру
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func TaskTwentyFour() float64 {
	a := NewPoint(5, 7)
	b := NewPoint(3, 4)
	dist := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
	fmt.Println(dist)
	return dist
}
