package main

import "fmt"

func main() {
	all := []int{5, 10, 15, 20, 25}
	sum := 0

	for i := 0; i < len(all); i++ {
		sum += all[i]
	}

	fmt.Println("Сумма среза:", sum)
}
