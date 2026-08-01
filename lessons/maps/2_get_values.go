package main

import "fmt"

func main() {
	//var m map[string]int

	m := make(map[string]int)

	m = map[string]int{
		"apple":  5,
		"pear":   3,
		"orange": 7,
	}

	m["cherry"] = 10

	fmt.Println(m)

	val := m["apple"] // получаем значение по известному ключу

	fmt.Println(val)

	unknownVal := m["unknown"] // получаем дефолтное значения для типа

	fmt.Println(unknownVal)

	foundVal, ok := m["unknown"]
	if ok {
		fmt.Println("Значение найдено:", foundVal)
	} else {
		fmt.Println("Ключ не найден")
	}
}
