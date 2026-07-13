package main

import (
	"fmt"
)

func main() {
	score := 85
	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 75 {
		fmt.Println("Хорошо")
	} else if score >= 60 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Неудовлетворительно")
	}
}
