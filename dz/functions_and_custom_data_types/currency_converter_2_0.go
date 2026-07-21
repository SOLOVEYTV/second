package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeResponse struct {
	Result string             `json:"result"`
	Rates  map[string]float64 `json:"rates"`
}

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
	rate, err := GetUsdRate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Актуальный курс полученный с API: ", float64(rate))

	cents := RubToCentsWithExternalRate(rub, kop, rate)

	fmt.Println("-----------------------------------------------")
	fmt.Println("|", FormatRubles(rub, kop), " --> ", FormatDollars(cents), "|")
	fmt.Println("-----------------------------------------------")
}

func RubToCentsWithExternalRate(r Rubles, c Kopecks, rate float64) Cents {
	usd := (float64(r) + (float64(c) / 100)) * rate

	return Cents(usd * 100)
}

func GetUsdRate() (float64, error) {
	resp, err := http.Get("https://open.er-api.com/v6/latest/RUB")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data ExchangeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.Rates["USD"], nil
}
