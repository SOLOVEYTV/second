package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  5,
		"pear":   3,
		"orange": 7,
	}

	for index, val := range m {
		fmt.Println(index, val)
	}
}
