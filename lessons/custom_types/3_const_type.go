package main

import (
	"fmt"
)

type (
	Weekday int
	Month   int
)

const (
	Sunday  Weekday = 1
	Monday  Weekday = 2
	Tuesday Weekday = 3
	// ...

	January  Month = 1
	February Month = 2
	// ...
)

func main() {
	today := Sunday
	fmt.Println(today)
}
