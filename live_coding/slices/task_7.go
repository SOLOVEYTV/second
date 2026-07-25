package main

import (
	"fmt"
)

// Линейный поиск элемента
//
// Напишите функцию search(arr []int, req int) int, которая ищет значение req в срезе arr
// и возвращает индекс первого вхождения. Если элемент не найден — верните -1.
//
// Пример:
// arr := []int{13, 22, 31, 14, 65, 456, 37}

func main() {
	arr := []int{13, 22, 31, 14, 65, 456, 37}

	index := search(arr, 456)
	fmt.Println(index)
}

func search(arr []int, req int) int {
	var index int
	for idx, val := range arr {
		if req == val {
			index = idx
			break
		} else {
			index = -1
		}
	}
	return index
}

//func search(arr []int, req int) int {
//	for idx, val := range arr {
//		if req == val {
//			return idx
//		}
//	}
//
//	return -1
//}
