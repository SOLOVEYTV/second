package main

import (
	"fmt"
)

func main() {
	// Пустой срез
	var s1 []int  // nil-срез (длина=0, ёмкость=0, не выделено памяти)
	s2 := []int{} // не nil, но пустой (длина=0, ёмкость=0)

	// Срез с начальными значениями
	s3 := []string{"Go", "Python", "Java"}
	s4 := []float64{1.2, 3.4, 5.6}

	fmt.Println(s1, s2)
	fmt.Println(len(s3), cap(s3), len(s4), cap(s3))
}
