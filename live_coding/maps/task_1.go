package main

import "fmt"

// Задача 1. Сумма значений мапы
// Напишите функцию SumValues(m map[string]int) int, которая возвращает сумму всех значений в мапе.

func main() {
	result := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	sum := SumValues(result)

	fmt.Println(sum)
}

func SumValues(m map[string]int) int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}
