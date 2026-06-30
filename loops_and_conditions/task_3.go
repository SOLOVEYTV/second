package main

import "fmt"

//Задача 3 (лёгкая)
//Дано целое число. Определить, является ли оно простым (делится без остатка только на 1 и на себя). Использовать цикл для проверки делителей.

func main() {
	var user_number int
	fmt.Print("Введите число: ")
	fmt.Scan(&user_number)
	var remainder = 0

	for counter := 1; counter <= user_number; counter++ {
		if user_number%counter == 0 {
			remainder++
			if remainder > 2 {
				break
			}
		}
	}
	if remainder == 2 {
		fmt.Println("Простое число")
	} else {
		fmt.Println("Не простое число!")
	}
}
