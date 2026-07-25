package main

import "fmt"

func main() {
	var all []int
	input := 1
	for input != 0 {
		fmt.Println("Введите число (0 для выхода):")
		fmt.Scan(&input)
		if input != 0 {
			all = append(all, input)
		}
	}
	fmt.Println(all, "Длина:", len(all))
}
