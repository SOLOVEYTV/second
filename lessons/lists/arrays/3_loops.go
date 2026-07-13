package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	for idx, val := range arr {
		fmt.Printf("Индекс %d, значение %d\n", idx, val)
	}

	for _, val := range arr {
		fmt.Printf("Значение %d\n", val)
	}
}
