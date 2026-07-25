package main

import "fmt"

//Среднее арифметическое
//
//Напишите функцию midNum(numbers []int32) float32, которая принимает срез чисел
//и возвращает их среднее арифметическое (сумму, делённую на количество элементов).
//
//Пример:
//numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 11}
//midNum(numbers) => 5.222...

func main() {
	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 11}
	midValue := midNum(numbers)

	fmt.Println(midValue)
}

func midNum(numbers []int32) float32 {
	var sum int32
	for _, val := range numbers {
		sum += val
	}
	mid := float32(sum) / float32(len(numbers))

	return mid
}
