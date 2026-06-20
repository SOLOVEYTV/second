package main

import (
	"fmt"
)

type (
	Rubles  int64
	Dollars int64
	Kopecks int64
	Cents   int64
)

const rateRubToUsd float64 = 0.012826

func main() {
	var rub Rubles
	fmt.Println("Введите колличество рублей:")
	fmt.Scan(&rub)
	for rub < 0 {
		fmt.Println("Введите положительное число.")
		fmt.Scan(&rub)
	}
	var kop Kopecks
	fmt.Println("Введите колличество копеек:")
	fmt.Scan(&kop)
	for kop < 0 || kop > 99 {
		fmt.Println("Введите число от 0 до 99.")
		fmt.Scan(&kop)
	}

	cents := RubToCents(rub, kop)

	fmt.Println("-----------------------------------------------")
	fmt.Println("|", FormatRubles(rub, kop), " --> ", FormatDollars(cents), "|")
	fmt.Println("-----------------------------------------------")
}

func RubToCents(r Rubles, c Kopecks) Cents {
	usd := (float64(r) + (float64(c) / 100)) * rateRubToUsd

	return Cents(usd * 100)
}
func FormatRubles(r Rubles, kop Kopecks) string {
	return fmt.Sprint(r, " руб. ", kop, " коп.")
}

func FormatDollars(cents Cents) string {
	usd := Dollars(cents / 100)

	return fmt.Sprint(usd, " долл. ", cents%100, " цент.")
}
