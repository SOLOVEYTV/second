package main

import "fmt"

func main() {
	all := []int{-5, 10, -3, 0, 8, -1, 4}
	var first []int
	for _, val := range all {
		if val > 0 {
			first = append(first, val)
		}
	}
	fmt.Println(first)
}
