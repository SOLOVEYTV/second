package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 42

	s := strconv.Itoa(n)
	fmt.Println(s) // "42"

	bin := strconv.FormatInt(int64(n), 2)
	fmt.Println(bin) // "101010"
}
