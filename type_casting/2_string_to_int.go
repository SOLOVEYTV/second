package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "7324657635245634"

	num1, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num1)

	s1 := "1101011"
	num2, _ := strconv.ParseInt(s1, 2, 64)
	fmt.Println(num2)

	// 16-тиричная система исчисления
	hex := "FF"

	val, _ := strconv.ParseInt(hex, 16, 64)
	fmt.Println(val)
}
