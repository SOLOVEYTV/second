package main

import (
	"fmt"
)

func main() {
	for count := 1; count <= 5; count++ {
		fmt.Println(count)

		if count == 4 {
			break
		}
	}
}
