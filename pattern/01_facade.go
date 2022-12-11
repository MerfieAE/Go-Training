package pattern

import (
	"fmt"
	"math/rand"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
Проблема:
Минимизировать зависимость подсистем некоторой сложной системы и обмен информацией между ними.
Решение:
Фасад — простой интерфейс для работы со сложным фреймворком. Фасад не имеет всей функциональности фреймворка,
но зато скрывает его сложность от клиентов.
В примере рассматривается сервис такси-автопилота, который довозит пассажира в указанное место.
Сервис состоит из функции выбора локации DriveToLocation, клиентского кошелька ClientBudget, и робота TaxiPilot, который
принимает решение о поездке.
плюсы:
	+Изолирует клиентов от компонентов сложной подсистемы.
минусы:
	-Фасад рискует стать божественным объектом, привязанным ко всем классам программы. (объект, который хранит
		 в себе «слишком много» или делает «слишком много». )
Фасад - интерфейс пользователя, который просто передает управление внутренним элементам системы торгового автомата
*/

//DriveToLocation - выбор локации пользователем
type DriveToLocation struct {
	LocationPrice int
	LocationName  string
}

func (d *DriveToLocation) SetLocation(m map[int]string) {
	//Выбирает рандомное значение из мапы
	for key, value := range m {
		d.LocationPrice = key
		d.LocationName = value
	}
}

//ClientBudget - бюджет клиента
type ClientBudget struct {
	Budget int
}

func (c *ClientBudget) SetBudget() {
	c.Budget = rand.Intn(1000)
}

//TaxiPilot - робот-такси
type TaxiPilot struct {
	Payment bool
}

func (p *TaxiPilot) PaymentSuccess(cost, price int) {
	if cost > price {
		p.Payment = true
	} else {
		p.Payment = false
	}
}

//Фасад
type PrepareTrip struct {
	destination *DriveToLocation
	budget      *ClientBudget
	management  *TaxiPilot
}

func NewStart() *PrepareTrip {
	return &PrepareTrip{
		destination: &DriveToLocation{},
		budget:      &ClientBudget{},
		management:  &TaxiPilot{},
	}
}

//Реализация паттерна
func (p *PrepareTrip) Start() {
	fmt.Println("Prepare pilot for a drive, the robot is going to check your wallet...")
	p.destination.SetLocation(map[int]string{54: "Ближайшее метро", 98: "Ближайшее кино", 102: "Ближайший отель", 35: "Ближайший ресторан", 79: "Ближайший автопарк"})
	p.budget.SetBudget()
	p.management.PaymentSuccess(p.budget.Budget, p.destination.LocationPrice)
	fmt.Printf("Your destination is %s, the cost of trip is %v, budget is %v\n", p.destination.LocationName, p.destination.LocationPrice, p.budget.Budget)
	if p.management.Payment == true {
		fmt.Println("Your budget is ok, we start trip...")
	} else {
		fmt.Println("Sorry, you haven't enough money for the trip, try another time...")
	}
}
