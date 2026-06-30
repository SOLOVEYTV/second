package main

import "fmt"

// Программа запрашивает у пользователя целое число n.
// Вывести таблицу умножения для этого числа от 1 до 10, но если результат умножения кратен 3 – вместо числа вывести "boom".

func main() {
	var user_number int
	fmt.Print("Введите число: ")
	fmt.Scan(&user_number)
	var result = 1

	for counter := 1; counter <= 10; counter++ {
		result = user_number * counter
		if result%3 == 0 {
			fmt.Println("BOOM")
		} else {
			fmt.Println(result)
		}
	}
}
