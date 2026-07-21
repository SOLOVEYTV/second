package main

import "fmt"

func main() {
	var index int
	all := []int{3, 7, 2, 9, 1, 6}
	maximum := all[0]

	for idx, val := range all {
		if val > maximum {
			maximum = val
			index = idx
		}
	}

	fmt.Printf("Максимум: %d, индекс: %d", maximum, index)
}
