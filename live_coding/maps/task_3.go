package main

import "fmt"

// Задача 7. Срез уникальных значений из мапы

// Напишите функцию UniqueValues(m map[string]int) []int, которая возвращает срез всех уникальных значений (без повторений). Порядок не важен.

func main() {
	original := map[string]int{
		"Alice": 1,
		"Bob":   2,
		"Carol": 2,
		"Dave":  3,
		"Eve":   5,
		"Nick":  1,
	}
	unic := UniqueValues(original)
	fmt.Println(unic)
}

func UniqueValues(m map[string]int) []int {
	result := make([]int, 0)
	temp := make(map[int]bool)
	for _, val := range m {
		if !temp[val] {
			temp[val] = true
			result = append(result, val)
		}
	}

	return result
}
