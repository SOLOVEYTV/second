package main

import "fmt"

func main() {
	var stack []int
	stack = Push(stack, 10)
	stack = Push(stack, 20)
	stack = Push(stack, 30)

	newStack, elem, ok := Pop(stack)
	newStack, elem, ok = Pop(newStack)
	newStack, elem, ok = Pop(newStack)
	newStack, elem, ok = Pop(newStack)
	fmt.Println(newStack, elem, ok)

}

func Push(stack []int, elem int) []int {
	stack = append(stack, elem)
	fmt.Println("Добавлено:", elem, " текущий срез:", stack)

	return stack
}

func Pop(stack []int) ([]int, int, bool) {
	var ok bool
	var newStack []int
	if len(stack) > 0 {
		ok = true
		elem := stack[len(stack)-1]
		newStack = stack[:len(stack)-1]
		fmt.Println("Извлечено:", elem, " очередь:", newStack, " статус:", ok)

		return newStack, elem, ok
	}

	ok = false
	fmt.Println("Извлекать нечего!", " статус:", ok)

	return newStack, 0, ok
}
