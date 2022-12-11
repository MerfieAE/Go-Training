package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


Строитель (англ. Builder) — порождающий шаблон проектирования предоставляет способ создания составного объекта.
Отделяет конструирование сложного объекта от его представления так, что в результате одного и того
же процесса конструирования могут получаться разные представления.

Проблема:
    Инициализация очень сложного, большого объекта со множеством параметров инициализации. Использовать один конструктор
    с множеством параметров - плохо (телескопический конструктор - анти-паттерн)

Решение:
    Паттерн Строитель предлагает вынести конструирование объекта за пределы его собственного класса, поручив это дело
    отдельным объектам, называемым строителями.

плюсы:
    +Позволяет создавать продукты пошагово.
    +Позволяет использовать один и тот же код для создания различных продуктов.
    +Изолирует сложный код сборки продукта от его основной бизнес-логики.

минусы:
    -Усложняет код программы из-за введения дополнительных классов.
    -Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.

    В качесте примера реализации создадим автомобиль.
*/

type Car struct {
	Engine int
	Body   string
	Wheels string
}

type CarBuilder interface {
	Engine(val int) CarBuilder
	Body(val string) CarBuilder
	Wheels(val string) CarBuilder

	Build() Car
}

type carBuilder struct {
	engine int
	body   string
	wheels string
}

func NewCarBuilder() CarBuilder {
	return carBuilder{}
}

func (c carBuilder) Engine(val int) CarBuilder {
	c.engine = val
	return c
}
func (c carBuilder) Body(val string) CarBuilder {
	c.body = val
	return c
}
func (c carBuilder) Wheels(val string) CarBuilder {
	c.wheels = val
	return c
}
func (c carBuilder) Build() Car {
	return Car{
		Engine: c.engine,
		Body:   c.body,
		Wheels: c.wheels,
	}
}
