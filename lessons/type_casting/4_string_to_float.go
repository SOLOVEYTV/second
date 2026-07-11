package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3.14159"

	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f) // 3.14159
}
