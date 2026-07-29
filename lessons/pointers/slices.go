package main

type Slice struct {
	// pointerOnStart *на начало массива
	// len
	// cap
}

// Из-за копирования не видно изменений в main
//func main() {
//	nums := make([]int, 0, 10)
//	nums = append(nums, 1, 2, 3, 4)
//
//	append1(nums)
//
//	fmt.Println(nums)
//}
//
//func append1(nums []int) {
//	nums = append(nums, 1)
//}

//func main() {
//	nums := []int64{1, 2, 3, 4, 5}
//	fmt.Println(len(nums), cap(nums))
//
//	append1(&nums)
//	fmt.Println(len(nums), cap(nums))
//
//	fmt.Println(nums)
//}
//
//func append1(nums *[]int64) {
//	*nums = append(*nums, 1)
//}
