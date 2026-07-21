package main

import "fmt"

func main() {
	first := [5]int{10, 20, 30, 40, 50}
	second := first[1:4]

	fmt.Println(second)
}
