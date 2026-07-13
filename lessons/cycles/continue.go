package main

import (
	"fmt"
)

func main() {
	for count := 1; count <= 5; count++ {
		if count == 4 {
			continue
		}

		fmt.Println(count)
	}
}
