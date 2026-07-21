package main

import "fmt"

func main() {
	var first []int

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			first = append(first, i)
		}
	}

	fmt.Println(first)
}
