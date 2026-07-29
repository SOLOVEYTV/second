package main

import "fmt"

func main() {
	value := 5
	increment(&value)

	fmt.Println(value)
}

func increment(pointerOnValue *int) {
	*pointerOnValue++
}

// Из-за копирования не видно изменений в main
//func main() {
//	value := 5
//	increment(value)
//
//	fmt.Println(value)
//}
//
//func increment(pointerOnValue int) {
//	pointerOnValue++
//}
