package main

import "fmt"

// Удаление ключей из мапы по срезу ключей

// Условие:
// Напишите функцию DeleteKeys(m *map[string]int, keys []string) int, которая удаляет из мапы m все ключи,
// присутствующие в срезе keys. Возвращает количество удалённых ключей. Если ключа нет – он игнорируется.
// Мапа изменяется по указателю.

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	keys := []string{"a", "c", "e"}

	deleted := DeleteKeys(&m, keys)

	fmt.Println(deleted) // 2
	fmt.Println(m)       // map[b:2 d:4]
}

func DeleteKeys(m *map[string]int, keys []string) int {
	var sum int
	for _, val := range keys {
		if _, ok := (*m)[val]; ok {
			delete(*m, val)
			sum++
		}
	}

	return sum
}
