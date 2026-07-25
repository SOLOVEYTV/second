package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8, 10}
	result := MergeSorted(a, b)

	fmt.Println(result)
}

func MergeSorted(a, b []int) []int {
	var newSlice []int
	j, i := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			newSlice = append(newSlice, a[i])
			i++
		} else {
			newSlice = append(newSlice, b[j])
			j++
		}

	}
	for j < len(b) {
		newSlice = append(newSlice, b[j])
		j++
	}
	for i < len(a) {
		newSlice = append(newSlice, a[i])
		i++
	}

	return newSlice
}
