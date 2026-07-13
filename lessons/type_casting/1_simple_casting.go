package main

import (
	"fmt"
)

//func main() {
//	var a = 42
//	var b = float64(a) // int → float64
//	var c = int8(a)    // int → int8 (может быть потеря данных)
//
//	fmt.Printf("Тип a: %T\n", a)
//	fmt.Printf("Тип b: %T\n", b)
//	fmt.Printf("Тип c: %T\n", c)
//}

func main() {
	var x int64 = 300
	var y = int8(x) // 300 не помещается в int8 (диапазон -128..127)
	fmt.Println(y)  // выведет 44 (300 % 256 = 44)

	var f = 3.84
	var i = int(f) // дробная часть отбрасывается → 3
	fmt.Println(i)

	var u uint = 10
	var f2 = float32(u)
	fmt.Println(f2) // 10

	var negative int8 = -42
	var test = uint8(negative)
	fmt.Println(test) // 214
}
