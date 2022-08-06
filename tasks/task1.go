package tasks

/*
Задача 1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

import "fmt"

type Human struct {
	Name      string
	Age       int
	IsWorking bool
}

func (h *Human) getHuman() string {
	h.Name = "John"
	h.Age = 23
	h.IsWorking = true
	return fmt.Sprintf("Human name is %s, he is %v, he is working: %t.", h.Name, h.Age, h.IsWorking)
}

type Action struct {
	Verb string
	Human
}

func (a *Action) getAction() string {
	a.Verb = "progressive"
	return fmt.Sprintf("He is very %s!", a.Verb)
}

func TaskOne() {
	action := Action{}
	fmt.Println(action.getHuman(), action.getAction())
}
