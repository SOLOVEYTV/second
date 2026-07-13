package main

import (
	"fmt"
)

//func main() {
//	// сумма всех элементов в массиве
//	nums := [5]int{3, 7, 1, 9, 4}
//	sum := 0
//
//	for _, v := range nums {
//		sum += v
//	}
//
//	fmt.Println("Сумма:", sum) // 24
//}

// Поиск максимального элемента
//func main() {
//	arr := [6]int{5, 2, 9, 1, 7, 3}
//
//	maxValue := arr[0] // предполагаем, что массив не пуст
//	for _, v := range arr {
//		if v > maxValue {
//			maxValue = v
//		}
//	}
//
//	fmt.Println("Максимум:", maxValue) // 9
//}

// Подсчёт количества чётных чисел
func main() {
	numbers := [7]int{2, 5, 8, 11, 14, 17, 20}

	evenCount := 0
	for _, v := range numbers {
		if v%2 == 0 {
			evenCount++
		}
	}
	fmt.Println("Чётных:", evenCount) // 4
}
