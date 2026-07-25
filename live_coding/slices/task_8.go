package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 3, 3, 3, 4}
	newArr := dubFind(arr)
	fmt.Println(newArr)
}

func dubFind(arr []int) []int {
	var result []int
	for _, val := range arr {
		var found bool

		for _, valResult := range result {
			if valResult == val {
				found = true
				break
			}
		}

		if found == false {
			result = append(result, val)
		}
	}

	return result
}
