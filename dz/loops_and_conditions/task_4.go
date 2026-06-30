package main

import "fmt"

//Задача 4 (средняя)
//Пользователь вводит 5 оценок (целые числа от 0 до 100). Программа должна вывести:
//Средний балл (как float64)
//Количество пятёрок (оценка == 5)
//Массив оценок не хранить – обрабатывать каждую оценку сразу после ввода.

func main() {

	var (
		sum = 0.0
		col = 0
	)

	for count := 1; count <= 5; count++ {
		var numbers = 0
		fmt.Print("Введите ", count, "е число: ")
		fmt.Scan(&numbers)
		for numbers < 0 || numbers > 100 {
			fmt.Print("Число не входит в диапозон от 0 до 100!", "\nВведите ", count, "е число: ")
			fmt.Scan(&numbers)
		}

		sum = sum + float64(numbers)
		if numbers == 5 {
			col = col + 1
		}
	}
	fmt.Println("Ср. балл: ", sum/5)
	fmt.Println("Кол-во пятёрок: ", col)
}
