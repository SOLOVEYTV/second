package main

import "fmt"

func main() {
	all := []int{10, 20, 40, 50}
	var two []int

	for idx := range all {
		if idx == 2 {
			two = append(two, 30)
		}
		two = append(two, all[idx])
	}
	fmt.Println(two)
}
