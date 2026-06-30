package main

import (
	"fmt"
)

func main() {
	// присваиваем анонимную функцию переменной
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(3, 4)) // 7
}
