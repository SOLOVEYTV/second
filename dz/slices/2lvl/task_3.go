package main

import "fmt"

func main() {
	var q []int
	q = Enqueue(q, 10)
	q = Enqueue(q, 20)
	q = Enqueue(q, 30)

	newq, elem, ok := Dequeue(q)
	newq, elem, ok = Dequeue(newq)
	newq, elem, ok = Dequeue(newq)
	newq, elem, ok = Dequeue(newq)
	fmt.Println(newq, elem, ok)

}

func Enqueue(q []int, elem int) []int {
	q = append(q, elem)
	fmt.Println("Добавлено:", elem, " текущий срез:", q)

	return q
}

func Dequeue(q []int) ([]int, int, bool) {
	var ok bool
	var newq []int
	if len(q) > 0 {
		ok = true
		elem := q[0]
		newq = q[1:]
		fmt.Println("Извлечено:", elem, " очередь:", newq, " статус:", ok)

		return newq, elem, ok
	}

	ok = false
	fmt.Println("Извлекать нечего!", " статус:", ok)

	return newq, 0, ok
}
