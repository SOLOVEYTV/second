package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)     // длина=5, ёмкость=5, элементы равны нулю [0,0,0,0,0]
	s2 := make([]int, 3, 5) // длина=3, ёмкость=5 [0,0,0, , ]

	fmt.Println(s, s2)
}
