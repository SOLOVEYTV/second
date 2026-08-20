package main

import "fmt"

// Удаление из мапы по значениям (срез значений)

// Условие:
// Напишите функцию DeleteByValues(m *map[string]int, values []int) int, которая удаляет все пары, у которых значение присутствует в срезе values. Возвращает количество удалённых элементов.

// Пример:

func main() {
	m := map[string]int{"a": 10, "b": 20, "c": 10, "d": 30}
	vals := []int{10, 30}

	deleted := DeleteByValues(&m, vals)

	fmt.Println(deleted) // 3 (удалены a, c, d)
	fmt.Println(m)       // map[b:20]
}

func DeleteByValues(m *map[string]int, values []int) int {
	valMap := make(map[int]bool)
	for _, value := range values {
		valMap[value] = true
	}

	keysForDelete := []string{}
	for key, value := range *m {
		if valMap[value] == true {
			keysForDelete = append(keysForDelete, key)
		}
	}

	for _, value := range keysForDelete {
		delete(*m, value)
	}

	return len(keysForDelete)
}
