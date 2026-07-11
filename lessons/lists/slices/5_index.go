package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:4] // элементы с индекса 1 по 3 (не включая 4) → [2, 3, 4]

	fmt.Println(slice)
}
