package main

import (
	"fmt"
)

//func main() {
//	count := 10
//	num := 1
//
//	for range count {
//		fmt.Printf("Значение: %d\n", num)
//
//		num++
//	}
//}

func main() {
	s := "Hello"
	for index, value := range s {
		fmt.Printf("Позиция: %d, Символ: %c\n", index, value)
	}
}

// 0101100110100100
// [10101101011001, 101010101010101, 101010101010, 1010101010, 010101010010]
