package main

import "fmt"

type (
	Kilometers float64
	Miles      float64
)

const rateKmToMiles = 1.60934

func main() {
	var km Kilometers
	var mil Miles
	fmt.Println("Введите дистанцию в КМ:")
	fmt.Scan(&km)
	for km <= 0 {
		fmt.Println("Введите положительное число.")
		fmt.Scan(&km)
	}
	mil = KmToMiles(km)
	fmt.Printf("%.2f миль \n", mil)

	fmt.Println("Введите дистанцию в Милях:")
	fmt.Scan(&mil)
	for mil <= 0 {
		fmt.Println("Введите положительное число.")
		fmt.Scan(&mil)
	}
	km = MilesToKm(mil)
	fmt.Printf("%.2f км", km)
}

func KmToMiles(km Kilometers) Miles {
	mi := km / rateKmToMiles
	return Miles(mi)
}
func MilesToKm(mi Miles) Kilometers {
	km := mi * rateKmToMiles
	return Kilometers(km)
}
