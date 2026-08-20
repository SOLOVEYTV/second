package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 1,
		"d": 3,
		"e": 2,
	}
	DeleteByValue(&m, 2)
	fmt.Println(m)
}

func DeleteByValue(m *map[string]int, value int) {
	for idx, val := range *m {
		if val == value {
			delete(*m, idx)
		}
	}
}
