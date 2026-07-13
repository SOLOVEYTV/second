package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 2)

	copied := copy(dst, src) // копирует первые 2 элемента из src в dst
	fmt.Println(dst)         // [1, 2]
	fmt.Println(copied)      // 2
}
