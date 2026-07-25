package main

import "fmt"

func main() {
	all := []int{2, 5, 8, 11, 14, 17, 20, 23}
	var even int
	var odd int
	for _, val := range all {
		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Printf("Чётных: %d, Нечётных: %d", even, odd)
}
