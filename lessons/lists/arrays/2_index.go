package main

import (
	"fmt"
)

func main() {
	arr := [4]int{10, 20, 30, 40}

	fmt.Println(arr[0]) // 10
	fmt.Println(arr[3]) // 40

	arr[1] = 99 // изменяем второй элемент

	fmt.Println(arr) // [10, 99, 30, 40]

	arr1 := [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(len(arr1)) // 7
}
