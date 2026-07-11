package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}    // len=3, cap=3
	s = append(s, 4)       // [1, 2, 3, 4] // len=4, cap=6
	s = append(s, 5, 6, 7) // [1, 2, 3, 4, 5, 6, 7] // len=7, cap=12

	fmt.Println(s)
}
