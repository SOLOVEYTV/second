package main

import (
	"fmt"
)

func main() {
	var age int64
	fmt.Println("Введите свой возвраст: ")
	fmt.Scan(&age)

	if age >= 18 || age <= 27 {
		fmt.Println("Вы несовершеннолетний")
	} else {
		fmt.Println("Вы совершеннолетний")
	}
}
