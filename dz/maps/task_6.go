package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 10, "c": 20, "d": 40}
	merged := MergeMaps(m1, m2)
	fmt.Println(merged) // map[a:1 b:12 c:23 d:40]
}

func MergeMaps(m1, m2 map[string]int) map[string]int {
	merged := make(map[string]int)
	for key, val := range m1 {
		merged[key] = val
	}
	for key, val := range m2 {
		merged[key] += val
	}

	return merged
}
