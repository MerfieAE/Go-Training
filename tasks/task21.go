package tasks

/*
 Задача 21.	Реализовать паттерн «адаптер» на любом примере.
*/

import "fmt"

type celsius float64

func FahrenheitToCelsius(fahrenheit float64) (result celsius) {
	result = celsius((fahrenheit - 32) * 5 / 9)
	return
}

func KelvinToCelsius(kelvin float64) (result celsius) {
	result = celsius(kelvin - 273.15)
	return
}

func TaskTwentyOne() {
	fahrenheit := 9.0
	kelvin := 15.0
	fmt.Println(FahrenheitToCelsius(fahrenheit))
	fmt.Println(KelvinToCelsius(kelvin))
}
