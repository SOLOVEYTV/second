package main

//Задача 2 (лёгкая)
//Найти сумму всех чётных чисел от 1 до N. N вводит пользователь.

import "fmt"

func main() {
	var user_number int
	fmt.Print("Введите N: ")
	fmt.Scan(&user_number)

	sum := 0

	for i := 2; i <= user_number; i += 2 {
		sum += i
		fmt.Println(i)
	}
	fmt.Println("Сумма чётных чисел:", sum)
}
