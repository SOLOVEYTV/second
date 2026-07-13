package main

import (
	"fmt"
)

//func main() {
//	s := []int{1, 2, 3}
//	fmt.Println("До:", s, len(s), cap(s)) // 3, 3
//
//	s = append(s, 4, 5)
//	fmt.Println("После:", s, len(s), cap(s)) // 5, 6
//}

func main() {
	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
	weekend := days[5:7] // Сб, Вс

	fmt.Println(weekend)
	fmt.Println(len(weekend), cap(weekend))

	fmt.Println(days)

	weekend[0] = "jdsnjknsdj" // изменяем элементы в обоих массивах

	fmt.Println(days)

	days = append(days, "Воварраваав") // выделяем новую память, days и weekend больше не ссылаются на одну область памяти

	weekend[0] = "Сб" // пытаемся вернуть субботу

	fmt.Println(days)    // не суббота
	fmt.Println(weekend) // суббота

	//[Сб Вс]
	//2 2
	//[Пн Вт Ср Чт Пт Сб Вс]
	//[Пн Вт Ср Чт Пт jdsnjknsdj Вс]
	//[Пн Вт Ср Чт Пт jdsnjknsdj Вс Воварраваав]
	//[Сб Вс]
}
