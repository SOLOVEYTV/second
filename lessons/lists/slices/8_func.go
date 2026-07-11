package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}

	incrementCopy(list) // ничего не меняется, потому что в main не видно изменений, которые были сделаны над локальной копией внутри функции

	list1 := incrementInitLen(list)
	list2 := incrementInitLen(list1)

	fmt.Println(list2)
}

func incrementCopy(array []int) { // неправильно, теряем модифицируемые данные, так как работаем с локальной копией
	for _, value := range array {
		value++
	}
}

func incrementInitCap(array []int) []int { // правильно, выделяем заранее емкость
	result := make([]int, 0, len(array))

	for _, value := range array {
		value++
		result = append(result, value)
	}

	return result
}

func incrementInitLen(array []int) []int { // тоже правильно, выделяем заранее длину == емкость
	result := make([]int, len(array))

	for i, value := range array {
		value++
		result[i] = value
	}

	return result
}
