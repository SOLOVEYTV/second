package main

import (
	"fmt"
)

// Минимум и максимум в срезе
//
// Напишите функцию minMaxNum(numbers []int) (int, int), которая принимает срез целых чисел
// и возвращает минимальное и максимальное значения.
//
//Пример:
//numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//min, max := minMaxNum(numbers)
//min => 1, max => 10

func main() {
	numbers := []int{-1, -2, -3}
	min, max := minMaxNum(numbers)

	fmt.Printf("min: %d, max: %d", min, max)
}

func minMaxNum(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for _, val := range numbers {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return min, max
}
