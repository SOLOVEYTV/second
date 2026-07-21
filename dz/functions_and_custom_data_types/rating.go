package main

import (
	"fmt"
)

type (
	Rating int
)

const (
	Terrible  Rating = 1
	Bad       Rating = 2
	Fair      Rating = 3
	Good      Rating = 4
	Excellent Rating = 5
)

func main() {
	var r Rating
	fmt.Println("Введите оценку от 1 до 5: ")
	fmt.Scan(&r)
	if Validate(r) {
		fmt.Println(Grade(r))
	} else {
		fmt.Println("Некорректная оценка")
	}

}

func Validate(r Rating) bool {
	return r >= Terrible && r <= Excellent
}

func Grade(r Rating) string {
	switch {
	case r == Terrible:
		return "Ужасно"
	case r == Bad:
		return "Плохо"
	case r == Fair:
		return "Удовлетворительно"
	case r == Good:
		return "Хорошо"
	case r == Excellent:
		return "Отлично"
	default:
		return "Некорректная оценка"
	}
}
