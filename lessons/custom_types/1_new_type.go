package main

import (
	"fmt"
)

type (
	Celsius    float64 // новый тип на основе float64
	Fahrenheit float64
)

func main() {
	var t Celsius = 25.0
	var f Fahrenheit = 78.8

	//t = f          // Ошибка: cannot use f (type Fahrenheit) as type Celsius
	t = Celsius(f) // явное приведение

	fmt.Println(t, f)
}
