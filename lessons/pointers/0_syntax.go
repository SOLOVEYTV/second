package main

import "fmt"

func main() {
	var value int = 10

	var pointer *int = &value // pointer 0xc000234243434
	*pointer = 20

	fmt.Println(pointer)
	fmt.Println(*pointer)
}
