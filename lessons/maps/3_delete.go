package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  5,
		"pear":   3,
		"orange": 7,
	}

	m["cherry"] = 10

	delete(m, "pear")

	fmt.Println(m)

	fmt.Println(len(m))
}
