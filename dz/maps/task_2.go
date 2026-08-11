package main

import "fmt"

func main() {
	set := make(map[int]bool)
	Add(&set, 10)
	Add(&set, 20)
	Add(&set, 10)
	fmt.Println(Contains(set, 10))
	fmt.Println(Contains(set, 30))
	Remove(&set, 10)
	fmt.Println(Contains(set, 10))
}

func Add(set *map[int]bool, val int) {
	(*set)[val] = true
}

func Remove(set *map[int]bool, val int) {
	delete(*set, val)
}

func Contains(set map[int]bool, val int) bool {
	_, ok := set[val]
	return ok
}
