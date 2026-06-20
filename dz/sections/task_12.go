package main

import "fmt"

func main() {
	all := make([]int, 0, 3)
	value := 0
	for i := 0; i < 5; i++ {
		value = value + 10
		all = append(all, value)
		println("len=", len(all), " ", "cap=", cap(all))
	}
	fmt.Println(all)
}
