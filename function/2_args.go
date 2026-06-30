package main

import (
	"fmt"
)

func add(a int, b int) int {
	sum := a + b

	return sum
}

// можно сократить, если типы совпадают:
func multiply(a, b int) int {
	if a == 0 {
		return 0
	}

	return a * b
}

func main() {
	res := add(5, 3) // 8

	product := multiply(4, 2) // 8

	fmt.Println(res, product)
}
