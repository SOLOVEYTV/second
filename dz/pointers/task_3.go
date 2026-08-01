package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	ToUpper(&str)

	fmt.Println(str)
}

func ToUpper(s *string) {
	*s = strings.ToUpper(*s)
}
