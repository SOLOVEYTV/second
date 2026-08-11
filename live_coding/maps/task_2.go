package main

import "fmt"

// Задача 2. Ключ с максимальным значением

// Напишите функцию KeyWithMaxValue(m map[string]int) string, которая возвращает ключ, у которого значение максимальное. Если мапа пуста, вернуть пустую строку.

func main() {
	result := map[string]int{
		"one":   -1,
		"two":   -2,
		"three": -3,
	}
	maxKey := KeyWithMaxValue(result)

	fmt.Println(maxKey)
}

func KeyWithMaxValue(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}
	var maxKey string
	var maxVal int
	for key, val := range m {
		if maxKey == "" || val > maxVal {
			maxVal = val
			maxKey = key
		}
	}

	return maxKey
}
