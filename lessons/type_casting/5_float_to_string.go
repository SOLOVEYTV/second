package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 3.14159

	s := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println(s) // "3.14"
}
