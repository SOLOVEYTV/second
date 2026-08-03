package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "apple orange orange orange orange apple banana cherry"

	words := strings.Fields(text)
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}

	fmt.Println(count)
}
