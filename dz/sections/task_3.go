package main

import "fmt"

func main() {
	src := []string{"a", "b", "c", "d", "e"}
	dst := make([]string, 3)
	copy(dst, src)
	fmt.Println(dst)
}
