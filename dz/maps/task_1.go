package main

import "fmt"

func main() {
	original := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 25,
		"Dave":  30,
		"Eve":   35,
	}
	inverted := InvertMap(original)
	fmt.Println(inverted)
}

func InvertMap(m map[string]int) map[int][]string {
	result := make(map[int][]string)
	for key, val := range m {
		result[val] = append(result[val], key)
	}

	return result
}
