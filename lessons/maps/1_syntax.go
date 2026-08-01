package main

import "fmt"

func main() {
	//var m map[string]int

	m := make(map[string]int)

	m = map[string]int{
		"apple":  5,
		"pear":   3,
		"orange": 7,
	}

	m["cherry"] = 10

	fmt.Println(m)
}
