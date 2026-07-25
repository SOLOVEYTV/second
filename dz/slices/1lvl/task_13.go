package main

import "fmt"

func main() {
	all := []int{1, 2, 3, 4, 5}
	copy(all, all[1:])
	all[4] = 0
	fmt.Println(all)
}
