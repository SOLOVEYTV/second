package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	Chunks := Chunk(s, 4)

	fmt.Println(Chunks)
}

func Chunk(s []int, size int) [][]int {
	result := make([][]int, 0, len(s)/size+1)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}

		result = append(result, s[i:end])

	}

	return result
}
