package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 5, 5, 5)) // 6
	fmt.Println(sum(10, 20))           // 30
	fmt.Println(sum())                 // 0
}
