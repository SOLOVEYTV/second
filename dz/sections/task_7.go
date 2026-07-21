package main

import "fmt"

func main() {
	all := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(all)/2; i++ {
		j := len(all) - i - 1
		all[i], all[j] = all[j], all[i]
	}

	fmt.Println(all)
}
