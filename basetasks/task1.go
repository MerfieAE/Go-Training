package basetasks

/*
Задача 1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	Kind string
}

func TaskOne() *Action {
	return &Action{
		Human: Human{Name: "John Weak", Age: 21},
		Kind:  "get passport",
	}
}
