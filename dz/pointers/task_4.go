package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	AppendElement(&nums, 4)
	fmt.Println(nums)
}

func AppendElement(s *[]int, val int) {
	*s = append(*s, val)
}
