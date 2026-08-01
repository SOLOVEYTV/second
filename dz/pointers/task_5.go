package main

import "fmt"

func main() {
	var p *int
	val := SafeDeref(p, 42)
	fmt.Println(val)

	x := 100
	val2 := SafeDeref(&x, 0)
	fmt.Println(val2)
}

func SafeDeref(p *int, defaultVal int) int {
	if p != nil {
		return *p
	}

	return defaultVal
}
