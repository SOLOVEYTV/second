package main

import "fmt"

// Задача 14. Разность мап (ключи первой, отсутствующие во второй)

// Условие:
// Напишите функцию MapDifference(a, b map[string]int) map[string]int, которая возвращает новую мапу,
// содержащую пары из a, ключи которых не присутствуют в b. Значения берутся из a.

func main() {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := map[string]int{"b": 10, "d": 4}

	diff := MapDifference(a, b)

	fmt.Println(diff) // map[a:1 c:3]
}

func MapDifference(a, b map[string]int) map[string]int {
	temp := make(map[string]int)
	for key, val := range a {
		if _, ok := b[key]; !ok {
			temp[key] = val
		}
	}

	return temp
}
