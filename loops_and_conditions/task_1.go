package main

import "fmt"

//Задача 1 (очень лёгкая)
//Вывести все числа от 1 до 10, но пропустить число 5. Использовать continue.

func main() {
	for count := 1; count <= 10; count++ {
		if count == 5 {
			continue
		}

		fmt.Println(count)
	}
}
