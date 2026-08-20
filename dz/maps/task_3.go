package main

import "fmt"

func main() {
	words := []string{"cat", "dog", "elephant", "bat", "lion", "ant"}
	grouped := GroupByLength(words)

	fmt.Println(grouped)
}

func GroupByLength(words []string) map[int][]string {
	result := make(map[int][]string)
	for _, word := range words {
		length := len(word)
		result[length] = append(result[length], word)
	}

	return result
}
