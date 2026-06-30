package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}

	return a / b, nil
}

func main() {
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}
}

//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//
//	return // возвращает x и y
//}
//
//func main() {
//	a, b := split(17)
//	fmt.Println(a, b) // 7 10 (примерно)
//}
