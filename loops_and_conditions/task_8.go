package main

import "fmt"

//Задача 8 (средняя)
//С помощью вложенных циклов вывести следующий узор (для n=5):

func main() {
	var user_number int
	fmt.Print("Введите число: ")
	fmt.Scan(&user_number)
	for user_number <= 0 {
		fmt.Println("*ПУСТО*\nВведите число больше 0:")
		fmt.Scan(&user_number)
	}
	var star = "*"
	for count := 1; count <= user_number; count++ {
		fmt.Println(star)
		star = star + "*"
	}
}
