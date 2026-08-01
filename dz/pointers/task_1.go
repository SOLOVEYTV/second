package main

import "fmt"

func main() {
	x := 6
	Double(&x)

	fmt.Println(x)
}

func Double(val *int) {
	*val *= 2
}
